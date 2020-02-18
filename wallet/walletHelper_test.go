package wallet

import (
	"testing"

	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletHelper(t *testing.T) {
	txBuilder := tx.NewTransactionBuilder("http://seed1.ngd.network:20332")
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.NotNil(t, txBuilder)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	walletHelper := NewWalletHelper(txBuilder.Client, account)
	assert.NotNil(t, walletHelper)
	assert.Equal(t, "http://seed1.ngd.network:20332", walletHelper.Client.GetUrl())
	assert.Equal(t, "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619", walletHelper.Account.KeyPair.PublicKey.String())
}

func TestWalletHelper_ClaimGas(t *testing.T) {

}

func TestWalletHelper_Transfer(t *testing.T) {

}

//func TestWallet_Transfer(t *testing.T) {
//	client := rpc.NewClient("http://127.0.0.1:10332")
//	client.SetBasicAuth("krain", "123456")
//	account, _ := NewAccountFromWIF("KyoYyZpoccbR6KZ25eLzhMTUxREwCpJzDsnuodGTKXSG8fDW9t7x")
//	address := "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
//	api := NewWalletHelper(client, account)
//	gasBalace, _ := api.GetBalance(tx.GasToken, address)
//
//	assert.Equal(t, true, gasBalace.Int64() > 0)
//
//	result, err := api.Transfer(tx.NeoToken, "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW", big.NewInt(100))
//	assert.Nil(t, err)
//	assert.True(t, len(result) > 0)
//	fmt.Println(result)
//
//	claimable, _ := api.GetUnClaimedGas(address)
//	assert.True(t, claimable > 0)
//	fmt.Println(claimable)
//
//	result, _ = api.ClaimGas()
//	assert.True(t, len(result) > 0)
//	fmt.Println(result)
//}

// func TestWallet_NEP5Helper(t *testing.T) {
// 	client := rpc.NewClient("http://127.0.0.1:10332")
// 	nep5Api := nep5.NewNep5Helper(client)

// 	total, err := nep5Api.TotalSupply(tx.NeoToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, big.NewInt(100000000), total)

// 	name, err := nep5Api.Name(tx.NeoToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "NEO", name)

// 	tokensymbol, err := nep5Api.Symbol(tx.NeoToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "neo", tokensymbol)

// 	decimals, err := nep5Api.Decimals(tx.GasToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 8, decimals)

// 	acc, _ := helper.AddressToScriptHash("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ")
// 	balance, err := nep5Api.BalanceOf(tx.GasToken, acc)
// 	assert.Nil(t, err)
// 	assert.Equal(t, true, balance.Int64() > 0)

// }
