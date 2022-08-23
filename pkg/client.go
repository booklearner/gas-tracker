package tracker

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

var client *ethclient.Client

func InitClient() error {
	node := viper.GetString("node")
	c, err := ethclient.Dial(node)

	if err != nil {
		return err
	}

	client = c
	return nil
}
