package wallet

import (
	"encoding/binary"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
	"math"
	"math/big"
	"sort"
	"strconv"
)

type WalletHelper struct {
	Client rpc.IRpcClient
	wallet *NEP6Wallet
	Magic  uint32
}

var dummy = "dummy"

func NewWalletHelperFromPrivateKey(rpc rpc.IRpcClient, priKey []byte) (*WalletHelper, error) {
	dummyWallet, _ := NewNEP6Wallet("", &helper.DefaultProtocolSettings, &dummy, DefaultScryptParameters)
	_ = dummyWallet.Unlock("")
	_, err := dummyWallet.CreateAccountWithPrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	return &WalletHelper{
		Client: rpc,
		wallet: dummyWallet,
	}, err
}

func NewWalletHelperFromContract(rpc rpc.IRpcClient, contract *sc.Contract, pair *keys.KeyPair) (*WalletHelper, error) {
	dummyWallet, _ := NewNEP6Wallet("",&helper.DefaultProtocolSettings, &dummy, DefaultScryptParameters)
	if pair != nil {
		_ = dummyWallet.Unlock("")
	}
	_, err := dummyWallet.CreateAccountWithContract(contract, pair)
	if err != nil {
		return nil, err
	}
	return &WalletHelper{
		Client: rpc,
		wallet: dummyWallet,
	}, err
}

// Create a WalletHelper using your own private key, password is "" by default
func NewWalletHelperFromWIF(rpc rpc.IRpcClient, wif string) (*WalletHelper, error) {
	dummyWallet, _ := NewNEP6Wallet("", &helper.DefaultProtocolSettings, &dummy, DefaultScryptParameters)
	_ = dummyWallet.Unlock("")
	_, err := dummyWallet.ImportFromWIF(wif)
	if err != nil {
		return nil, err
	}
	return &WalletHelper{
		Client: rpc,
		wallet: dummyWallet,
	}, err
}

func NewWalletHelperFromNEP2(rpc rpc.IRpcClient, nep2 string, passphrase string, N, R, P int) (*WalletHelper, error) {
	dummyWallet, _ := NewNEP6Wallet("", &helper.DefaultProtocolSettings, &dummy, NewScryptParameters(N, R, P))
	_, err := dummyWallet.ImportFromNEP2(nep2, passphrase, N, R, P)
	if err != nil {
		return nil, err
	}
	err = dummyWallet.Unlock(passphrase)
	if err != nil {
		return nil, err
	}
	return &WalletHelper{
		Client: rpc,
		wallet: dummyWallet,
	}, err
}

func NewWalletHelperFromWallet(rpc rpc.IRpcClient, wlt *NEP6Wallet) *WalletHelper {
	return &WalletHelper{
		Client: rpc,
		wallet: wlt,
	}
}

func (w *WalletHelper) CalculateNetworkFee(trx *tx.Transaction) (uint64, error) {
	if trx == nil {
		return 0, fmt.Errorf("no transaction to calculate")
	}
	hashes := trx.GetScriptHashesForVerifying()

	// base size for transaction: includes const_header + signers + attributes + script + hashes
	size := trx.HeaderSize() +
		tx.SignerSlice(trx.GetSigners()).GetVarSize() +
		tx.TransactionAttributeSlice(trx.GetAttributes()).GetVarSize() +
		sc.ByteSlice(trx.GetScript()).GetVarSize() +
		helper.GetVarSize(len(hashes))

	exec_fee_factor := int64(30)
	nf := uint64(0)
	for _, hash := range hashes {
		var witness_script []byte
		account := w.wallet.GetAccountByScriptHash(&hash)
		if account != nil {
			c := account.GetContract()
			if c != nil {
				witness_script = c.GetScript()
			}
		}

		if witness_script == nil && trx.GetWitnesses() != nil {
			// try to find the script in the witnesses
			for _, witness := range trx.GetWitnesses() {
				if witness.GetScriptHash().Equals(&hash) {
					witness_script = witness.VerificationScript
					break
				}
			}
		}

		if witness_script == nil {
			// todo, NativeContract case
			continue
		} else if sc.IsSignatureContract(witness_script) {
			size += 67 + sc.ByteSlice(witness_script).GetVarSize()
			nf += uint64(exec_fee_factor * (sc.OpCodePrices[sc.PUSHDATA1] + sc.OpCodePrices[sc.PUSHDATA1] + sc.OpCodePrices[sc.PUSHNULL] + tx.ECDsaVerifyPrice))
		} else if b, m, n, _ := sc.IsMultiSigContract(witness_script); b {
			size_inv := 66 * m
			size += helper.GetVarSize(size_inv) + size_inv + sc.ByteSlice(witness_script).GetVarSize()

			nf += uint64(exec_fee_factor * sc.OpCodePrices[sc.PUSHDATA1] * int64(m))
			sb := sc.NewScriptBuilder()
			sb.EmitPushInteger(m)
			script, _ := sb.ToArray()
			nf += uint64(exec_fee_factor * sc.OpCodePrices[sc.OpCode(script[0])])

			nf += uint64(exec_fee_factor * sc.OpCodePrices[sc.PUSHDATA1] * int64(n))
			sb = sc.NewScriptBuilder()
			sb.EmitPushInteger(n)
			script, _ = sb.ToArray()
			nf += uint64(exec_fee_factor * sc.OpCodePrices[sc.OpCode(script[0])])

			nf += uint64(exec_fee_factor * (sc.OpCodePrices[sc.PUSHNULL] + int64(tx.ECDsaVerifyPrice*n)))
		} else {
			// support more cotnract types in the future
		}
	}
	nf += uint64(size * tx.FeePerByte)
	return nf, nil
}

