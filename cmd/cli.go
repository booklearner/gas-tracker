package main

import (
	"fmt"
	"os"

	tracker "github.com/booklearner/gas-tracker/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "the version of the gas-tracker command line application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v" + tracker.GetVersion())
	},
}

var gasCmd = &cobra.Command{
	Use:   "gas",
	Short: "get current gas from the Ethereum network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "%d\n", tracker.GetGas())
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run the gas tracker as a daemon, exposing simple API to get gas prices",
	Run: func(cmd *cobra.Command, args []string) {
		tracker.RunServer()
	},
}

var EthereumNodeAddress string
var ServerBindAddress string

func main() {
	var cmd = &cobra.Command{Use: "gas-tracker"}
	cmd.PersistentFlags().StringVarP(
		&EthereumNodeAddress,
		"node",
		"n",
		"eth.booklearner.org",
		"address for the Ethereum Node to make the RPC calls to",
	)
	viper.BindPFlag("node", cmd.PersistentFlags().Lookup("node"))
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(gasCmd)
	cmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(
		&ServerBindAddress,
		"bind",
		"b",
		"0.0.0.0:5001",
		"local address and port to bind to when running as a daemon",
	)
	viper.BindPFlag("bind", serverCmd.Flags().Lookup("bind"))
	cmd.Execute()
}
