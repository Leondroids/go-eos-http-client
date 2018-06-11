package eoshttp

import "testing"

func TestWallet_CreateWallet(t *testing.T) {
	client := NewEOSClient()

	wallet, alreadyExist, err := client.Wallet.CreateWallet("test9")

	if err != nil {
		t.Error(err)
		return
	}

	if alreadyExist {

	}

	CheckStringNotEmpty(wallet, "wallet id", t)
}
