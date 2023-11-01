package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"os"
)

func GetAirDropAmount() {
	contractAddress := common.HexToAddress(TPLAddress())
	fmt.Println("addr:", contractAddress)

	contractABI := GetABI("./ABI/erc20.json")
	log.Println("合约ABI加载成功...")

	name, err := contractABI.Pack("name")
	fmt.Println("abi:", contractABI)

	if err != nil {
		log.Fatal(err)
	}
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: name,
	}

	fmt.Println("msg:", callMsg)

	nameResult, err := ethClient.CallContract(context.Background(), callMsg, big.NewInt(33953364))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("erc20:name-", nameResult)

}

func CouponAddress() string {
	return os.Getenv("COUPON_ADDR")
}

func Airdropddress() string {
	return os.Getenv("AIRDROP_ADDR")
}

func DepositAddress() string {
	return os.Getenv("DOPOSIT_ADDR")
}

func FellowNFTAddress() string {
	return os.Getenv("FELLOWNFT_ADDR")
}

func GenesisNFTAddress() string {
	return os.Getenv("GENESISNFT_ADDR")
}

func MedalNFTAddress() string {
	return os.Getenv("MEDALNFT_ADDR")
}

func CommunityNFTAddress() string {
	return os.Getenv("COMMUNITY_ADDR")
}

func USDTAddress() string {
	return os.Getenv("USDT_ADDR")
}

func TPLAddress() string {
	return os.Getenv("TPL_ADDR")
}

func FormatAddress(address string) string {
	return "0x" + address[27:]
}

func HashToInt(hash common.Hash) uint {
	bigInt := new(big.Int)
	bigInt.SetBytes(hash[:])
	return BigIntToUint(bigInt)
}
