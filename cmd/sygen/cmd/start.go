package cmd

import (
	"fmt"
	"github.com/mchudnovskiy/sygen/pkg/server"
	"github.com/mchudnovskiy/sygen/pkg/server/settings"

	"github.com/spf13/cobra"
)

var a = &settings.Args{}
var payloadFilePath string
var headersFilePath string

func init() {
	startCmd.Flags().IntVarP(&a.RequestRate, "rate", "r", 1, "Request rate per second")
	startCmd.Flags().IntVarP(&a.ExecutionTime, "time", "t", 10, "Execution time in seconds")
	startCmd.Flags().StringVarP(&a.Endpoint, "endpoint", "e", "", "Endpoint for traffic. Example: http||localhost:90/test or mq||localhost(1414)/TEST.SVRCONN/TEST.QUEUE")
	startCmd.Flags().StringVarP(&payloadFilePath, "payload", "p", "", "Payload file path")
	startCmd.Flags().StringVarP(&headersFilePath, "headers", "", "", "Headers file path")
	startCmd.MarkFlagRequired("endpoint")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Sygen traffic generator",
	Run: func(cmd *cobra.Command, args []string) {
		a.Payload = payloadFilePath         //TODO open and read body file
		a.Headers = make(map[string]string) //TODO open  and read headers file and conver it's conent to map
		s := server.New(a)
		if err := s.Start(); err != nil {
			fmt.Println("error")
		}
	},
}
