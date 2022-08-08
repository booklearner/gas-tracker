package tracker

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetGas() string {
	node := viper.GetString("node")
	var s string = "Current gas: "
	var d int = 0
	return fmt.Sprintf("%s\n%s %d\n", node, s, d)
}
