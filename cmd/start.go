package cmd

import (
	"fmt"
	"github.com/spf13/cobra"


)

func init() {
  rootCmd.AddCommand(startCmd)
}


var startCmd = &cobra.Command{
  Use:   "start",
  Short: "Start Sygen traffic generator",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Sygen has been started")
  },
}

