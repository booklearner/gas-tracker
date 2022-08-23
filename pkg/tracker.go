package tracker

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type Gas struct {
	GasPrice            int64  `json:"gas_price"`
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
	g.GasPrice = gp.Int64()
	g.PendingTransactions = tc

	return g, nil
}
