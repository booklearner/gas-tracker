package tracker

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type Gas struct {
	PriceGwei           int64  `json:"gas_price_gwei"`
	PriceWei            int64  `json:"gas_price_wei"`
	BlockNumber         uint64 `json:"block_number"`
	PendingTransactions uint   `json:"pending_transactions"`
}

type Client struct {
	ethClient *ethclient.Client
}

func InitClient() (*Client, error) {
	node := viper.GetString("node")
	c, err := ethclient.Dial(node)
	if err != nil {
		return nil, err
	}
	return &Client{ethClient: c}, nil
}

func (client *Client) Close() {
	client.ethClient.Close()
}

func (client *Client) GetGas() (Gas, error) {
	ctx := context.Background()
	g := Gas{}

	bn, err := client.ethClient.BlockNumber(ctx)
	if err != nil {
		return g, err
	}
	gp, err := client.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return g, err
	}
	tc, err := client.ethClient.PendingTransactionCount(ctx)
	if err != nil {
		return g, err
	}

	g.BlockNumber = bn
	g.PriceGwei = new(big.Int).Div(gp, big.NewInt(1_000_000_000)).Int64()
	g.PriceWei = gp.Int64()
	g.PendingTransactions = tc

	return g, nil
}
