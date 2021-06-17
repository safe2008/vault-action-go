package main

import (
	"os"
	"strings"

	"github.com/safe2008/vault-actions/pkg/vault"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	address := os.Getenv("INPUT_ADDRESS")
	token := os.Getenv("INPUT_TOKEN")
	secrets := os.Getenv("INPUT_SECRETS")
	s := strings.Fields(secrets)
	n := s[0] + " " + s[1]

	result := vault.GetSecret(address, token, n)
	githubactions.SetOutput(s[3], result)
}
