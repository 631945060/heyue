package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	// 引入以太坊相关的包，用于与以太坊网络进行交互
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 定义常量 contractBytecode，存储合约的字节码
const (
	// store合约的字节码
	contractBytecode = "6080604052348015600e575f5ffd5b506101d98061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806306661abd1461004e5780632baeceb71461006c5780632e64cec114610076578063d09de08a14610094575b5f5ffd5b61005661009e565b60405161006391906100f7565b60405180910390f35b6100746100a3565b005b61007e6100bd565b60405161008b91906100f7565b60405180910390f35b61009c6100c5565b005b5f5481565b60015f5f8282546100b4919061013d565b92505081905550565b5f5f54905090565b60015f5f8282546100d69190610170565b92505081905550565b5f819050919050565b6100f1816100df565b82525050565b5f60208201905061010a5f8301846100e8565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610147826100df565b9150610152836100df565b925082820390508181111561016a57610169610110565b5b92915050565b5f61017a826100df565b9150610185836100df565b925082820190508082111561019d5761019c610110565b5b9291505056fea2646970667358221220dc37cfad815b3c92db5c1e21f43fe50bd10efcea6dd8c1a6b7be2a7c5194bf0964736f6c634300081e0033"
)

func main1() {
	// 连接到以太坊网络（这里使用 sepolia 测试网络作为示例）
	// 通过 ethclient.Dial 函数连接到指定的以太坊节点
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/45dc4d4d5e1b4cf992f3532b186fc58f")
	if err != nil {
		// 如果连接失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 创建私钥（在实际应用中，您应该使用更安全的方式来管理私钥）
	// 通过 crypto.HexToECDSA 函数将十六进制字符串转换为 ECDSA 私钥
	privateKey, err := crypto.HexToECDSA("e642666a050baebccd0f108edf718a79038c895164b6c40fa732cdf367369159")
	if err != nil {
		// 如果转换失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	// 将公钥转换为 ECDSA 公钥类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// 如果转换失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal("error casting public key to ECDSA")
	}

	// 从公钥生成以太坊地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取nonce
	// nonce 是一个单调递增的计数器，用于确保交易的唯一性
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		// 如果获取 nonce 失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 获取建议的gas价格
	// gas 价格决定了交易的优先级和成本
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		// 如果获取 gas 价格失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 解码合约字节码
	// 将十六进制的合约字节码解码为字节切片
	data, err := hex.DecodeString(contractBytecode)
	if err != nil {
		// 如果解码失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 创建交易
	// types.NewContractCreation 用于创建一个合约部署交易
	// 参数依次为 nonce、交易金额（这里为 0）、gas 限制、gas 价格和合约字节码
	tx := types.NewContractCreation(nonce, big.NewInt(0), 3000000, gasPrice, data)

	// 签名交易
	// 获取当前网络的链 ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		// 如果获取链 ID 失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 使用私钥对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		// 如果签名失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 发送交易
	// 通过 client.SendTransaction 函数将签名后的交易发送到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		// 如果发送交易失败，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 输出交易哈希
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// 等待交易被挖矿
	// 调用 waitForReceipt 函数等待交易被打包进区块
	receipt, err := waitForReceipt(client, signedTx.Hash())
	if err != nil {
		// 如果等待过程中出现错误，使用 log.Fatal 输出错误信息并终止程序
		log.Fatal(err)
	}

	// 输出合约部署地址
	fmt.Printf("Contract deployed at: %s\n", receipt.ContractAddress.Hex())
}

// waitForReceipt 函数用于等待交易被挖矿并返回交易收据
func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		// 通过 client.TransactionReceipt 函数查询交易收据
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			// 如果没有错误，说明交易已经被打包，返回交易收据
			return receipt, nil
		}
		if err != ethereum.NotFound {
			// 如果出现除交易未找到之外的错误，返回错误信息
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