// ClaimGas for NEP6Account
func (w *WalletHelper) ClaimGas(magic uint32) (string, error) {
	if w.wallet == nil {
		return "", fmt.Errorf("wallet is nil")
	}
	cosigners := make([]tx.Signer, 0)
	sb := sc.NewScriptBuilder()
	for _, account := range w.wallet.accounts {
		neoBalance, err := w.GetBalanceFromAccount(tx.NeoToken, account.scriptHash)
		if err != nil {
			return "", err
		}
		sb.EmitDynamicCall(tx.NeoToken, "transfer", []interface{}{
			sc.ContractParameter{Type: sc.Hash160, Value: account.scriptHash},
			sc.ContractParameter{Type: sc.Hash160, Value: account.scriptHash},
			sc.ContractParameter{Type: sc.Integer, Value: neoBalance},
			sc.ContractParameter{Type: sc.String, Value: ""},
		})
		sb.Emit(sc.ASSERT)
		cosigners = append(cosigners, tx.Signer{
			Account: account.scriptHash,
			Scopes:  tx.CalledByEntry,
		})
	}
	script, err := sb.ToArray()
	if err != nil {
		return "", err
	}
	balancesGas, err := w.GetAccountAndBalance(tx.GasToken)
	if err != nil {
		return "", err
	}
	trx, err := w.MakeTransaction(script, cosigners, []tx.ITransactionAttribute{}, balancesGas)
	if err != nil {
		return "", err
	}
	// sign the tx
	trx, err = w.SignTransaction(trx, magic)
	if err != nil {
		return "", err
	}

	// use RPC to send the tx
	response := w.Client.SendRawTransaction(crypto.Base64Encode(trx.ToByteArray()))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	return response.Result.Hash, nil
}

// GetAccountAndBalance gets account and balance pair
func (w *WalletHelper) GetAccountAndBalance(assetHash *helper.UInt160) ([]AccountAndBalance, error) {
	balances := make([]AccountAndBalance, 0)
	for _, account := range w.wallet.accounts {
		balance, err := w.GetBalanceFromAccount(assetHash, account.scriptHash)
		if err != nil {
			return nil, err
		}
		balances = append(balances, AccountAndBalance{
			Account: account.scriptHash,
			Value:   balance,
		})
	}
	return balances, nil
}

// GetBalanceFromAccount is used to get balance of an asset of an account
func (w *WalletHelper) GetBalanceFromAccount(assetHash *helper.UInt160, account *helper.UInt160) (*big.Int, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitDynamicCall(assetHash, "balanceOf", []interface{}{
		sc.ContractParameter{
			Type:  sc.Hash160,
			Value: account,
		},
	})
	script, err := sb.ToArray()
	if err != nil {
		return nil, err
	}
	response := w.Client.InvokeScript(crypto.Base64Encode(script), nil)
	stack, err := rpc.PopInvokeStack(response)
	if err != nil {
		return nil, err
	}
	r, err := stack.ToParameter()
	if err != nil {
		return nil, err
	}
	return r.Value.(*big.Int), nil
}

