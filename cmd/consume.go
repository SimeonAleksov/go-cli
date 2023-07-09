package cmd

import (
	"github.com/spf13/cobra"

	"github.com/SimeonAleksov/go-cli/consumer"
	"github.com/SimeonAleksov/go-cli/tui"
)

var consumeCommand = &cobra.Command{
	Use:   "consume",
	Short: "Start listening",
	Run: func(cmd *cobra.Command, args []string) {
		kafkaConsumer := consumer.NewKafkaConsumer("kafka:9092", "accounts")
		listen := tui.NewTUI(kafkaConsumer)
		go listen.Start()
		<-make(chan struct{})
		listen.Stop()

	},
}

func init() {
	rootCmd.AddCommand(consumeCommand)
}
