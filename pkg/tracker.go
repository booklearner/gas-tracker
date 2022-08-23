package tracker

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type Gas struct {
	PriceGwei        int64  `json:"gas_price_gwei"`
	PriceWei         int64  `json:"gas_price_wei"`
	BlockNumber         uint64 `json:"block_number"`
	PendingTransactions uint   `json:"pending_transactions"`
}

func initClient() (*ethclient.Client, error) {
	node := viper.GetString("node")
	c, err := ethclient.Dial(node)

	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetGas() (Gas, error) {
	g := Gas{}

	ctx := context.Background()
	c, err := initClient()
	if err != nil {
		return g, err
	}

	bn, err := c.BlockNumber(ctx)
	if err != nil {
		return g, err
	}
	gp, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return g, err
	}
	tc, err := c.PendingTransactionCount(ctx)
	if err != nil {
		return g, err
	}

	g.BlockNumber = bn
	g.PriceGwei = new(big.Int).Div(gp, big.NewInt(1_000_000_000)).Int64()
	g.PriceWei = gp.Int64()
	g.PendingTransactions = tc

	return g, nil
}
