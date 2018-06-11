package eoshttp

import "testing"

func TestChain_GetInfo(t *testing.T) {
	client := NewEOSClient()

	info, err := client.Chain.GetInfo()

	if err != nil {
		t.Error(err)
		return
	}

	CheckStringNotEmpty(info.ServerVersion, "ServerVersion", t)
	CheckStringNotEmpty(info.ChainID, "ChainID", t)
	CheckInt64NotEmpty(info.HeadBlockNum, "HeadBlockNum", t)
	CheckInt64NotEmpty(info.LastIrreversibleBlockNum, "LastIrreversibleBlockNum", t)
	CheckStringNotEmpty(info.LastIrreversibleBlockID, "LastIrreversibleBlockID", t)
	CheckStringNotEmpty(info.HeadBlockID, "HeadBlockID", t)
	CheckStringNotEmpty(info.HeadBlockTime, "HeadBlockTime", t)
	CheckStringNotEmpty(info.HeadBlockProducer, "HeadBlockProducer", t)
	CheckInt64NotEmpty(info.VirtualBlockCPULimit, "VirtualBlockCPULimit", t)
	CheckInt64NotEmpty(info.VirtualBlockNetLimit, "VirtualBlockNetLimit", t)
	CheckInt64NotEmpty(info.BlockCPULimit, "BlockCPULimit", t)
	CheckInt64NotEmpty(info.BlockNetLimit, "BlockNetLimit", t)
}



// AUX

func CheckStringNotEmpty(value string, label string, t *testing.T) {
	if value == "" {
		CannotEmpty(label, t)
	}
}

func CheckInt64NotEmpty(value int64, label string, t *testing.T) {
	if value == 0 {
		CannotEmpty(label, t)
	}
}

func CannotEmpty(label string, t *testing.T) {
	t.Errorf("%v cannot be empty", label)
}
