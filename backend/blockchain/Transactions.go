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
// This function will ensure that transactions dont get stuck in the pending state
// when we send the TxBlock again. The next TxBlock will override pending
// transactions of the previous block.
//
// Usage:
//
// - First Call - Creates a new block of tx, Returns a func that is used for sending individual txs in that block.
// Takes in a private key for signing txs.
//
// - Second Call - Takes in a func that should send the tx to the mempool and return results.
// 				   txFunction example:
// 					func(auth *bind.TransactOpts) (*types.Transaction, error) {
// 						return yourContractApi.FunctionCall(auth, param1, param2..)
// 					}
//
func CreateTxBlock(privateKey string) func(txFunction func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {

	auth, err := getAccountAuth(privateKey)
	if err != nil {
		fmt.Println(err)
	}

	return func(txFunction func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
		fmt.Println("Sending with:")
		fmt.Println("nonce = ", auth.Nonce)
		fmt.Println("gasPrice = ", auth.GasPrice)

		// try 3 times to increase the gas price
		for i:=0; i < 3; i++{
			// send tx to mempool
			tx, err := txFunction(auth)
			if err != nil {

				// if the gas price is lower than the last block gas price we add 10% and try again
				if err.Error() == "replacement transaction underpriced" {
					addition := big.NewInt(0)
					addition.Div(auth.GasPrice, big.NewInt(10))
					auth.GasPrice.Add(auth.GasPrice, addition)
					continue
				}

				return nil, err
			}
			// if tx was sent into the mempool we inc the nonce
			auth.Nonce.Add(auth.Nonce, big.NewInt(1))

			return tx, nil
		}
		return nil, errors.New("Pending tx has a high gas price")
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
