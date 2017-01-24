package main

import vaultapi "github.com/hashicorp/vault/api"

func main() {

	client, err := vaultapi.NewClient(vaultapi.Config())
	if err != nil {
		panic(err)
	}

}
