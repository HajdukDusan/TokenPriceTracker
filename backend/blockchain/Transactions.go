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

// higher-order func

func InitializeTxBlockSender() func() func(string, *big.Int) (*types.Transaction, error) {

	var sentTransactions []chan bool

	return func() func(string, *big.Int) (*types.Transaction, error) {

		//cancel all routines that are waiting for pending transactions
		for indx := range sentTransactions {
			sentTransactions[indx] <- true
		}

		// empty the array
		sentTransactions = make([]chan bool, 0)

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

			// make a routine that waits for the pending tx
			sentTransactions = append(sentTransactions, waitTxConfirmed(tx.Hash()))

			// if tx was sent into the mempool we inc the nonce
			auth.Nonce.Add(auth.Nonce, big.NewInt(1))

			return tx, nil
		}
	}
}

func waitTxConfirmed(hash common.Hash) chan bool{
	ch := make(chan bool)
	go func() {
		for {

			select {
				case ch <- n: {
					
				}
				case <-ch: {
					fmt.Println("Tx stopped!")
					break
				}
            }

			_, pending, err := Client.TransactionByHash(context.Background(), hash)

			if err != nil {
				fmt.Println("Tx finished with error: ", err)
				break
			}
			if !pending {
				fmt.Println("Tx confirmed!")
				break
			}
		}
		fmt.Println("Tx hash: ", hash.Hex())
	}()
	return ch
}

// function to create auth for any account from its private key
func getAccountAuth(privateKey string) (*bind.TransactOpts, error) {

	// check if the rpc conn is initialized
	if Client == nil {
		return nil, errors.New("RPC connection is not initialized, try to call Connect() first...")
	}
	
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units

	// gasPrice x2 for fast mining
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))

	return auth, nil
}
