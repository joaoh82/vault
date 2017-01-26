package vault

import (
	"fmt"
	"log"

	vaultapi "github.com/hashicorp/vault/api"
)

// Client is a wrapper around the vault client
type Client struct {
	Client *vaultapi.Client
}

const (
	token           = "726aeeac-2df2-0c89-22a2-5bee019a912d"
	secretShares    = 5
	secretThreshold = 3
	backendPath     = "etcd"
)

var vaultKeys []string

func (c *Client) NewClient() error {
	vaultClient, err := vaultapi.NewClient(vaultapi.DefaultConfig())
	if err != nil {
		return err
	}
	c.Client = vaultClient
	return nil
}

func (c *Client) InitializeVault() (bool, error) {
	// Creates a Sys to return the client for sys-related API calls.
	sys := c.Client.Sys()
	// Checks if Vault is initialized already
	init, err := sys.InitStatus()
	if err != nil {
		return false, err
	}
	// If Vault is not initialized, we initialize here and save its keys and token on dick for now.
	if !init {
		initRes, err := sys.Init(&vaultapi.InitRequest{SecretShares: secretShares, SecretThreshold: secretThreshold})
		if err != nil {
			return false, err
		}
		// Here we save the keys and token on the disk
		fmt.Println("initResponse", initRes)
	}
	return true, nil
}

func (c *Client) CheckSeal() {
	vaultKeys = append(vaultKeys, "xQIXfz9KewZOSIC1zTjMFbdj8+5oQn6WpeWMIafSMOgB")
	vaultKeys = append(vaultKeys, "z49U6x2c84Sfz/jZZd4e2+2Ll+R5zQB+K3EjT6FAuQMC")
	vaultKeys = append(vaultKeys, "MeBfs74y5Hl9LQcyGZY6EuRk0pzlf7f5XXZ93tj0lkgD")
	vaultKeys = append(vaultKeys, "hELvT9jou0tXvvl4ss/R3Aenbkx7mYY8aJ0EYCJlhSAE")
	vaultKeys = append(vaultKeys, "ei3kF3tGrLa1XAaTzof1FQ5IKzTnKzG7Hppa8VvRqmsF")

	sys := c.Client.Sys()

	sealStatus, err := sys.SealStatus()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is sealed:", sealStatus.Sealed)
	// If Vault is Sealed we unseal
	if sealStatus.Sealed {
		keysUsed := 0
		// Here we go through the stored keys until the  Vault is unSealed
		for sealStatus.Sealed {
			sealStatus, err = sys.Unseal(vaultKeys[keysUsed])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Is sealed", sealStatus.Sealed)
			keysUsed++
		}
	}
}
