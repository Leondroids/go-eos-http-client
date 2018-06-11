package main

import (
	"github.com/Leondroids/go-eos-http-client/eoshttp"
	"log"
)

func main() {
	client := eoshttp.NewEOSClient()

	info, err := client.Chain.GetInfo()

	if err != nil {
		panic(err)
	}

	log.Println("Info: ", info)

	createWallet, alreadExisted, err := client.Wallet.CreateWallet("test6")

	if err != nil {
		panic(err)
	}

	if alreadExisted{
		log.Println("Wallet already exists")
	} else {
		log.Println("Wallet created: ", createWallet)
	}
}
