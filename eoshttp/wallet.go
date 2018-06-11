package eoshttp

import "strings"

const (
	WalletPath = "wallet"
)

type Wallet struct {
	path   string
	client *EOSClient
}

func newWallet(client *EOSClient) *Wallet {
	return &Wallet{
		client: client,
		path:   WalletPath,
	}
}

/*
	Method: wallet/create_wallet
	Creates a new wallet with the given name.
	curl --request POST --url http://127.0.0.1:8888/v1/wallet/create
	Returns password(string)
 */
func (it *Wallet) CreateWallet(name string) (string, bool, error) {
	result := it.client.Call(it.createTextRequest("create", []byte(name)))

	if !result.Success {
		if result.IsEOSError && result.EOSError.Error.Name == "wallet_exist_exception" {
			return "", true, nil

		}
		return "", false, result.Error
	}

	return strings.Replace(string(result.Body), "\"", "", -1), false, nil
}

func (it *Wallet) createTextRequest(path string, body []byte) *EOSRequest {
	return it.client.NewEOSJSONRequest(it.path, path, body)
}
