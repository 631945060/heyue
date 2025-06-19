package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main4() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/45dc4d4d5e1b4cf992f3532b186fc58f")
	if err != nil {
		log.Fatal(err)
	}

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

	// 获取 nonce
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 估算 gas 价格
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 准备交易数据
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[],"name":"count","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decrement","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"increment","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}
	hexToAddress := common.HexToAddress(contractAddr)
	methodName := "retrieve"
	input, err := contractABI.Pack(methodName)
	err = funcGetNum(fromAddress, hexToAddress, input, err, client, contractABI, methodName)

	//funcName(nonce, hexToAddress, gasPrice, input, err, privateKey, client, contractABI, methodName)
}

func funcGetNum(fromAddress common.Address, hexToAddress common.Address, input []byte, err error, client *ethclient.Client, contractABI abi.ABI, methodName string) error {
	// 准备调用合约的消息
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &hexToAddress,
		Data: input,
	}
	// 调用合约的 retrieve 方法
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 解析返回值
	var count *big.Int
	err = contractABI.UnpackIntoInterface(&count, methodName, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("count: ", count)
	return err
}

func funcName(nonce uint64, hexToAddress common.Address, gasPrice *big.Int, input []byte, err error, privateKey *ecdsa.PrivateKey, client *ethclient.Client, contractABI abi.ABI, methodName string) {
	tx := types.NewTransaction(nonce, hexToAddress, big.NewInt(0), 300000, gasPrice, input)
	// 创建交易并签名交易
	chainID := big.NewInt(int64(11155111))
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signedTx:  ", signedTx)
	fmt.Println("input:  ", input)
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReceipt2(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	to := common.HexToAddress(contractAddr)
	callInput, err := contractABI.Pack(methodName)
	if err != nil {
		log.Fatal(err)
	}
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}
	fmt.Println("callMsg:  ", callMsg)
	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	contractABI.UnpackIntoInterface(&unpacked, methodName, result)
	if err != nil {
		log.Fatal(err)
	}
}
func waitForReceipt2(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
