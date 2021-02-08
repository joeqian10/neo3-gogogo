# neo3-gogogo

## 1. Overview

This is a light-weight golang SDK for Neo 3.0 network.

## 2. Getting Started

### 2.1 Installation

Simply add this SDK to GOPATH:

```bash
go get github.com/joeqian10/neo3-gogogo
```

## 3. Modules

### 3.1 "crypto" module

This module offers methods used for cryptography purposes, such as AES encryption/decryption, Base58 encoding/decoding, Hash160/Hash256 hashing functions. For more information about the crypto algorithms used in neo, refer to [Cryptography]().

#### 3.1.1 AES encryption

```golang
func AESEncrypt(src, key []byte) ([]byte, error)
```

#### 3.1.2 AES decryption

```golang
func AESDecrypt(crypted, key []byte) ([]byte, error)
```

#### 3.1.3 Base58Check encoding

```golang
func Base58CheckEncode(input []byte) string
```

#### 3.1.4 Base58Check decoding

```golang
func Base58CheckDecode(input string) ([]byte, error)
```

#### 3.1.5 Sha256 hash function

```golang
func Sha256(b []byte) []byte
```

#### 3.1.6 Hash256 function

```golang
func Hash256(ba []byte) []byte
```

`Hash256` gets the twice SHA-256 hash value of `ba`.

#### 3.1.7 Hash160 function

```golang
func Hash160(ba []byte) []byte
```

`Hash160` first calculates SHA-256 hash result of `ba`, then calcaulates RIPEMD-160 hash of the result.

*Typical usage:*

```golang
package sample

import "encoding/hex"
import "github.com/joeqian10/neo3-gogogo/crypto"

func SampleMethod() {
    // AES encryption/decryption
    src := hex.DecodeString("3b75c0cee38f1e4fa123fad71c3f3e43dc8599c9bedb3aa16e4f8b9239a6d946")
    key := hex.DecodeString("e23c14a11c4ccefda68918331cbd2caf3e680d78b72e19c1fc8675b9636d0de8")
    encrypted, err := crypto.AESEncrypt(xr, derivedKey2)
    decrypted, err := crypto.AESDecrypt(encrypted, derivedKey2)

    // Base58Check encoding/decoding
    var b58CheckEncoded = "KxhEDBQyyEFymvfJD96q8stMbJMbZUb6D1PmXqBWZDU2WvbvVs9o"
    var b58CheckDecodedHex = "802bfe58ab6d9fd575bdc3a624e4825dd2b375d64ac033fbc46ea79dbab4f69a3e01"

    b58CheckDecoded, _ := hex.DecodeString(b58CheckDecodedHex)
    encoded := crypto.Base58CheckEncode(b58CheckDecoded)
    decoded, err := crypto.Base58CheckDecode(b58CheckEncoded)

    // Sha256, Hash256, Hash160
    b := []byte("Hello World")
    s1 := crypto.Sha256(b)
    s1 := crypto.Hash256(b)
    s1 := crypto.Hash160(b)

    ...
}
```

### 3.2 "helper" module

As its name indicated, this module acts as a helper and provides some standard param types used in neo, such as `UInt160`, `UInt256`, and some auxiliary methods with basic functionalities including conversion between a hex string and a byte array, conversion between a script hash and a standard neo address, concatenating/reversing byte arrays and so on.

#### 3.2.1 Create a new UInt160 object from a big-endian hex string

```golang
func UInt160FromString(s string) (UInt160, error)
```

#### 3.2.2 Create a new UInt160 object from a littel-endian byte array

```golang
func UInt160FromBytes(b []byte) (u UInt160, err error)
```

#### 3.2.3 Convert a UInt160 object to a big-endian hex string

```golang
func (u UInt160) String() string
```

#### 3.2.4 Convert a UInt160 object to a littel-endian byte array

```golang
func (u UInt160) ToByteArray() []byte
```

#### 3.2.5 Create a new UInt256 object from a big-endian hex string

```golang
func UInt256FromString(s string) (UInt160, error)
```

#### 3.2.6 Create a new UInt256 object from a byte array

```golang
func UInt256FromBytes(b []byte) (u UInt160, err error)
```

#### 3.2.7 Convert a UInt256 object to a big-endian hex string

```golang
func (u UInt256) String() string
```

#### 3.2.8 Convert a UInt256 object to a littel-endian byte array

```golang
func (u UInt256) ToByteArray() []byte
```

