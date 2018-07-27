package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Querying blocks with go-ethereum is a simple process. You can call the client's HeaderByNumber to return header
	// information about a block. It'll return the latest block header if you pass nil .
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String())

	// Call the client's BlockByNumber method to get the full block. You can read all the contents and metadata of the block
	// such as block number, block timestamp, block hash, block difficulty, as well as the list of transactions and much much
	// more.
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time().Uint64())
	fmt.Println(block.Difficulty().Uint64())
	fmt.Println(block.Hash().Hex())
	fmt.Println(len(block.Transactions()))

	// Call TransactionCount to return just the count of transactions in a block.
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
