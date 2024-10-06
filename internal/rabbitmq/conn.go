package rabbitmq

import (
	"fmt"
	"net/url"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitmq(
	username string,
	password string,
	rabbitmqAddress string,
	rabbitmqVhost string,
) (*amqp091.Connection, error) {

	var credentials *url.Userinfo = url.UserPassword(username, password)
	rabbitmqUrl, err := url.Parse(rabbitmqAddress)
	if err != nil {
		return nil, fmt.Errorf("Failed to pase URL, %w", err)
	}
	rabbitmqUrl.User = credentials
	rabbitmqUrl.Path = rabbitmqVhost

	return amqp091.Dial(rabbitmqUrl.String())
}