#### 3.2.9 Concatenate two byte arrays

```golang
func ConcatBytes(b1 []byte, b2 []byte) []byte
```

#### 3.2.10 Convert a byte array to a hex string

```golang
func BytesToHex(b []byte) string
```

#### 3.2.11 Convert a hex string to a byte array

```golang
func HexTobytes(hexstring string) (b []byte)
```

#### 3.2.12 Reverse a byte array

```golang
func ReverseBytes(param []byte) []byte
```

#### 3.2.13 Convert an address string to a script hash

```golang
func AddressToScriptHash(address string) (UInt160, error)
```

#### 3.2.14 Convert a script hash to an address string

```golang
func ScriptHashToAddress(scriptHash UInt160) string
```

#### 3.2.15 Reverse a string

```golang
func ReverseString(input string) string
```

#### 3.2.16 Generate a random byte array

```golang
func GenerateRandomBytes(size int) ([]byte, error)
```

*Typical usage:*

```golang
package sample

import "encoding/hex"
import "github.com/joeqian10/neo3-gogogo/helper"

func SampleMethod() {
    // UInt160
    hexStr := "0x2d3b96ae1bcc5a585e075e3b81920210dec16302"
    v1, err := helper.UInt160FromString(hexStr)
    b1, err := hex.DecodeString(hexStr)
    v2, err := helper.UInt160FromBytes(ReverseBytes(b))
    s1 := v1.String()
    ba1 := v2.ToByteArray()
    // v1 and v2 are equal

    // UInt256
    str := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f32d"
    u1, err := helper.UInt256FromString(str)
    b2, err := hex.DecodeString(hexStr)
    u2, err := helper.UInt256FromBytes(ReverseBytes(b))
    s2 := u1.String()
    ba2 := u2.ToByteArray()
    // u1 and u2 are equal

    // reverse bs
    b3 := []byte{1, 2, 3}
    r := helper.ReverseBytes(b3)

    // concatenate bs
    b4 := []byte{4, 5, 6}
    c := helper.ConcatBytes(b3, b4)

    // convert byte array to hex string
    s := helper.BytesToHex(b3)

    // convert hex string to byte array
    b5 := helper.HexToBytes(s)

    // convert ScriptHash to address string
    a := helper.ScriptHashToAddress(v1)

    // convert address string to ScriptHash
    v3, err := helper.AddressToScriptHash(a)

    // reverse a string
    s3 := helper.ReverseString(s)

    // generate random bs
    ba, err := helper.GenerateRandomBytes(10)
    ...
}
```

### 3.3 "rpc" module

This module provides structs and methods which can be used to send RPC requests to and receive RPC responses from a neo node. For more information about neo RPC API, refer to [API Reference]().

#### 3.3.1 Create a new RPC client

```golang
func NewClient(endpoint string) *RpcClient
```

`endPoint` can be the RPC port of a MainNet, TestNet or a LocalNet neo node.

#### 3.3.2 Get current block count

```golang
func (n *RpcClient) GetBlockCount() GetBlockCountResponse
```

#### 3.3.3 Get block information

```golang
func (n *RpcClient) GetBlock(hashOrIndex string) GetBlockResponse
```

#### 3.3.4 Get smart contract execution results

```golang
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse
```

.....  

There are more RPC APIs and they will not be all listed in this document. Please find what you need from the source code.

*Typical usage:*

```golang

package sample

import "github.com/joeqian10/neo3-gogogo/rpc"

func SampleMethod() {
    // create a rpc client
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    client := rpc.NewClient(TestNetEndPoint)

    // get block count
    r1 := client.GetBlockCount()
    height := r1.Result

    // get raw mempool, get all the transactions' id in this node's mempool
    r2 := client.GetRawMemPool()
    transactions := r2.Result

    // get transaction detail by its id
    r3 := client.GetRawTransaction("your transaction id string")
    tx := r3.Result

    // send raw transaction
    r4 := client.SendRawTransaction("raw transaction hex string")

    ...
}

```

### 3.4 "sc" module

This module is mainly used to build smart contract scripts which can be run in a neo virtual machine. For more information about neo smart contract and virtual machine, refer to [NeoContract]() and [NeoVM]().

#### 3.4.1 Create a new ScriptBuilder object

```golang
func NewScriptBuilder() ScriptBuilder
```

#### 3.4.2 Converts a ScriptBuilder to a byte array

```golang
func (sb *ScriptBuilder) ToArray() []byte
```

