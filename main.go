package main

import (
	"log"

	"github.com/joaoh82/vault/vault"
)

const (
	token = "726aeeac-2df2-0c89-22a2-5bee019a912d"
)

var client vault.Client

func main() {

	// Creating a new Vault Client
	err := client.NewClient()
	if err != nil {
		panic(err)
	}

	ok, err := client.InitializeVault()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		log.Println("initialized")
		client.CheckSeal()
	}
	// After Vault is initialized and Unsealed we set the Token for the Vault so we can perform operations.
	client.Client.SetToken(token)

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
