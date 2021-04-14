package keys

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/text/unicode/norm"
	"math/big"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
)

type KeyPairSlice []KeyPair

func (kps KeyPairSlice) Len() int           { return len(kps) }
func (kps KeyPairSlice) Less(i, j int) bool { return kps[i].PublicKey.CompareTo(kps[j].PublicKey) == -1 }
func (kps KeyPairSlice) Swap(i, j int)      { kps[i], kps[j] = kps[j], kps[i] }

type KeyPair struct {
	PrivateKey []byte
	PublicKey  *crypto.ECPoint
}

const (
	N = 16384
	R = 8
	P = 8
)

func NewKeyPair(privateKey []byte) (*KeyPair, error) {
	length := len(privateKey)
	if length != 32 {
		return nil, fmt.Errorf("argument length is wrong %v", length)
	}
	ecdsaKey := ToECDsa(privateKey)
	pubKey, err := crypto.CreateECPoint(ecdsaKey.X, ecdsaKey.Y, &ecdsaKey.Curve)
	if err != nil {
		return nil, err
	}
	key := &KeyPair{privateKey, pubKey}
	return key, nil
}

func NewKeyPairFromWIF(wif string) (*KeyPair, error) {
	if wif == "" {
		return nil, fmt.Errorf("wif string is empty")
	}
	data, err := crypto.Base58CheckDecode(wif)
	if err != nil {
		return nil, err
	}
	length := len(data)
	if length != 34 || data[0] != 0x80 || data[33] != 0x01 {
		return nil, fmt.Errorf("invalid parameter format")
	}
	privateKey := data[1:33]
	return NewKeyPair(privateKey)
}

func NewKeyPairFromNEP2(nep2 string, passphrase string, version byte, N, R, P int) (*KeyPair, error) {
	if nep2 == "" {
		return nil, fmt.Errorf("NEP2 string is empty")
	}
	data, err := crypto.Base58CheckDecode(nep2)
	if err != nil {
		return nil, err
	}
	if len(data) != 39 || data[0] != 0x01 || data[1] != 0x42 || data[2] != 0xe0 {
		return nil, fmt.Errorf("format error: invalid nep2 string")
	}
	addressHash := make([]byte, 4)
	copy(addressHash, data[3:7])
	dataPassphrase := norm.NFC.Bytes([]byte(passphrase)) // Normalize the passphrase according to the NFC standard.
	derivedKey, err := scrypt.Key(dataPassphrase, addressHash, N, R, P, 64)
	if err != nil {
		return nil, err
	}

	derivedKey1 := derivedKey[:32]
	derivedKey2 := derivedKey[32:]
	encryptedKey := make([]byte, 32)
	copy(encryptedKey, data[7:])
	decrypted, err := crypto.AESDecrypt(encryptedKey, derivedKey2)
	if err != nil {
		return nil, err
	}

	priKey := helper.XOR(decrypted, derivedKey1)
	pair, err := NewKeyPair(priKey)
	if err != nil {
		return nil, err
	}

	address := PublicKeyToAddress(pair.PublicKey, version)
	hash := crypto.Hash256([]byte(address))[:4]
	if !bytes.Equal(addressHash, hash) {
		return nil, fmt.Errorf("format error: address hash not equal")
	}
	return pair, nil
}

func GenerateKeyPair() (*KeyPair, error) {
	ecdsaKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	privateKey := ecdsaKey.D.Bytes()
	return NewKeyPair(privateKey)
}

func (p *KeyPair) CompareTo(q *KeyPair) int {
	if p == q {
		return 0
	}
	return p.PublicKey.CompareTo(q.PublicKey)
}

// ecdsa converts the key to a usable ecdsa.PrivateKey for signing data.
func (p *KeyPair) ToECDsa() *ecdsa.PrivateKey {
	return ToECDsa(p.PrivateKey)
}

// ecdsa converts the private key byte[] to a usable ecdsa.PrivateKey for signing data.
func ToECDsa(key []byte) *ecdsa.PrivateKey {
	ecdsaKey := new(ecdsa.PrivateKey)
	ecdsaKey.PublicKey.Curve = elliptic.P256()
	ecdsaKey.D = new(big.Int).SetBytes(key)
	ecdsaKey.PublicKey.X, ecdsaKey.PublicKey.Y = ecdsaKey.PublicKey.Curve.ScalarBaseMult(key)
	return ecdsaKey
}

// export wif string
func (p *KeyPair) Export() string {
	data := make([]byte, 34)
	data[0] = 0x80
	copy(data[1:], p.PrivateKey)
	data[33] = 0x01
	wif := crypto.Base58CheckEncode(data)
	return wif
}

// export nep2 key string
func (p *KeyPair) ExportWithPassword(password string, version byte, N, R, P int) (string, error) {
	s := ""
	address := PublicKeyToAddress(p.PublicKey, version)
	addressHash := crypto.Hash256([]byte(address))[:4]
	// Normalize the passphrase according to the NFC standard.
	phraseNorm := norm.NFC.Bytes([]byte(password))
	derivedKey, err := scrypt.Key(phraseNorm, addressHash, N, R, P, 64)
	if err != nil {
		return s, err
	}
	derivedKey1 := derivedKey[:32]
	derivedKey2 := derivedKey[32:]
	xr := helper.XOR(p.PrivateKey, derivedKey1)
	encryptedKey, err := crypto.AESEncrypt(xr, derivedKey2)
	if err != nil {
		return s, err
	}
	buffer := make([]byte, 39)
	buffer[0] = 0x01
	buffer[1] = 0x42
	buffer[2] = 0xe0
	copy(buffer[3:7], addressHash)
	copy(buffer[7:], encryptedKey)
	return crypto.Base58CheckEncode(buffer), nil
}

// String implements the Stringer interface.
func (p *KeyPair) String() string {
	return helper.BytesToHex(p.PrivateKey)
}

// sign message with KeyPair
func (p *KeyPair) Sign(message []byte) ([]byte, error) {
	privateKey := p.ToECDsa()
	hash := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])

	if err != nil {
		return nil, err
	}

	params := privateKey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// ExistsIn returns true if p is in list
func (p *KeyPair) ExistsIn(list []KeyPair) bool {
	for _, item := range list {
		if item.CompareTo(p) == 0 {
			return true
		}
	}
	return false
}

// Verify returns true if the signature is valid and corresponds
// to the hash and public key
func VerifySignature(message []byte, signature []byte, p *crypto.ECPoint) bool {
	hash := sha256.Sum256(message)
	publicKey := p.ToECDsa()

	if p.X == nil || p.Y == nil {
		return false
	}
	rBytes := new(big.Int).SetBytes(signature[0:32])
	sBytes := new(big.Int).SetBytes(signature[32:64])
	return ecdsa.Verify(publicKey, hash[:], rBytes, sBytes)
}

func VerifyMultiSig(message []byte, signatures [][]byte, pubKeys []crypto.ECPoint) bool {
	m := len(signatures)
	n := len(pubKeys)
	if m==0 || n==0 || m>n {return false}
	var success bool = true
	for i, j := 0, 0; success && i < m && j < n; {
		if VerifySignature(message, signatures[i], &pubKeys[j]) {i++}
		j++
		if m-i > n-j {success=false}
	}
	return success
}