#### 3.4.3 Emit an operation code to the script

```golang
func (sb *ScriptBuilder) Emit(op OpCode, arg ...byte) error
```

#### 3.4.4 Emit an operation code to call a smart contract

```golang
func (sb *ScriptBuilder) EmitAppCall(scriptHash helper.UInt160, operation string, args []ContractParameter) error
```

#### 3.4.5 Emit an operation code to call a method at a specific position in the script

```golang
func (sb *ScriptBuilder) EmitCall(offset int) error
```

#### 3.4.6 Emit an operation code to jump to another position in the script

```golang
func (sb *ScriptBuilder) EmitJump(op OpCode, offset int) error
```

#### 3.4.7 Emit an operation code to push a BigInteger

```golang
func (sb *ScriptBuilder) EmitPushBigInt(number big.Int) error
```

#### 3.4.8 Emit an operation code to push an integer

```golang
func (sb *ScriptBuilder) EmitPushInt(number int) error
```

#### 3.4.9 Emit an operation code to push a boolean value

```golang
func (sb *ScriptBuilder) EmitPushBool(param bool) error
```

#### 3.4.10 Emit an operation code to push a byte array

```golang
func (sb *ScriptBuilder) EmitPushBytes(param []byte) error
```

#### 3.4.11 Emit an operation code to push a string

```golang
func (sb *ScriptBuilder) EmitPushString(param string) error
```

#### 3.4.12 Emit an operation code to push a raw byte array

```golang
func (sb *ScriptBuilder) EmitRaw(arg []byte) error
```

#### 3.4.13 Emit an operation code to push a ContractParameter

```golang
func (sb *ScriptBuilder) EmitPushParameter(param ContractParameter) error
```

#### 3.4.12 Emit an operation code to call a system method

```golang
func (sb *ScriptBuilder) EmitSysCall(api uint) error
```

*Typical usage:*

```golang

package sample

import "github.com/joeqian10/neo-gogogo/sc"

func SampleMethod() {
    // create a script builder
    sb := sc.NewScriptBuilder()

    // call a specific method from a specific contract
    scriptHash, _ := helper.UInt160FromString("b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    sb.EmitAppCall(scriptHash.ToByteArray(), "name", nil)
    bs := sb.ToArray()

    ...
}

```

### 3.5 "tx" module

This module defines the transaction and the parts which make up a transaction in the neo network, and also provides structs and methods for building transactions from scratch. For more information about neo transactions, refer to [Transaction]().

#### 3.5.1 Create a TransactionBuilder

```golang
func NewTransactionBuilder(endPoint string) *TransactionBuilder
```

#### 3.5.2 Create a TransactionBuilder from a rpc client

```golang
func NewTransactionBuilderFromClient(client rpc.IRpcClient) *TransactionBuilder
```

#### 3.5.3 Make an unsigned transaction

```golang
func (tb *TransactionBuilder) MakeTransaction(script []byte, sender helper.UInt160, attributes []*TransactionAttribute, cosigners []*Signer) (*Transaction, error)
```

#### 3.5.4 Add a single signed signature to TransactionBuild's SignStore

```golang
func (tb *TransactionBuilder) AddSignature(keyPair *keys.KeyPair) error
```

#### 3.5.5 Add a multi signed signature to TransactionBuild's SignStore

```golang
func (tb *TransactionBuilder) AddMultiSig(keyPairs []*keys.KeyPair, m int, pubKeys1 []*keys.PublicKey) error
```

#### 3.5.6 Add a SignItem to TransactionBuilder's SignStore

```golang
func (tb *TransactionBuilder) AddSignItem(contract *sc.Contract, keyPair *keys.KeyPair) error
```

#### 3.5.7 Sign a transaction

```golang
func (tb *TransactionBuilder) Sign() error
```

#### 3.5.8 Calculate network fee with TransactionBuilder's SignStore

```golang
func (tb *TransactionBuilder) CalculateNetworkFeeWithSignStore() int64
```

#### 3.5.9 Calculate network fee of a specific script

```golang
func (tb *TransactionBuilder) CalculateNetWorkFee(witness_script []byte, size *int) int64
```

#### 3.5.10 Get current blockchain height

```golang
func (tb *TransactionBuilder) GetBlockHeight() (uint32, error)
```

#### 3.5.11 Get the balance of an asset of an account

```golang
func (tb *TransactionBuilder) GetBalance(account helper.UInt160, assetId helper.UInt160) (int64, error)
```

