package tracker

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func InitClient() (*ethclient.Client, error) {
	node := viper.GetString("node")
	c, err := ethclient.Dial(node)

	if err != nil {
		return nil, err
	}
	return c, nil
}
