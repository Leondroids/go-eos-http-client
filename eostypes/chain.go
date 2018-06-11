package eostypes

import (
	"encoding/json"
)

/*
	ChainGetInfoResponse

	Example Response:
	{
  		"server_version": "0961a560",
  		"chain_id": "cf057bbfb72640471fd910bcb67639c22df9f92470936cddc1ade0e2f2e7dc4f",
  		"head_block_num": 349,
  		"last_irreversible_block_num": 348,
  		"last_irreversible_block_id": "0000015c333a10b501cfca6786f3237072456ab039f6fe412ac7e1dbc10e1e4b",
  		"head_block_id": "0000015d6ad4a66ff184fe73f8720ab3070a55e453e1ef11abfe694dc46f840a",
  		"head_block_time": "2018-06-08T14:54:53",
  		"head_block_producer": "eosio",
  		"virtual_block_cpu_limit": 283089,
  		"virtual_block_net_limit": 1485079,
  		"block_cpu_limit": 199900,
  		"block_net_limit": 1048576
	}
 */
type ChainGetInfoResponse struct {
	ServerVersion            string `json:"server_version"`
	ChainID                  string `json:"chain_id"`
	HeadBlockNum             int64  `json:"head_block_num"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID  string `json:"last_irreversible_block_id"`
	HeadBlockID              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
	VirtualBlockCPULimit     int64  `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit     int64  `json:"virtual_block_net_limit"`
	BlockCPULimit            int64  `json:"block_cpu_limit"`
	BlockNetLimit            int64  `json:"block_net_limit"`
}

func (cir *ChainGetInfoResponse) FromJson(b []byte) (*ChainGetInfoResponse, error) {
	err := json.Unmarshal(b, cir)

	if err != nil {
		return nil, err
	}

	return cir, nil
}

////////////////////////////////////////////////////////////////////

type ChainGetBlockRequest struct {
	BlockNumOrID string `json:"block_num_or_id"`
}
