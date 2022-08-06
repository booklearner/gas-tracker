package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "The version of the gas-tracker command line application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version")
		// fmt.Println("Print: " + version)
	},
}

func main() {
	var cmd = &cobra.Command{Use: "gas-tracker"}
	cmd.AddCommand(version)
	cmd.Execute()
}
