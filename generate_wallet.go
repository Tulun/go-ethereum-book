package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func main() {
	// To generate a new wallet first we need to import the go-ethereum crypto package that provides the GenerateKey
	// method for generating a random private key.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Then we can convert it to bytes by importing the golang crypto/ecdsa package and using the FromECDSA method.
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// We can now convert it to a hexadecimal string by using the go-ethereum hexutil package which provides the Encode
	// method which takes a byte slice. Then we strip of the 0x after it's hex encoded.
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	// This is the private key which is used for signing transactions and is to be treated like a password and never be shared,
	// since who ever is in possesion of it will have access to all your funds.
	// Since the public key is derived from the private key, go-ethereum's crypto private key has a Public method that will
	// return the public key.
	publicKey := privateKey.Public()

	// Converting it to hex is a similar process that we went through with the private key. We strip off the 0x and the first 2
	// characters 04 which is always the EC prefix and is not required.

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0

	// Now that we have the public key we can easily generate the public address which what you're used to seeing. In order
	// to do the the go-ethereum crypto package has a PubkeyToAddress method which accepts an ECDSA public key, and
	// returns the public address.
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	// The public address is simply the Keccak-256 hash of the public key, and then we take the last 40 characters (20
	// bytes) and prefix it with 0x . Here's how you can do it manually using the go-ethereum's crypto/sha3 Keccak256
	// functions.
	hash := sha3.NewKeccak256()
	hash.Write(publicKeyBytes[1:])

	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
