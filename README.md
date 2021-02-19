# neo3-gogogo

## 1. Overview

This is a light-weight golang SDK for Neo 3.0 network.

### 1.1 Version

neo: v3.0.0-preview5

## 2. Getting Started

### 2.1 Installation

Simply add this SDK to GOPATH:

```bash
go get github.com/joeqian10/neo3-gogogo
```

## 3. Modules

### 3.1 "block" module

This module defines the blockhead and block struct along with their functions used in neo, such as serializing/deserializing, hashing of a blockhead.

### 3.2 "blockchain" module

This module defines the storage key/value struct in neo.

### 3.3 "crypto" module

This module offers methods used for cryptography purposes, such as AES encryption/decryption, Base58/Base64 encoding/decoding, Hash160/Hash256 hashing functions, eliptic curve points (public key) related functions, and script hash/address conversion functions. For more information about the crypto algorithms used in neo, refer to [Cryptography](https://github.com/neo-project/neo/tree/master/src/neo/Cryptography).

*Typical usage:*

```golang
package sample

import "crypto/elliptic"
import "encoding/hex"
import "github.com/joeqian10/neo3-gogogo/crypto"
import "math/big"

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

    // Base64 encoding/decoding
    b64encoded := crypto.Base64Encode([]byte{0x01, 0x02})
    b64decoded, err := crypto.Base64Decode("DCEC6W884XWN8uDUF64bnkk64et86LWWDjdHd+AZQ+2vyC0LQZVEDXg=")

    // Sha256, Hash256, Hash160
    b := []byte("Hello World")
    s1 := crypto.Sha256(b)
    s2 := crypto.Hash256(b)
    s3 := crypto.Hash160(b)

    // ec point related
    var p256 = elliptic.P256()
    point, err := crypto.CreateECPoint(big.NewInt(100), big.NewInt(200), &p256)
    point, err := crypto.NewECPointFromBytes([]byte{})
    point, err := crypto.NewECPointFromString("")
    point, err := crypto.DecodePoint([]byte{}, &p256)
    point, err := crypto.FromBytes([]byte{}, &p256)

    ...

    // script hash/address conversion
    scriptHash, err := crypto.AddressToScriptHash("NdtB8RXRmJ7Nhw1FPTm7E6HoDZGnDw37nf")
    address := ScriptHashToAddress(helper.UInt160FromBytes(Hash160([]byte{0x01})))

    ...
}
```

### 3.4 "helper" module

As its name indicated, this module acts as a helper and provides some standard param types used in neo, such as `UInt160`, `UInt256`, and some auxiliary methods with basic functionalities including conversion between a hex string and a byte array, conversion between a script hash and a standard neo address, concatenating/reversing byte arrays and so on.

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

    // reverse bytes
    b3 := []byte{1, 2, 3}
    r := helper.ReverseBytes(b3)

    // concatenate bytes
    b4 := []byte{4, 5, 6}
    c := helper.ConcatBytes(b3, b4)

    // convert byte array to hex string
    s := helper.BytesToHex(b3)

    // convert hex string to byte array
    b5 := helper.HexToBytes(s)

    // reverse a string
    s3 := helper.ReverseString(s)

    // integer types and byte array conversion
    bs := helper.UInt64ToBytes(uint64(12345678))
    n := helper.BytesToUInt64([]byte{})

    // generate random bytes
    bs, err := helper.GenerateRandomBytes(10)

    // big.Int and byte array in neo conversion
    neoBytes := helper.BigIntToNeoBytes(big.NewInt(12345678))
    bigInteger := helper.BigIntFromNeoBytes([]byte{})

    ...
}
```

### 3.5 "io" module

This module provides structs, methods and interfaces for serializing/deserializing operations in neo.

*Typical usage:*

```golang
package sample

import "encoding/hex"
import "github.com/joeqian10/neo3-gogogo/io"

