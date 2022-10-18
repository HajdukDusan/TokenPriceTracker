# Token Price Tracker

Token Price Tracker is a server implemented in golang that tracks and updates multiple token prices on a deployed smart contract.

### Contract
The contract is currently deployed on the Goerli testnet [here](https://goerli.etherscan.io/address/0x7fcc8d5733782d38b84675e92703a4b7d26a77cb). It stores prices of multiple tokens that are updated every minute. Token price is updated only if the difference in the on-chain and market price is more than 2%.

### Server

The golang server updates the prices of tokens by pulling that information from the [Coingecko](https://www.coingecko.com/en/api/documentation) API. If the price difference is more than 2% the server will send new transactions to the contract that override current pending transactions in the mempool, because by that time they would be deprecated. All blockchain related functionality is done through the [Go-Ethereum](https://github.com/ethereum/go-ethereum) library. The server does not use any kind of permanent storage and all history information is fetched from the contract logs. 

Adding a new token to be tracked is done by adding its Coingecko ID in the symbols array of the Config.go file.
```js
    var Symbols = []string{
        "ethereum",
        "bitcoin",
        "tether",
        "litecoin",
        "decentraland",
        "ravencoin",
    }
```


### Endpoints

- API endpoint for retrieving the price history of a token in the specified timestamp period.

    example:`
        api/history?symbol=ravencoin&from=0&to=2663405165
    `

- API endpoint for retrieving the on-chain and market current prices for the specified token.

    example:`
        api/price?symbol=ravencoin
    `

- Web Socket subscription for on-chain price change events.

### Setup
1. Create a .env file and fill the information specified in the .env.dev file.
2. Run the make file. It will generate the relevant abi and bin files from the smart contract.