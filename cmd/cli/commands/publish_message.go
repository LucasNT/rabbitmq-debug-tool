package commands

import (
	"os"

	"github.com/LucasNT/rabbitmq-helper-tool/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	publishMessage = &cobra.Command{
		Use:   "publish",
		Short: "publish message to rabbitmq Instance",
		Long:  "Read messagens from the stdin, each line is a message.",
		Run:   publishCommand,
	}
	publishExchange    string
	publishRoutingKey  string
	publishContentType string
)

func init() {
	rootComand.AddCommand(publishMessage)
	flags := publishMessage.Flags()
	flags.StringVarP(&publishExchange, "exchance", "e", "", "exchange name, empty is the default exchange (default \"\")")
	flags.StringVarP(&publishRoutingKey, "key", "k", "", "routing key to pass to exchnge")
	publishMessage.MarkFlagRequired("key")
	flags.StringVarP(&publishContentType, "contentType", "c", "text/plain", "ContentType of the messagem")
}

func publishCommand(cmd *cobra.Command, args []string) {
	conn, err := rabbitmq.ConnectToRabbitmq(username, password, rabbitmqAddress, rabbitmqVhost)
	defer conn.Close()
	if err != nil {
		logrus.Fatalf("Failed to connect to rabbitmq: %v", err)
	}
	channel, err := rabbitmq.CreateChannel(conn)
	defer channel.Close()
	if err != nil {
		logrus.Fatalf("Failed to create channel: %v", err)
	}
	_, err = rabbitmq.PublishMessage(publishExchange, publishRoutingKey, channel, publishContentType, os.Stdin)
	if err != nil {
		logrus.Fatal(err)
	}
}
