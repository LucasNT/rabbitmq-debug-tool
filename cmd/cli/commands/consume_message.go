package commands

import (
	"strconv"

	"github.com/LucasNT/rabbitmq-helper-tool/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	consumeMessage = &cobra.Command{
		Use:   "consume",
		Short: "Consume messages from rabbitmq Instance",
		Run:   consumeCommand,
	}
	consumeQueue             string
	consumeDoAck             bool
	consumeRequeue           bool
	consumeStreamOffset      string
	consumeMessagesToConsume uint32
	//consumeQueueArguments []string
)

func init() {
	rootComand.AddCommand(consumeMessage)
	flags := consumeMessage.Flags()
	flags.StringVarP(&consumeQueue, "queue", "q", "", "Name of the queue that will be consumed")
	flags.BoolVarP(&consumeDoAck, "autoack", "a", false, "If the message has auto ack")
	flags.BoolVarP(&consumeRequeue, "requeue", "R", false, "If the message will not be requeued, need the autoack to be false")
	flags.StringVar(&consumeStreamOffset, "stream-offset", "", "Stream offset in the case that queue is a stream")
	flags.Uint32VarP(&consumeMessagesToConsume, "messages-to-consume", "n", 1, "Amount of messages to be requested")
	//flags.StringArrayVar(&consumeQueueArguments, "args", nil, "Pass arguments for the consume, in the  format label=value")
}

func consumeCommand(cmd *cobra.Command, args []string) {
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
	var queueArgs map[string]interface{} = make(map[string]interface{})
	if consumeStreamOffset != "" {
		if num, err := strconv.Atoi(consumeStreamOffset); err == nil {
			queueArgs["x-stream-offset"] = num
		} else {
			queueArgs["x-stream-offset"] = consumeStreamOffset
		}
	}
	err = rabbitmq.Consume(channel, consumeQueue, consumeDoAck, !consumeRequeue, consumeMessagesToConsume, queueArgs)
	if err != nil {
		logrus.Fatal(err)
	}
}
