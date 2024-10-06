package commands

import (
	"github.com/LucasNT/rabbitmq-debug-tool/internal/rabbitmq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	validateUser = &cobra.Command{
		Use:   "validate",
		Short: "Validate username and password",
		Run:   validateCommand,
	}
)

func init() {
	rootComand.AddCommand(validateUser)
}

func validateCommand(cmd *cobra.Command, args []string) {
	conn, err := rabbitmq.ConnectToRabbitmq(username, password, rabbitmqAddress, rabbitmqVhost)
	defer conn.Close()
	if err != nil {
		logrus.Fatalf("Failed to connect to rabbitmq: %v", err)
	}
}