func SampleMethod() {
    // create a BinaryReader
    b := make([]byte, 4)
    br := io.NewBinaryReaderFromBuf(b)
    
    // use the BinaryReader to read something
    var result1 uint32
    br.ReadLE(&result1)
    result2 := br.ReadVarUInt() // result2 is uint64

    ...

    // create a BinaryWriter
    b := new(bytes.Buffer)
    bw := io.NewBinaryWriterFromIO(b)
    // or from a BufBinaryWriter
    bbw := io.NewBufBinaryWriter()
    bw := bbw.BinaryWriter

    // use the BinaryWriter to write something
    bw.WriteBE(result1)
    bw.WriteVarBytes([]byte{})

    ...
}
```

### 3.6 "keys" module

This module defines the KeyPair struct which is a wrapper of a pair of private key and public key. The KeyPair can be used to sign messages.

*Typical usage:*

```golang
package sample

import "encoding/hex"
import "github.com/joeqian10/neo3-gogogo/keys"

func SampleMethod() {
    // create a KeyPair
    privateKey := make([]byte, 32)
    pair, err := keys.NewKeyPair(privateKey)
    pair, err := keys.NewKeyPairFromNEP2("6PYX7SaH7EdDgeHo1V8yXfhrgrGXjHaMzsgWLMpaxLkDR9u6HHnEDMmjhh", "neo3-gogogo", 16384, 8, 8)
    pair, err := keys.NewKeyPairFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")

    // export a KeyPair
    wif := pair.Export()
    nep2 := pair.ExportWithPassword("neo3-gogogo", 16384, 8, 8)

    // sign messages
    data := []byte("hello world")
    signature, err := pair.Sign(data)

    // verify signature
    valid := keys.VerifySignature(data, signature, pair.PublicKey)

    ...
}
```

### 3.7 "mpt" module

This module provides structs and methods used to interact with the StateRoot in neo. For more information about the state service in neo, refer to [StateService](https://github.com/neo-project/neo-modules/tree/master/src/StateService).

*Typical usage:*

```golang
package sample

import "github.com/joeqian10/neo3-gogogo/helper"
import "github.com/joeqian10/neo3-gogogo/mpt"

func SampleMethod() {
    // resolve proof
    proofData := helper.HexToBytes("36a5f9b4f70c841154060b132ff1e253eecaa7db8a61737365740034c1e9f7413e8a5eb3c500ffd06acd41644d5d5a96000000000000060772000000000000000000000020dbe90d0674546fd0e6dc879013ab743b5f171cc1f4e97caab98be714bb7f125400208837fe543b1bfcd58d480ea9988cb7acf58c5f4cf0518a987f70924320ceab390020ff409ca133d8a3a5ea8b72d5b82214d6190d3dd44d19d67062681af2b4b8302000004b0128050f090b040f07000c0804010105040006000b0103020f0f010e0205030e0e0c0a0a070d0b080a06202da36d644929144a9869d0ffb581c645f4d29b7b38a068ca1cc3bbd34b0771e252000020a41c8cc1e18bcbe0f237fa00544455ab92d5a86bc8fa6cd0f0ba294e9f02202d002002ad9e7adebf3057801f6918a97efd49741949f8a538b58b4fae31700b976e7d000000000000000000000000002d010a0703070306050704000020afb2d68b4ae87095c68d113f42234c3e36e15597ec1c11ecce4261f3a9169b0a5200207c00cb1580a95b55ecf7e82829997ebf4176cf95df65712a0b14b37eb13c1c5f0000207e9c938de1d3f1f95753c0867e3326bdb16b2caccefb6f490810bf13ccc7440e000000000000000000000000005a0137040c010e090f070401030e080a050e0b030c0500000f0f0d00060a0c0d04010604040d050d050a090600000000000000000000000000062046a12d0bfc2f3f1d9e18a51f55d32ba4855578c1954d277180addb5b69210c970903070004c071b50400")
    root, _ := helper.UInt256FromString("34db5a993a95e0db79efe8220bf142e5952056bb59834fe3b91fc1611ed4385e")

    scriptHash, key, proofs, err := ResolveProof(proofdata)

    // verify proof
    value, err := VerifyProof(root.ToByteArray(), scriptHash, key, proofs)

    ...
}
```

### 3.8 "nep17" module

This module provides wrapper structs and methods for NEP17 tokens. For more information about NEP17, refer to [NEP17](https://github.com/neo-project/proposals/tree/nep-17).

*Typical usage:*

```golang
package sample

