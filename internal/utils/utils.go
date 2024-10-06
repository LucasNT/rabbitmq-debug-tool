package utils

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword() (string, error) {
	fmt.Fprintln(os.Stderr, "Password: ")
	bytePwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("Failed to read password, %w", err)
	}
	pass := strings.TrimSpace(string(bytePwd))
	return pass, nil
}
