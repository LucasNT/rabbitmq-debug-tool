package rabbitmq

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func PublishMessage(
	exchange string,
	routeKey string,
	channel *amqp091.Channel,
	contentType string,
	data io.Reader,
) (uint32, error) {
	scanner := bufio.NewScanner(data)
	buf := make([]byte, 4*1024*1024)
	scanner.Buffer(buf, 4*1024*1024)
	var count uint32 = 0
	for scanner.Scan() {
		line := scanner.Text()
		err := channel.Publish(exchange,
			routeKey,
			false,
			false,
			amqp091.Publishing{
				ContentType: contentType,
				Body:        []byte(line),
			},
		)
		if err != nil {
			fmt.Errorf("Failed to publish message: %w", err)
		}
		logrus.Debugf("Publish message: %s", line)
		count += 1
	}
	return count, nil
}
