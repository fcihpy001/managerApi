package utils

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

var (
	ethClient *ethclient.Client
)

func GetABI(abiJson string) abi.ABI {
	file, err := os.ReadFile(abiJson)
	if err != nil {
		log.Panicln("文件读取失败")
	}
	wrapABI, err := abi.JSON(bytes.NewReader(file))
	return wrapABI
}

func GetEthClient() *ethclient.Client {
	if ethClient == nil {
		dial, err := ethclient.Dial("https://bsc-testnet.public.blastapi.io")
		if err != nil {
			log.Fatalf("连接以太坊节点失败：%v", err)
		}
		ethClient = dial
		fmt.Println("eth_client节点初始化成功")
		defer dial.Close()
	}
	return ethClient
}

func BigIntToUint(i *big.Int) uint {
	// 检查是否为负数
	if i.Sign() == -1 {
		return 0
	}

	// 将 big.Int 转换为 *big.Float
	f := new(big.Float).SetInt(i)

	// 将 *big.Float 转换为 uint
	uintVal, _ := f.Uint64()

	return uint(uintVal)
}
