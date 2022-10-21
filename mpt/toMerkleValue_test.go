package mpt

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeserializeMerkleValue(t *testing.T) {
	value := helper.HexToBytes("201ca7d4dcecd3705602d6e285deafa2c8240f2883045d237822ae33d50451f59f208d735f64159b5f9c7a5670dbab887dbb895c012f39003e5fbc55b6fbff9df44414cc9f88a9e96be8e91131b06ea674fbba51c7c99e0500000000000000144f5f702b3f459f222d371052940bb9ce2d86d2ed06756e6c6f636b4a14e14fdd69cf7bf6afb9265ac806e09fea438df7b81425820465d41a57dca24529e88387ac2d787227780f42400000000000000000000000000000000000000000000000000000000000")
	tmv, err := DeserializeCrossChainTxParameter(value, 0)
	assert.Nil(t, err)
	fmt.Println(helper.BytesToHex(tmv.TxHash))
}

func TestDeserializeArgs(t *testing.T) {
	var m = make(map[string]bool)
	m["lock"] = true
	m["unlock"] = true

	b := m["hello"]
	//if !ok {
	//	fmt.Println(ok)
	//}
	fmt.Println(b)
}

func TestMerkleProve(t *testing.T) {
	path := helper.HexToBytes("ef204ce9a62083b29dd888394c124fc57f0b8533171ed3f869dd350b0375d8027c00020000000000000020000000000000000000000000000000000000000000000000000000000000e269204df8f0fc252d2bd3f21c474510fed914cf5fb5ba98510ddfe83b3d6d5a3715ff14250e76987d838a75310c34bf422ea9f1ac4cc9060e0000000000000014cb569453781497dcb067b73d95b28802cb01553806756e6c6f636b4a149328aec1e84c93855e2fb4a01f5eb7ec15e1abd614e9cdc1efd22c74b5706f0068f79b69b46fa85a0d2035f50500000000000000000000000000000000000000000000000000000000")
	crossStateRoot := helper.HexToBytes("d0acc6ea0e2cd2560ee298d4846bec230f590879090c83b235728488d1ab0fe0")
	tmv, err := MerkleProve(path, crossStateRoot)
	assert.Nil(t, err)

	fmt.Println(helper.BytesToHex(tmv))
}
