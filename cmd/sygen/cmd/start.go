package cmd

import (
	"fmt"
	"github.com/mchudnovskiy/sygen/pkg/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Sygen traffic generator",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.New()
		if err := s.Start(); err != nil {
			fmt.Println("error")
		}
	},
}