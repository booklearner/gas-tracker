package main

import (
	"fmt"
	"github.com/booklearner/gas-tracker/pkg"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of the gas-tracker command line application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v" + tracker.GetCliVersion())
	},
}

var gasCmd = &cobra.Command{
	Use:   "gas",
	Short: "Get current gas from the Ethereum network",
	Run: func(cmd *cobra.Command, args []string) {
		var s string = "Current gas: "
		var d int = tracker.GetGas()
		fmt.Fprintf(os.Stdout, "%s %d\n", s, d)
	},
}

func main() {
	var cmd = &cobra.Command{Use: "gas-tracker"}
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(gasCmd)
	cmd.Execute()
}
