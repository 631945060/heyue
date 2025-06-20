package main

import (
	// "context"
	"context"
	"crypto/ecdsa"
	Begging "ethc/begging"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	//contractAddr = "0x2FBEdC849dD62b5c8CCb49d65beD44a3FA295c4B"
	contractAddr = "0x9496D1EE9747ce59317A39F2fD3419290479139c"
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
	ethAmount := big.NewFloat(0.00001)

	ethToWei := big.NewFloat(1e18)
	weiAmount := new(big.Int)
	ethAmount.Mul(ethAmount, ethToWei)
	ethAmount.Int(weiAmount)
	opt.GasLimit = uint64(300000)
	opt.GasPrice = nil
	opt.Value = weiAmount
	//tt, err := storeContract.Donate(opt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Transaction Hash: %s\n", tt.Hash().Hex())
	// 捐赠者地址，你需要替换为实际的捐赠者地址
	donorAddress := common.HexToAddress("0x0109244e392b0613b4048F7C97601bd94a638B16")

	callOpts := &bind.CallOpts{
		Context: context.Background(),
		From:    fromAddress,
	}
	donation, err := storeContract.GetDonation(callOpts, fromAddress)
	if err != nil {
		log.Fatalf("Failed to call GetDonation method: %v", err)
	}
	fmt.Printf("Donation amount for donor %s: %s\n", donorAddress.Hex(), donation.String())

	//getDonation(address donor)
	methodSignature := []byte("getDonation(address)")
	methodSelector := crypto.Keccak256(methodSignature)[:4] // 函数选择器

	// 构造参数：address 是 20 字节，需右对齐填充到 32 字节
	var paddedAddress [32]byte
	copy(paddedAddress[12:], fromAddress.Bytes()) // 地址放在右边 20 字节位置

	// 构造调用数据
	input := append([]byte{}, methodSelector...)
	input = append(input, paddedAddress[:]...)

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: input,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	var unpacked [32]byte
	copy(unpacked[:], result)
	fmt.Println("is value saving in contract equals to origin value:", unpacked)

}
