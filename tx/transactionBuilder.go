package tx

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
)

const NeoTokenId = "9bde8f209c88dd0e7ca3bf0af0f476cdd8207789"
const GasTokenId = "8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b"

const GasFactor = 100000000

var NeoToken, _ = helper.UInt160FromString(NeoTokenId)
var GasToken, _ = helper.UInt160FromString(GasTokenId)

type SignItem struct {
	Hash     helper.UInt160
	Contract *sc.Contract
	KeyPairs []*keys.KeyPair
}

type TransactionBuilder struct {
	Client    rpc.IRpcClient
	SignStore []*SignItem
	Tx        *Transaction
}

func NewTransactionBuilder(endPoint string) *TransactionBuilder {
	client := rpc.NewClient(endPoint)
	if client == nil {
		return nil
	}
	return &TransactionBuilder{
		Client: client,
	}
}

func NewTransactionBuilderFromClient(client rpc.IRpcClient) *TransactionBuilder {
	if client == nil {
		return nil
	}
	return &TransactionBuilder{
		Client: client,
	}
}

func (tb *TransactionBuilder) MakeTransaction(script []byte, sender helper.UInt160, attributes []*TransactionAttribute, cosigners []*Cosigner) (*Transaction, error) {
	rb, err := helper.GenerateRandomBytes(4)
	if err != nil {
		return nil, err
	}
	nonce := binary.LittleEndian.Uint32(rb)
	tb.Tx = new(Transaction)
	// version
	tb.Tx.SetVersion(0)
	// nonce
	tb.Tx.SetNonce(nonce)
	// script
	if script != nil {
		tb.Tx.SetScript(script)
	} else {
		tb.Tx.SetScript([]byte{})
	}
	// sender
	tb.Tx.SetSender(sender)
	// validUntilBlock
	blockHeight, err := tb.GetBlockHeight()
	if err != nil {
		return nil, err
	}
	tb.Tx.SetValidUntilBlock(blockHeight + MaxValidUntilBlockIncrement)
	// attributes
	if attributes != nil {
		tb.Tx.SetAttributes(attributes)
	} else {
		tb.Tx.SetAttributes([]*TransactionAttribute{})
	}
	// cosigners
	if cosigners != nil {
		tb.Tx.SetCosigners(cosigners)
	} else {
		tb.Tx.SetCosigners([]*Cosigner{})
	}
	// sysfee
	gasConsumed, err := tb.GetGasConsumed(script)
	if err != nil {
		return nil, err
	}
	gasConsumed = int64(math.Max(float64(gasConsumed), 0))
	if gasConsumed > 0 {
		remainder := gasConsumed % GasFactor
		if remainder > 0 {
			gasConsumed += GasFactor - remainder
		} else {
			gasConsumed -= remainder
		}
	}
	tb.Tx.SetSystemFee(gasConsumed)
	return tb.Tx, nil
}

func (tb *TransactionBuilder) AddSignature(keyPair *keys.KeyPair) error {
	contract := keys.CreateSignatureContract(keyPair.PublicKey)
	return tb.AddSignItem(contract, keyPair)
}

