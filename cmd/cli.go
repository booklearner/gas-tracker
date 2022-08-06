package main

import (
	"fmt"
	"github.com/booklearner/gas-tracker/pkg"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "the version of the gas-tracker command line application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v" + tracker.GetCliVersion())
	},
}

var gasCmd = &cobra.Command{
	Use:   "gas",
	Short: "get current gas from the Ethereum network",
	Run: func(cmd *cobra.Command, args []string) {
		var s string = "Current gas: "
		var d int = tracker.GetGas()
		fmt.Fprintf(os.Stdout, "%s %d\n", s, d)
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run the gas tracker as a daemon, exposing simple API to get gas prices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("...")
	},
}

var EthereumNodeAddress string
var ServerBindAddress string

func main() {
	var cmd = &cobra.Command{Use: "gas-tracker"}
	cmd.PersistentFlags().StringVarP(&EthereumNodeAddress, "node", "n", "eth.booklearner.org", "address for the Ethereum Node to make the RPC calls to")
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(gasCmd)
	cmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&ServerBindAddress, "bind", "b", "0.0.0.0:5001", "local address and port to bind to when running as a daemon")
	cmd.Execute()
}
