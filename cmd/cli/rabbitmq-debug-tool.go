package main

import (
	"github.com/LucasNT/rabbitmq-helper-tool/cmd/cli/commands"
	"github.com/sirupsen/logrus"
)

var ()

func main() {
	if err := commands.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
