package tracker

import (
	"context"
	"math/big"
)

type Gas struct {
	PriceGwei           int64  `json:"gas_price_gwei"`
	PriceWei            int64  `json:"gas_price_wei"`
	BlockNumber         uint64 `json:"block_number"`
	PendingTransactions uint   `json:"pending_transactions"`
}

func GetGas() (Gas, error) {
	ctx := context.Background()
	g := Gas{}

	bn, err := client.BlockNumber(ctx)
	if err != nil {
		return g, err
	}
	gp, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return g, err
	}
	tc, err := client.PendingTransactionCount(ctx)
	if err != nil {
		return g, err
	}

	g.BlockNumber = bn
	g.PriceGwei = new(big.Int).Div(gp, big.NewInt(1_000_000_000)).Int64()
	g.PriceWei = gp.Int64()
	g.PendingTransactions = tc

	return g, nil
}