func (tb *TransactionBuilder) AddMultiSig(keyPairs []*keys.KeyPair, m int, pubKeys []*keys.PublicKey) error {
	contract := keys.CreateMultiSigContract(m, pubKeys)
	for _, key := range keyPairs {
		err := tb.AddSignItem(contract, key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tb *TransactionBuilder) AddSignItem(contract *sc.Contract, keyPair *keys.KeyPair) error {
	// check the contract is needed for signing
	hashes := tb.Tx.GetScriptHashesForVerifying()
	if !contract.GetScriptHash().Exists(hashes) {
		return fmt.Errorf("AddSignItem error: mismatch hash %v", contract.GetScriptHash())
	}

	// add keyPair to existed item
	isInStore := false
	for _, item := range tb.SignStore {
		if item.Hash == contract.GetScriptHash() {
			isInStore = true
			if !keyPair.Exists(item.KeyPairs) {
				item.KeyPairs = append(item.KeyPairs, keyPair)
			}
		}
	}

	//add new SignItem
	if !isInStore {
		tb.SignStore = append(tb.SignStore, &SignItem{
			Hash:     contract.GetScriptHash(),
			Contract: contract,
			KeyPairs: []*keys.KeyPair{keyPair},
		})
	}
	return nil
}

func (tb *TransactionBuilder) Sign() error {
	tb.Tx.SetNetworkFee(tb.CalculateNetworkFeeWithSignStore())
	// get gas balance of sender
	value, err := tb.GetBalance(GasToken, tb.Tx.sender)
	if err != nil {
		return err
	}

	if value.Int64() < tb.Tx.GetNetworkFee()+tb.Tx.GetSystemFee() {
		return fmt.Errorf("insufficient GAS balance")
	}

	// Sign with signStore
	for _, item := range tb.SignStore {
		tb.Tx.AddSignature(item.KeyPairs, item.Contract)
	}

	return nil
}

// GetBlockHeight gets the current blockchain height via rpc
func (tb *TransactionBuilder) GetBlockHeight() (uint32, error) {
	response := tb.Client.GetBlockCount()
	if response.HasError() {
		return 0, fmt.Errorf(response.Error.Message)
	}
	count := uint32(response.Result)
	return count - 1, nil // height = index = count - 1, genesis block is index 0
}

// CalculateNetworkFee
func (tb *TransactionBuilder) CalculateNetworkFeeWithSignStore() int64 {
	var networkFee int64 = 0
	hashes := tb.Tx.GetScriptHashesForVerifying()
	size := tb.Tx.HeaderSize() + TransactionAttributeSlice(tb.Tx.attributes).GetVarSize() + CosignerSlice(tb.Tx.cosigners).GetVarSize() + sc.ByteSlice(tb.Tx.script).GetVarSize() + helper.GetVarSize(len(hashes))
	for _, hash := range hashes {
		witness_script, _ := tb.GetWitnessScript(hash)

		if witness_script == nil {
			continue
		}
		networkFee += tb.CalculateNetworkFee(witness_script, &size)
	}
	networkFee += int64(size) * 1000 // FeePerByte
	return networkFee
}

// CalculateNetworkFee for single witness
func (tb *TransactionBuilder) CalculateNetworkFee(witness_script []byte, size *int) int64 {
	var networkFee int64 = 0
	if sc.ByteSlice(witness_script).IsSignatureContract() {
		*size += 67 + sc.ByteSlice(witness_script).GetVarSize()
		networkFee += sc.OpCodePrices[sc.PUSHDATA1] + sc.OpCodePrices[sc.PUSHDATA1] + sc.OpCodePrices[sc.PUSHNULL] + 1000000 // InteropService.GetPrice(InteropService.Crypto.ECDsaVerify, null)
	} else if b, m, n := sc.ByteSlice(witness_script).IsMultiSigContract(); b {
		size_inv := 66 * m
		*size += helper.GetVarSize(size_inv) + size_inv + sc.ByteSlice(witness_script).GetVarSize()
		networkFee += sc.OpCodePrices[sc.PUSHDATA1] * int64(m)
		sb := sc.NewScriptBuilder()
		sb.EmitPushInt(m)
		networkFee += sc.OpCodePrices[sc.OpCode(sb.ToArray()[0])]
		networkFee += sc.OpCodePrices[sc.PUSHDATA1] * int64(n)
		sb = sc.NewScriptBuilder()
		sb.EmitPushInt(n)
		networkFee += sc.OpCodePrices[sc.OpCode(sb.ToArray()[0])]
		networkFee += sc.OpCodePrices[sc.PUSHNULL] + 1000000*int64(n)
	} else {
		// support more contract types in the future
	}
	return networkFee
}

// GetBalance is used to get balance of an asset of an account
func (tb *TransactionBuilder) GetBalance(assetHash helper.UInt160, account helper.UInt160) (*big.Int, error) {
	sb := sc.NewScriptBuilder()
	cp := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: account.Bytes(),
	}
	sb.EmitAppCall(assetHash, "balanceOf", []sc.ContractParameter{cp})
	script := sb.ToArray()
	response := tb.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return nil, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return nil, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	return stack.ToParameter().Value.(*big.Int), nil
}

// GetGasConsumed runs a script in ApplicationEngine in test mode and returns gas consumed
func (tb *TransactionBuilder) GetGasConsumed(script []byte) (int64, error) {
	response := tb.Client.InvokeScript(helper.BytesToHex(script))
	if response.HasError() {
		return 0, fmt.Errorf(response.Error.Message)
	}
	gasConsumed, err := strconv.ParseInt(response.Result.GasConsumed, 10, 64)
	if err != nil {
		return 0, err
	}
	return gasConsumed, nil
}

// GetWitnessScript is used to get the script of a contract via its scriptHash
func (tb *TransactionBuilder) GetWitnessScript(hash helper.UInt160) ([]byte, error) {
	// try get script from SignStore
	for _, item := range tb.SignStore {
		if item.Hash == hash {
			return item.Contract.Script, nil
		}
	}

	// try get script from on-chain contract
	response := tb.Client.GetContractState(hash.String())
	if response.HasError() {
		return nil, fmt.Errorf(response.Error.Message)
	}
	script, err := base64.StdEncoding.DecodeString(response.Result.Script)
	if err != nil {
		return nil, err
	}
	return script, nil
}
