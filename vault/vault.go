package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	vaultapi "github.com/hashicorp/vault/api"
)

// Client is a wrapper around the vault client
type Client struct {
	Client *vaultapi.Client
}

const (
	secretShares    = 5
	secretThreshold = 3
	initFileName    = "initInfo.json"
)

// NewClient creates a new Vault Client
func (c *Client) NewClient() error {
	vaultClient, err := vaultapi.NewClient(vaultapi.DefaultConfig())
	if err != nil {
		return err
	}
	c.Client = vaultClient
	return nil
}

// InitializeVault checks if vault is initialized and if not, initializes.
func (c *Client) InitializeVault() error {
	// Creates a Sys to return the client for sys-related API calls.
	sys := c.Client.Sys()
	// Checks if Vault is initialized already
	init, err := sys.InitStatus()
	if err != nil {
		return err
	}
	// If Vault is not initialized, we initialize here and save its keys and token on dick for now.
	if !init {
		initRes, err := sys.Init(&vaultapi.InitRequest{SecretShares: secretShares, SecretThreshold: secretThreshold})
		if err != nil {
			return err
		}
		// Here we save the keys and token on the disk
		bs, _ := json.Marshal(initRes)
		fmt.Println(bs)
		err = ioutil.WriteFile(initFileName, bs, 0777)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	return nil
}

// CheckSeal checks if Vault is sealed and if it is, it unseals.
func (c *Client) CheckSeal() (vaultapi.InitResponse, error) {
	sys := c.Client.Sys()

	sealStatus, err := sys.SealStatus()
	if err != nil {
		return vaultapi.InitResponse{}, err
	}
	var initRes vaultapi.InitResponse
	// If Vault is Sealed we unseal
	if sealStatus.Sealed {
		bs, err := ioutil.ReadFile(initFileName)
		if err != nil {
			return vaultapi.InitResponse{}, err
		}
		json.Unmarshal(bs, &initRes)
		keysUsed := 0
		// Here we go through the stored keys until the  Vault is unSealed
		for sealStatus.Sealed {
			sealStatus, err = sys.Unseal(initRes.Keys[keysUsed])
			if err != nil {
				return vaultapi.InitResponse{}, err
			}
			log.Println("Is sealed", sealStatus.Sealed)
			keysUsed++
		}
	}
	return initRes, nil
}
