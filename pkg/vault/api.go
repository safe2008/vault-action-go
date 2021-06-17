package vault

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vault/api"
)

const error = "not found"

// GetSecret get secret form vault.

func GetSecret(address string, token string, secrets string) string {
	s := strings.Fields(secrets)

	config := &api.Config{
		Address: address,
	}

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return error
	}

	client.SetToken(token)

	secret, err := client.Logical().Read(s[0])
	if err != nil {
		fmt.Println(err)
		return error
	}

	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return error
	}

	return fmt.Sprintf("%v", m[s[1]])
}