#### 3.5.12 Get the gas consumed when running a script in ApplicationEngine in test mode

```golang
func (tb *TransactionBuilder) GetGasConsumed(script []byte) (int64, error)
```

#### 3.5.13 Get the script of a contract via its scriptHash

```golang
func (tb *TransactionBuilder) GetWitnessScript(hash helper.UInt160) ([]byte, error)
```

Typical usage:

```golang

package sample

import "github.com/joeqian10/neo3-gogogo/tx"
import "github.com/joeqian10/neo3-gogogo/sc"

func SampleMethod() {
    // create a transaction builder
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    tb := tx.NewTransactionBuilder(TestNetEndPoint)

    sb := sc.NewScriptBuilder()
    scriptHash, _ := helper.UInt160FromString("b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    sb.EmitAppCall(scriptHash.ToByteArray(), "name", nil)
    bs := sb.ToArray()

    // build a transaction
    sender, _ := helper.AddressToScriptHash("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")

    t, _ := tb.MakeTransaction(bs, sender, nil, nil)
    // get the raw byte array of this transaction
    unsignedRaw := t.UnsignedRawTransaction()

    // sign the transaction
    tb.Tx = t
    keyPair, _ := keys.NewKeyPairFromWIF("L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp")
    _ = tb.AddSignature(keyPair)
    err := tb.Sign()

    // refer to unit tests for more examples
    ......
}

```

### 3.6 "wallet" module

This module defines the account and the wallet in the neo network, and methods for creating an account or a wallet, signing a message/verifying signature with private/public key pair are also provided. For more information about the neo wallet, refer to [NEP6Wallet]().

#### 3.6.1 Create a new account

```golang
func NewAccount() (*NEP6Account, error)
```

#### 3.6.2 Create a new account from a WIF key

```golang
func NewAccountFromWIF(wif string) (*NEP6Account, error)
```

#### 3.6.3 Create a new account from an encrypted key according to NEP-2

```golang
func NewAccountFromNEP2(nep2Key, passphrase string) (*NEP6Account, error)
```

#### 3.6.4 Create a new wallet

```golang
func NewWallet() *NEP6Wallet
```

#### 3.6.5 Add a new account into the wallet

```golang
func (w *NEP6Wallet) AddNewAccount() error
```

#### 3.6.6 Import an account from a WIF key

```golang
func (w *NEP6Wallet) ImportFromWIF(wif string) error
```

#### 3.6.7 Import an account from an encrypted key according to NEP-2

```golang
func (w *NEP6Wallet) ImportFromNEP2Key(nep2Key, passphare string) error
```

#### 3.6.8 Add an existing account into the wallet

```golang
func (w *NEP6Wallet) AddAccount(acc *NEP6Account)
```

#### 3.6.9 Create a WalletHelper

```golang
func NewWalletHelper(rpc rpc.IRpcClient, account *NEP6Account) *WalletHelper
```

#### 3.6.10 Transfer assets

```golang
func (w *WalletHelper) Transfer(assetId helper.UInt160, to string, amount *big.Int) (string, error)
```

#### 3.6.11 Claim gas

```golang
func (w *WalletHelper) ClaimGas() (string, error)
```

#### 3.6.12 Get asset balance

```golang
func (w *WalletHelper) GetBalance(assetId helper.UInt160, address string) (balance *big.Int, err error)
```

#### 3.6.13 Get unclaimed gas

```golang
func (w *WalletHelper) GetUnClaimedGas(address string) (float64, error)
```

#### 3.6.14 Generate a random key pair

```golang
func GenerateKeyPair() (key *KeyPair, err error)
```

#### 3.6.15 Create a new key pair from private key

```golang
func NewKeyPair(privateKey []byte) (key *KeyPair, err error)
```

#### 3.6.16 Create a new key pair from WIF

```golang
func NewKeyPairFromWIF(wif string) (key *KeyPair, err error)
```

#### 3.6.17 Create a new key pair from an encrypted key according to NEP-2

```golang
func NewKeyPairFromNEP2(nep2 string, password string) (key *KeyPair, err error)
```

#### 3.6.18 Use a key pair to sign arbitrary message

```golang
func (p *KeyPair) Sign(message []byte) ([]byte, error)
```

*Typical usage:*