import "github.com/joeqian10/neo3-gogogo/helper"
import "github.com/joeqian10/neo3-gogogo/nep17"
import "github.com/joeqian10/neo3-gogogo/rpc"

func SampleMethod() {
    // create a new Nep17Helper
    scriptHash := helper.NewUInt160FromBytes([]byte{})
    client := rpc.NewClient("http://seed1.ngd.network:20332")
    nep := nep17.NewNep17Helper(scriptHash, client)

    // get symbol
    symbol, err := nep.Symbol()

    ...
}
```

### 3.9 "rpc" module

This module provides structs and methods which can be used to send RPC requests to and receive RPC responses from a neo node. For more information about neo RPC API, refer to [APIs](https://docs.neo.org/v3/docs/en-us/reference/rpc/latest-version/api.html).

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

### 3.10 "sc" module

This module is mainly used to build smart contract scripts which can be run in a neo virtual machine. For more information about neo smart contract and virtual machine, refer to [NeoContract](https://docs.neo.org/v3/docs/en-us/develop/write/basics.html) and [NeoVM](https://docs.neo.org/v3/docs/en-us/basic/neovm.html).

*Typical usage:*

```golang

package sample

import "math/big"
import "github.com/joeqian10/neo3-gogogo/crypto"
import "github.com/joeqian10/neo3-gogogo/helper"
import "github.com/joeqian10/neo3-gogogo/sc"