// GetBalanceFromWallet is used to get balance from all accounts inside the wallet
func (w *WalletHelper) GetBalanceFromWallet(assetHash *helper.UInt160, wlt *NEP6Wallet) (*big.Int, error) {
	if wlt == nil {
		wlt = w.wallet
	}
	if wlt == nil {
		return nil, fmt.Errorf("wallet is nil")
	}
	r := big.NewInt(0)
	for k, _ := range wlt.accounts {
		b, err := w.GetBalanceFromAccount(assetHash, &k)
		if err != nil {
			return nil, err
		}
		r = r.Add(r, b)
	}
	return r, nil
}

// GetBlockHeight gets the current blockchain height via rpc
func (w *WalletHelper) GetBlockHeight() (uint32, error) {
	response := w.Client.GetBlockCount()
	if response.HasError() {
		return 0, fmt.Errorf(response.Error.Message)
	}
	count := uint32(response.Result)
	return count - 1, nil // height = index = count - 1, genesis block is index 0
}

func (w *WalletHelper) GetContractState(hash *helper.UInt160) (*models.RpcContractState, error) {
	response := w.Client.GetContractState(hash.String())
	if response.HasError() {
		return nil, fmt.Errorf(response.Error.Message)
	}
	return &response.Result, nil
}

// GetGasConsumed runs a script in ApplicationEngine in test mode and returns gas consumed
func (w *WalletHelper) GetGasConsumed(script []byte, signers []models.RpcSigner) (int64, error) {
	fmt.Println(crypto.Base64Encode(script))
	response := w.Client.InvokeScript(crypto.Base64Encode(script), signers)
	if response.HasError() {
		return 0, fmt.Errorf(response.Error.Message)
	}
	gasConsumed, err := strconv.ParseInt(response.Result.GasConsumed, 10, 64)
	if err != nil {
		return 0, err
	}
	return gasConsumed, nil
}

// GetUnClaimedGas gets the amount of unclaimed gas in the wallet
func (w *WalletHelper) GetUnClaimedGas() (uint64, error) {
	if w.wallet == nil {
		return 0, fmt.Errorf("wallet is nil")
	}
	t := uint64(0)
	for _, account := range w.wallet.accounts {
		response := w.Client.GetUnclaimedGas(account.Address)
		if response.HasError() {
			return 0, fmt.Errorf(response.Error.Message)
		}
		u, err := strconv.ParseUint(response.Result.Unclaimed, 10, 64)
		if err != nil {
			return 0, err
		}
		t += u
	}
	return t, nil
}

func (w *WalletHelper) MakeTransaction(script []byte, cosigners []tx.Signer, attributes []tx.ITransactionAttribute, balanceGas []AccountAndBalance) (*tx.Transaction, error) {
	for _, ab := range balanceGas {
		rb, err := helper.GenerateRandomBytes(4)
		if err != nil {
			return nil, err
		}
		nonce := binary.LittleEndian.Uint32(rb)
		trx := new(tx.Transaction)
		// version
		trx.SetVersion(0)
		// nonce
		trx.SetNonce(nonce)
		// script
		trx.SetScript(script)
		// validUntilBlock
		blockHeight, err := w.GetBlockHeight()
		if err != nil {
			return nil, err
		}
		trx.SetValidUntilBlock(blockHeight + tx.MaxValidUntilBlockIncrement)
		// signers
		signers := getSigners(ab.Account, cosigners)
		trx.SetSigners(signers)
		// attributes
		trx.SetAttributes(attributes)
		// sysfee
		gasConsumed, err := w.GetGasConsumed(script, models.CreateRpcSigners(signers))
		if err != nil {
			return nil, err
		}
		gasConsumed = int64(math.Max(float64(gasConsumed), 0))
		trx.SetSystemFee(gasConsumed)
		// netfee
		netFee, err := w.CalculateNetworkFee(trx)
		if err != nil {
			return nil, err
		}
		trx.SetNetworkFee(int64(netFee))

		if ab.Value.Int64() >= trx.GetSystemFee()+trx.GetNetworkFee() {
			return trx, nil
		}
	}
	return nil, fmt.Errorf("insufficient GAS")
}