```golang

package sample

import "github.com/joeqian10/neo3-gogogo/rpc"
import "github.com/joeqian10/neo3-gogogo/wallet"
import "github.com/joeqian10/neo3-gogogo/wallet/keys"

func SampleMethod() {
    // create an account with a random generated private key
    a1, err := wallet.NewAccount()
    // or create an account with your own private key in WIF format
    a2, err := wallet.NewAccountFromWIF("your private key in WIF format")
    // or create an account with a private key encrypted in NEP-2 standard and a passphrase
    a3, err := wallet.NewAccountFromNEP2("your private key encrypted in NEP-2 standard", "your passphrase")

    // create a new wallet
    w := wallet.NewWallet()
    // add a new account into the wallet
    w.AddNewAccount()
    // or import an account from a WIF key
    w.ImportFromWIF("your account private key in WIF format")
    // or import an account from a private key encrypted in NEP-2 standard and a passphrase
    w.ImportFromNep2Key("your account private key encrypted in NEP-2 standard", "your account passphrase")
    // or simply add an existing account
    w.AddAccount(a1)

    // create a WalletHelper
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    client := rpc.NewClient(TestNetEndPoint)
    wh := wallet.NewWalletHelper(client, a2)
    // transfer some neo
    wh.Transfer(tx.NeoToken, a2.Address, a3.Address, 80000)
    // claim gas
    wh.ClaimGas(a2.Address)

    // create a KeyPair
    kp1, err := keys.GenerateKeypair()
    kp2, err := keys.NewKeyPair(helper.HexToBytes("831cb932167332a768f1c898d2cf4586a14aa606b7f078eba028c849c306cce6"))
    kp3, err := keys.NewKeyPairFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
    kp4, err := keys.NewKeyPairFromNEP2("6PYMDM4ZwugPFbFqtDwen33yEYtnSSc3fdR8z3iseTH3FeuGxSc4GDcSMj", "neo3-gogogo")

    // sign arbitrary message
    b, err := kp1.Sign([]byte{})

    ......
}

```

### 3.7 "nep5" module

This module is to make life easier when dealing with NEP-5 tokens. Methods for querying basic information of a NEP-5 token, such as name, total supply, are provided. Also, it offers the ability to test run the scripts to transfer and get the balance of a NEP-5 token. For more information about NEP-5, refer to [NEP-5]().

#### 3.7.1 Create a new Nep17Helper with token script hash and a rpc client

```golang
func NewNep5Helper(scriptHash helper.UInt160, client rpc.IRpcClient) *Nep17Helper
```

#### 3.7.2 Get the total supply of a NEP-5 token

```golang
func (n *Nep17Helper) TotalSupply() (*big.Int, error)
```

#### 3.7.3 Get the name of a NEP-5 token

```golang
func (n *Nep17Helper) Name() (string, error)
```

#### 3.7.4 Get the symbol of a NEP-5 token

```golang
func (n *Nep17Helper) Symbol() (string, error)
```

#### 3.7.5 Get the decimals of a NEP-5 token

```golang
func (n *Nep17Helper) Decimals() (int, error)
```

#### 3.7.6 Get the balance of a NEP-5 token of an account

```golang
func (n *Nep17Helper) BalanceOf(account helper.UInt160) (*big.Int, error)
```

#### 3.7.7 Create a transaction to transfer NEP-5 token

```golang
func (n *Nep17Helper) CreateTransferTx(fromKey *keys.KeyPair, to helper.UInt160, amount *big.Int) (*tx.Transaction, error)
```

*Typical usage:*

```golang

package sample

import "github.com/joeqian10/neo-gogogo/nep5"
import "github.com/joeqian10/neo-gogogo/wallet"

func SampleMethod() {
    // create a Nep17Helper
    scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    var testNetEndPoint = "http://seed1.ngd.network:20332"
    nh := nep5.NewNep5Helper(scriptHash, testNetEndPoint)

    // get the name of a NEP-5 token
    name, err := nh.Name()

    // get the total supply of a NEP-5 token
    s, e := nh.TotalSupply()

    // get the balance of a NEP-5 token of an address
    address, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
    u, e := nh.BalanceOf(address)

    // test run the script for transfer a NEP-5 token
    from, err := keys.NewKeyPairFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
    to, _ := helper.AddressToScriptHash("AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV")
    t, e := nh.CreateTransferTx(from, to, 1)

    ......
}

```

## 4. Contributing

Any help is welcome! Please sign off your commits and pull requests, and add proper comments.

### 4.1 Lisense

This project is licensed under the MIT License.
