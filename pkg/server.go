package tracker

import (
	"fmt"

	"github.com/spf13/pflag"
)

func RunServer(f pflag.FlagSet) {
	node, err := f.GetString("node")
	if err != nil {
		panic("node flag not set")
	}

	bind, err := f.GetString("bind")
	if err != nil {
		panic("bind address not set")
	}

	fmt.Println(bind)
	fmt.Println(node)
}
