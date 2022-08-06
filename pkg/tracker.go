package tracker

import (
	"fmt"

	"github.com/spf13/pflag"
	"os"
)

func GetGas(f pflag.FlagSet) {
	node, err := f.GetString("node")
	if err != nil {
		panic("node flag not set")
	}

	fmt.Println(node)
	var s string = "Current gas: "
	var d int = 0
	fmt.Fprintf(os.Stdout, "%s %d\n", s, d)
}
