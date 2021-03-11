package wallet

//import (
//	"github.com/joeqian10/neo3-gogogo/crypto"
//	"github.com/joeqian10/neo3-gogogo/helper"
//	"github.com/joeqian10/neo3-gogogo/rpc"
//	"github.com/joeqian10/neo3-gogogo/sc"
//	"github.com/joeqian10/neo3-gogogo/tx"
//	"github.com/stretchr/testify/assert"
//	"log"
//	"math/big"
//	"testing"
//)
//
//var lh = rpc.NewClient("http://localhost:20002")
//
//var wh, _ = NewWalletHelperFromNEP2(lh, "6PYMKphknCsVDBLD9TgYRX1m8H4ra4BfhyRKMAcZ6Fz49p7UhHg19c5vjy", "t", 16384, 8, 8)
//
//var magic uint32 = 20210311
//
//func TestWalletHelper_ClaimGas2(t *testing.T) {
//	h, e := wh.ClaimGas(magic)
//	assert.Nil(t, e)
//	assert.Less(t, 0, len(h))
//}
//
//func TestWalletHelper_GetAccountAndBalance2(t *testing.T) {
//	r, e := wh.GetAccountAndBalance(tx.GasToken)
//	assert.Nil(t, e)
//	assert.Equal(t, 1, len(r))
//	acc, e := crypto.AddressToScriptHash("Nf9wYEnnurkYZpbCPTVQQ2eeiSLr4eDCL5", helper.DefaultAddressVersion)
//	assert.Nil(t, e)
//	assert.Equal(t, 0, r[0].Account.CompareTo(acc))
//	x := r[0].Value.Uint64()
//	assert.Less(t, uint64(0), x)
//}
//
//func TestWalletHelper_GetBalanceFromAccount2(t *testing.T) {
//	acc, e := crypto.AddressToScriptHash("Nf9wYEnnurkYZpbCPTVQQ2eeiSLr4eDCL5", helper.DefaultAddressVersion)
//	assert.Nil(t, e)
//	b, e := wh.GetBalanceFromAccount(tx.GasToken, acc)
//	assert.Nil(t, e)
//	x := b.Uint64()
//	assert.Less(t, uint64(0), x)
//}
//
//func TestWalletHelper_GetBalanceFromWallet2(t *testing.T) {
//	b, e := wh.GetBalanceFromWallet(tx.GasToken, nil)
//	assert.Nil(t, e)
//	x := b.Uint64()
//	assert.LessOrEqual(t, uint64(0), x)
//}
//
//func TestWalletHelper_GetBlockHeight2(t *testing.T) {
//	height, err := wh.GetBlockHeight()
//	assert.Nil(t, err)
//	assert.Less(t, uint32(0), height)
//}
//
//func TestWalletHelper_GetContractState2(t *testing.T) {
//	c, e := wh.GetContractState(tx.GasToken)
//	assert.Nil(t, e)
//	assert.Equal(t, "0xd2a4cff31913016155e38e474a2c06d08be276cf", c.Hash)
//}
//
//func TestWalletHelper_GetGasConsumed2(t *testing.T) {
//	script, e := sc.MakeScript(tx.GasToken, "symbol", nil)
//	assert.Nil(t, e)
//	log.Println(crypto.Base64Encode(script)) // wh8MBnN5bWJvbAwUz3bii9AGLEpHjuNVYQETGfPPpNJBYn1bUg==
//	b, e := wh.GetGasConsumed(script, nil)
//	assert.Nil(t, e)
//	assert.Less(t, int64(0), b)
//}
//
//func TestWalletHelper_GetUnClaimedGas2(t *testing.T) {
//	v, e := wh.GetUnClaimedGas()
//	assert.Nil(t, e)
//	assert.Less(t, uint64(0), v)
//}
//
//func TestWalletHelper_MakeTransaction2(t *testing.T) {
//	// tested in Transfer2
//}
//
//func TestWalletHelper_SignTransaction2(t *testing.T) {
//	// tested in Transfer2
//}
//
//func TestWalletHelper_Transfer2(t *testing.T) {
//	to := "Nf9wYEnnurkYZpbCPTVQQ2eeiSLr4eDCL5"
//	scriptHash, _ := crypto.AddressToScriptHash(to, helper.DefaultAddressVersion)
//	log.Println(scriptHash.String())
//	h, e := wh.Transfer(tx.NeoToken, to, big.NewInt(10000000), magic)
//	assert.Nil(t, e)
//	assert.Less(t, 0, len(h))
//}
