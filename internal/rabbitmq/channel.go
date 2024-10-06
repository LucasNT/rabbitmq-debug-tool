package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func CreateChannel(conn *amqp091.Connection) (*amqp091.Channel, error) {
	return conn.Channel()
}