func SampleMethod() {
    // create a script builder
    sb := sc.NewScriptBuilder()

    // emit an OpCode
    sb.Emit(NOP)

    // emit push a big integer
    sb.EmitPushBigInt(big.NewInt(-1))

    // emit push a bool value
    sb.EmitPushBool(true)

    // emit push bytes
    sb.EmitPushBytes([]byte{0x01, 0x02})

    // emit push a string
    sb.EmitPushString("hello world")

    // call a specific method from a specific contract without parameters
    scriptHash, _ := helper.UInt160FromString("b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    sb.EmitDynamicCall(scriptHash.ToByteArray(), "name")

    // call a specific method from a specific contract with parameters
    sb.EmitDynamicCallParam(scriptHash.ToByteArray(), "balanceOf", []ContractParameter{{
        Type: Hash160,
        Value: helper.NewUInt160(),
    }}...)

    // create an array
    a := []interface{}{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
    sb.CreateArray(a)

    // create a map
    a := map[interface{}]interface{}{}
    a[big.NewInt(1)] = big.NewInt(2)
    a[big.NewInt(3)] = big.NewInt(4)
    sb.CreateMap(a)

    // get the script
    script, err := sb.ToArray()

    ...

    // create a contract
    c := CreateContract([]ContractParameterType{Signature}, script)

    // create a single signature contract
    pubKey, _ := crypto.NewECPointFromString("03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619")
    c1, err := CreateSignatureContract(pubKey)

    // create a multi signature contract
    pubKey2, _ := crypto.NewECPointFromString("027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b")
    c2, err := CreateMultiSigContract(1, []crypto.ECPoint{*pubKey, *pubKey2})

    ...
}
```

### 3.11 "tx" module

This module defines the transaction and the parts which make up a transaction in the neo network, and also provides structs and methods for building transactions from scratch. For more information about neo transactions, refer to [Transaction](https://docs.neo.org/v3/docs/en-us/basic/concept/transaction.html).

Typical usage:

```golang
package sample

import "github.com/joeqian10/neo3-gogogo/helper"
import "github.com/joeqian10/neo3-gogogo/tx"
import "github.com/joeqian10/neo3-gogogo/sc"

func SampleMethod() {
    // generate a new empty transaction
    trx := NewTransaction()

    // get transaction hash
    hash := trx.GetHash()

    // get raw transaction
    raw := trx.ToByteArray()

    ...

    // create a new signer
    s := NewSigner(helper.NewUInt160(), CalledByEntry)

    // create a witness
    w, err := CreateWitness([]byte{}, []byte{})

    ...
}

```

### 3.12 "wallet" module

This module defines the account and the wallet in the neo network, and methods for creating an account or a wallet, signing a message/verifying signature with private/public key pair are also provided. For more information about the neo wallet, refer to [Wallet](https://docs.neo.org/v3/docs/en-us/basic/concept/wallets.html).

*Typical usage:*

```golang
package sample

import "github.com/joeqian10/neo3-gogogo/keys"
import "github.com/joeqian10/neo3-gogogo/rpc"
import "github.com/joeqian10/neo3-gogogo/sc"
import "github.com/joeqian10/neo3-gogogo/wallet"


func SampleMethod() {
    var password = "Satoshi"
    var privateKey = []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
                            0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
    var pair, _ = keys.NewKeyPair(privateKey)
    var wif = pair.Export()
    var nep2, _ = pair.ExportWithPassword(password, 2, 1, 1)
    var testContract, _ = sc.CreateSignatureContract(pair.PublicKey)
    var testScriptHash = testContract.GetScriptHash()

    // create a new wallet
    testName := "testName"
    w, err := NewNEP6Wallet("", &testName)

    // add password to the wallet
    err = w.Unlock(&password)

    // create an account with a random generated private key
    a1, err := w.CreateAccount()
    // or create an account with your own private key
    a2, err := w.CreateAccountWithPrivateKey(privateKey)
    // or create an account with contract and key pair
    a3, err := w.CreateAccountWithContract(testContract, pair)
    // or create an account with the script hash
    a4, err := w.CreateAccountWithScriptHash(testScriptHash)

    // import an account from WIF string
    a5, err := w.ImportFromWIF(wif)
    // import an account from NEP2 key and password
    a6, err := w.ImportFromNEP2(nep2, password, 2, 1, 1)

    // delete an account
    w.DeleteAccount(testScriptHash)

    // check if an account exists
    contained := w.Contains(testScriptHash)

    // get a single account by its script hash
    acc := w.GetAccountByScriptHash(testScriptHash)

    // get all accounts
    accounts := w.GetAccounts()

    // lock and unlock
    w.Lock()
    err = w.Unlock(&password)

    // verify password
    verified := w.VerifyPassword(&password)

    // decrypt key
    k, err := w.DecryptKey(nep2)

    ...

    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    client := rpc.NewClient(TestNetEndPoint)

    // create a WalletHelper from private key
    wh1, err := NewWalletHelperFromPrivateKey(client, privateKey)

    // create a WalletHelper from contract and key pair
    wh2, err := NewWalletHelperFromContract(client, testContract, pair)

    // create a WalletHelper from NEP2
    wh3, err := NewWalletHelperFromNEP2(client, nep2, password, 2, 1, 1)

    // create a WalletHelper from WIF
    wh4, err := NewWalletHelperFromWIF(client, wif)

    // create a WalletHelper from an existing wallet
    wh := NewWalletHelperFromWallet(client, w)

    // make transaction
    script := []byte{}
    ab := []AccountAndBalance{
        {
            Account: testScriptHash,
            Value: big.NewInt(1000000000000), // 10000 gas
        },
    }
    trx, err := wh.MakeTransaction(script, nil, nil, ab)

    // sign transaction
    trx, err := wh.SignTransaction(trx)
    
    ...
}
```

## 4. Contributing

Any help is welcome! Please sign off your commits and pull requests, and add proper comments.

### 4.1 Lisense

This project is licensed under the MIT License.
