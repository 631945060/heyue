package main

import (
	"context"
	"crypto/ecdsa" // 添加导入语句
	Begging "ethc/begging"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 导入 bind 包
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"log"
	"math/big"
)

func main2() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/45dc4d4d5e1b4cf992f3532b186fc58f")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.GenerateKey()
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
	// fmt.Println("Private Key:", privateKeyHex)
	privateKey, err := crypto.HexToECDSA("e642666a050baebccd0f108edf718a79038c895164b6c40fa732cdf367369159")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := Begging.DeployBeggingContractV1FlattenSolBeggingContractV1(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   //0xcdE4D48c503B654F248F23Da9C63AC9BCf5AD8B8
	fmt.Println(tx.Hash().Hex()) //0x3771221d0f3c14a6250489f5f0590fa801b508b823e1dd58b41f14025debf035

	_ = instance
}
