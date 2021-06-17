package main

import (
	"log"
	"os"

	"github.com/safe2008/vault-actions/pkg/vault"
)

func main() {
	address := os.Getenv("INPUT_ADDRESS")
	token := os.Getenv("INPUT_TOKEN")
	secrets := os.Getenv("INPUT_SECRETS")
	result := vault.GetSecret(address, token, secrets)
	log.Print(result)
}
