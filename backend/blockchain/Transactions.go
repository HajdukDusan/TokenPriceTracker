package blockchain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// CreateTxBlock is a higher-order function for sending multiple transactions.
//
// This function will ensure that transactions dont get stuck in the pending state
// when we send the TxBlock again.
//
// The next TxBlock will override pending transactions of the previous block.
//
// Params of the anonymous function are contract call params.
func CreateTxBlock() func(string, *big.Int) (*types.Transaction, error) {

	auth, err := getAccountAuth(PrivateKey)
	if err != nil {
		fmt.Println(err)
	}

	return func(symbol string, price *big.Int) (*types.Transaction, error) {

		fmt.Println("\nSending with:")
		fmt.Println("nonce = ", auth.Nonce)
		fmt.Println("gasPrice = ", auth.GasPrice)

		tx, err := PriceSetterContract.SetSymbolPrice(auth, symbol, price)
		if err != nil {
			return nil, err
		}

		// if tx was sent into the mempool we inc the nonce
		auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		return tx, nil
	}
}

// Function to create auth from privateKey
func getAccountAuth(privateKey string) (*bind.TransactOpts, error) {

	// check if the rpc connection is initialized
	if Client == nil {
		return nil, errors.New("RPC connection is not initialized, try to call Connect() first...")
	}
	
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	fromAddress, err := getAccountAddressFromPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	// get the current nonce of passed txs, ignore pending txs
	nonce, err := Client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return nil, err
	}

	chainID, err := Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

// Helper function to get the account address from private key
func getAccountAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return [common.AddressLength]byte{}, errors.New("Failed to recognize ECDSA in public key")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
