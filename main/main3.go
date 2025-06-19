package main

import (
	// "context"
	"context"
	"crypto/ecdsa"
	Begging "ethc/begging"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	//contractAddr = "0x2FBEdC849dD62b5c8CCb49d65beD44a3FA295c4B"
	contractAddr = "0x0F1E73191BA526F1120D2ea56b055D1CAE04445d"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/45dc4d4d5e1b4cf992f3532b186fc58f")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := Begging.NewBeggingContractV1FlattenSolBeggingContractV1(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("storeContract:%s", storeContract)
	_ = storeContract

	privateKey, err := crypto.HexToECDSA("e642666a050baebccd0f108edf718a79038c895164b6c40fa732cdf367369159")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//fmt.Printf("psivateKey:%s", privateKey)
	// 初始化交易opt实例
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	opt.Nonce = nil
	ethAmount := big.NewFloat(0.001)

	ethToWei := big.NewFloat(1e18)
	weiAmount := new(big.Int)
	ethAmount.Mul(ethAmount, ethToWei)
	ethAmount.Int(weiAmount)
	opt.GasLimit = uint64(300000)
	opt.GasPrice = nil
	opt.Value = weiAmount
	tt, err := storeContract.Donate(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction Hash: %s\n", tt.Hash().Hex())
	// 捐赠者地址，你需要替换为实际的捐赠者地址
	donorAddress := common.HexToAddress("0x0FF404A49B83Ae402fd3F667192A4Acf89d88ba2")

	callOpts := &bind.CallOpts{
		Context: context.Background(),
		From:    fromAddress,
	}
	donation, err := storeContract.GetDonation(callOpts, donorAddress)
	if err != nil {
		log.Fatalf("Failed to call GetDonation method: %v", err)
	}
	fmt.Printf("Donation amount for donor %s: %s\n", donorAddress.Hex(), donation.String())
}
