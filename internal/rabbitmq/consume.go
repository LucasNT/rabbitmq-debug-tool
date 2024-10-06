package rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func Consume(
	channel *amqp091.Channel,
	queueName string,
	doAck bool,
	requeue bool,
	messagesLimit uint32,
	args map[string]interface{},
) error {
	err := channel.Qos(1, 0, false)
	if err != nil {
		return fmt.Errorf("Failed to set prefetch count: %w", err)
	}
	delivery, err := channel.Consume(queueName, "", false, false, false, false, args)
	if err != nil {
		return err
	}
	for i := uint32(0); i < messagesLimit; i++ {
		del := <-delivery
		fmt.Println("======================================================")
		fmt.Printf("ContentType: %s\n", del.ContentType)
		fmt.Printf("Header: %v\n", del.Headers)
		fmt.Printf("Body: %s\n", string(del.Body))
		if doAck == false {
			del.Reject(requeue)
		} else {
			del.Ack(true)
		}

	}
	return nil
}
