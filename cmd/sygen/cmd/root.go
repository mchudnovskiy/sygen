//Package cmd is a cobra commands store for Sygen project
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "sygen",
	Short: "Sygen is a very fast traffic generator for load testing",
	Long:  `A Fast and Flexible Traffic Generator built in Go.`,
}

// Execute is a runner of sygen root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
