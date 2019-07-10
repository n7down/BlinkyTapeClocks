package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	if string(stderr.Bytes()) != "" {
		return "", errors.New(string(stderr.Bytes()))
	}
	return string(stdout.Bytes()), nil
}

func GetHostName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	hostname = strings.TrimSuffix(hostname, ".domain")
	hostnameTitle := strings.Title(hostname)
	return hostnameTitle, nil
}
