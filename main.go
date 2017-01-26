package main

import (
	"log"

	"github.com/joaoh82/vault/vault"
)

var client vault.Client

func main() {

	// Creating a new Vault Client
	err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	err = client.InitializeVault()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.CheckSeal()
	if err != nil {
		log.Fatal(err)
	}
	// After Vault is initialized and Unsealed we set the Token for the Vault so we can perform operations.
	client.Client.SetToken(res.RootToken)

	//c := vaultClient.Logical()
	//sec, err := c.Read("secret/myfirstkey")
	//		sec, err := c.Write("secret/mysecondkey",
	//			map[string]interface{}{
	//				"hello2": "world2",
	//			})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(sec)
}
