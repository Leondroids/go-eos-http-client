package eoshttp

import "github.com/Leondroids/go-eos-http-client/eostypes"

const (
	ChainPath = "chain"
)

type Chain struct {
	path   string
	client *EOSClient
}

func newChain(client *EOSClient) *Chain {
	return &Chain{
		client: client,
		path:   ChainPath,
	}
}

/*
	method: /chain/get_info
	Returns an object containing various details about the blockchain.
	curl --request POST --url http://127.0.0.1:8888/v1/chain/get_info
 */
func (it *Chain) GetInfo() (*eostypes.ChainGetInfoResponse, error) {
	result := it.client.Call(it.createRequest("get_info", nil))

	if !result.Success {
		return nil, result.Error
	}

	return new(eostypes.ChainGetInfoResponse).FromJson(result.Body)
}

func (it *Chain) createRequest(path string, body []byte) *EOSRequest {
	return it.client.NewEOSJSONRequest(it.path, path, body)
}