func (w *WalletHelper) Sign(ctx *ContractParametersContext, magic uint32) (bool, error) {
	fSuccess := false
	for _, scriptHash := range ctx.GetScriptHashes() {
		account := w.wallet.GetAccountByScriptHash(&scriptHash)
		if account != nil {
			// try to sign self-contained multisig
			msc := account.GetContract()
			b := false
			var m int
			var points []crypto.ECPoint
			if msc != nil {
				b, m, points = sc.ByteSlice(msc.GetScript()).IsMultiSigContractWithPoints()
			}
			if msc != nil && b {
				for _, point := range points {
					account = w.wallet.GetAccountByPublicKey(&point)
					if account == nil || account.HasKey() != true {
						continue
					}
					pair, err := account.GetKey()
					if err != nil {
						return false, err
					}
					signature, err := Sign(ctx.Verifiable, pair, magic)
					if err != nil {
						return false, err
					}
					addSigSuccess, err := ctx.AddSignature(msc.ToContract(), pair.PublicKey, signature)
					if err != nil {
						return false, err
					}
					fSuccess = fSuccess || addSigSuccess
					if fSuccess {
						m--
					}
					if ctx.GetCompleted() || m <= 0 {
						break
					}
				}
				continue
			} else if account.HasKey() {
				// Try to sign with regular accounts
				pair, err := account.GetKey()
				if err != nil {
					return false, err
				}
				signature, err := Sign(ctx.Verifiable, pair, magic)
				if err != nil {
					return false, err
				}
				addSigSuccess, err := ctx.AddSignature(account.GetContract().ToContract(), pair.PublicKey, signature)
				if err != nil {
					return false, err
				}
				fSuccess = fSuccess || addSigSuccess
				continue
			}
		}

		// try smart contract verification
		cs, err := w.GetContractState(&scriptHash)
		if err != nil {
			return false, err
		}
		c, err := cs.ToContract()
		if err != nil {
			return false, err
		}
		if c != nil {
			// Only works with verify without parameters
			if len(c.ParameterList) == 0 {
				fSuccess = fSuccess || ctx.AddItemWithParams(c, nil)
			}
		}
	}

	return fSuccess, nil
}

func (w *WalletHelper) SignTransaction(trx *tx.Transaction, magic uint32) (*tx.Transaction, error) {
	if w.wallet == nil {
		return nil, fmt.Errorf("wallet is nil")
	}
	ctx := NewContractParametersContract(trx)
	_, err := w.Sign(ctx, magic)
	if err != nil {
		return nil, err
	}
	if !ctx.GetCompleted() {
		return nil, fmt.Errorf("context is not completed")
	}
	witnesses, err := ctx.GetWitnesses()
	if err != nil {
		return nil, err
	}
	trx.SetWitnesses(witnesses)
	return trx, nil
}

// Transfer is used to transfer neo or gas or other nep17 asset, from NEP6Account
func (w *WalletHelper) Transfer(assetHash *helper.UInt160, toAddress string, amount *big.Int, magic uint32) (string, error) {
	to, err := crypto.AddressToScriptHash(toAddress, w.wallet.protocolSettings.AddressVersion)
	if err != nil {
		return "", err
	}

	balances, err := w.GetAccountAndBalance(assetHash)
	if err != nil {
		return "", err
	}
	sort.Sort(AccountAndBalanceSlice(balances))
	balancesUsed := findPayingAccounts(balances, amount)
	// add cosigner
	cosigners := make([]tx.Signer, 0)
	sb := sc.NewScriptBuilder()
	for _, used := range balancesUsed {
		cosigners = append(cosigners, tx.Signer{
			Account: used.Account,
			Scopes:  tx.CalledByEntry,
		})
		sb.EmitDynamicCall(assetHash, "transfer", []interface{}{
			sc.ContractParameter{Type: sc.Hash160, Value: used.Account},
			sc.ContractParameter{Type: sc.Hash160, Value: to},
			sc.ContractParameter{Type: sc.Integer, Value: used.Value},
			sc.ContractParameter{Type: sc.String, Value: ""}, // this field is used as a memo
		})
		sb.Emit(sc.ASSERT)
	}
	script, err := sb.ToArray()
	if err != nil {
		return "", err
	}
	balancesGas := make([]AccountAndBalance, 0)
	if assetHash.Equals(tx.GasToken) {
		balancesGas = balances
	} else {
		balancesGas, err = w.GetAccountAndBalance(tx.GasToken)
		if err != nil {
			return "", err
		}
	}
	trx, err := w.MakeTransaction(script, cosigners, []tx.ITransactionAttribute{}, balancesGas)
	if err != nil {
		return "", err
	}
	// sign the tx
	trx, err = w.SignTransaction(trx, magic)
	if err != nil {
		return "", err
	}

	fmt.Println(crypto.Base64Encode(trx.ToByteArray()))
	fmt.Println(helper.BytesToHex(trx.ToByteArray()))
	fmt.Println(trx.GetHash().String())
	// use RPC to send the tx
	response := w.Client.SendRawTransaction(crypto.Base64Encode(trx.ToByteArray()))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	return response.Result.Hash, nil
}
