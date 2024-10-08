package commands

import (
	"io"
	"os"
	"strings"

	"github.com/LucasNT/rabbitmq-helper-tool/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootComand = &cobra.Command{
		Use:               os.Args[0],
		Short:             "Programn used to get and send messages to rabbitmq instance",
		PersistentPreRunE: initializeCommand,
	}
	username              string
	password              string
	passwordFile          string
	readPasswordFromInput bool
	rabbitmqAddress       string
	rabbitmqVhost         string
	enableDebug           bool
)

func init() {
	flags := rootComand.PersistentFlags()
	flags.StringVarP(&username, "user", "u", "", "username for the rabbitmq authentication")
	rootComand.MarkPersistentFlagRequired("user")

	flags.StringVarP(&passwordFile, "password-file", "P", "", "file path for a file containg the password for the user of the rabbitmq")
	flags.BoolVarP(&readPasswordFromInput, "password-stdin", "p", false, "Read password from stdin")
	rootComand.MarkFlagsOneRequired("password-file", "password-stdin")
	rootComand.MarkFlagsMutuallyExclusive("password-file", "password-stdin")

	flags.StringVar(&rabbitmqAddress, "address", "amqp://localhost:5672", "Rabbimq address")

	flags.StringVar(&rabbitmqVhost, "vhost", "/", "Rabbitmq Vhost")

	flags.BoolVarP(&enableDebug, "debug", "d", false, "Enable debug logs")

}

func initializeCommand(cmd *cobra.Command, args []string) error {
	// validate arguments
	if username == "" {
		logrus.Fatal("Username can't be empty")
	}
	if readPasswordFromInput {
		var err error
		password, err = utils.ReadPassword()
		if err != nil {
			logrus.Fatalf("could not read the password from stdin %s", err)
		}
	} else {
		file, err := os.Open(passwordFile)
		if err != nil {
			logrus.Fatalf("Failed to open the password file, %s", err)
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			logrus.Fatalf("Could not read the password file, %s", err)
		}
		password = strings.TrimSpace(string(data))
	}
	if enableDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
}

func Execute() error {
	return rootComand.Execute()
}
