package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	airDropName    = "AirDrop"
	airDropVersion = "1.0"
	airdropAddr    = "0x4fC9E7E947Bb2269E0B356B038c3543f6b2f7f0C"

	signerPrivateKey = "8793cb44dbaf5a1e0adc780714b90f719620a63e509b334e8cc7224f77321581"
	signerAddr       = "0x6F17FeF0499809cC3988a11e625e8FcFF1C4be29"
	chainId          = 97
)

var (
	domainSeparator []byte
	privateKey      *ecdsa.PrivateKey
)

func init() {
	domainSeparator = buildDomainSeparator()
	rand.Seed(time.Now().UnixMilli())

	//将以十六进制表示的私钥转换为ECDSA（椭圆曲线数字签名算法）私钥对象
	key, err := crypto.HexToECDSA(signerPrivateKey)
	if err != nil {
		panic("private key invalid: " + err.Error())
	}
	privateKey = key
	validatePrivateKeyAddress()
}

// 签名
func Eip712Sign(nftAddr string, code string) (string, error) {
	hash, err := Eip712Digest(nftAddr, code)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	//使用私钥对数据进行签名
	sig, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Println("私钥签名错误:", err)
		return "", err
	}
	// THIS IS VERY IMPORTANT!!! add version by 27
	sig[len(sig)-1] += 27
	return hex.EncodeToString(sig), nil
}

// 将数据hash code MUST be hexify and prefix by 0x
func Eip712Digest(nftAddr string, code string) (hash common.Hash, err error) {
	if strings.HasPrefix(code, "0x") || strings.HasPrefix(code, "0X") {
		code = code[2:]
	}

	codeVal, err := strconv.ParseUint(code, 16, 64)
	if err != nil {
		return
	}
	eip712Prefix := []byte{0x19, 0x01}
	data := crypto.Keccak256(
		common.LeftPadBytes(common.HexToAddress(nftAddr).Bytes(), 32),
		common.LeftPadBytes(big.NewInt(0).SetUint64(codeVal).Bytes(), 32),
	)
	var buf []byte
	buf = append(buf, eip712Prefix...)
	buf = append(buf, domainSeparator...)
	buf = append(buf, common.LeftPadBytes(data, 32)...)

	return crypto.Keccak256Hash(buf), nil
}

// /////private method
// 使用私钥生成钱包地址
func privateKeyToAddress(pk string) (addr common.Address) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}
	addr = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}

// 验证私钥与签名地址是否匹配
func validatePrivateKeyAddress() {
	pk := signerPrivateKey
	address := signerAddr

	//使用私钥生成地址
	addr := privateKeyToAddress(pk)

	if !strings.EqualFold(address, addr.String()) {
		panic(fmt.Sprintf("address derived from privateKey NOT equal with config: %s %s", address, addr.String()))
	}
}

func buildDomainSeparator() []byte {
	i := big.NewInt(0)
	val, ok := i.SetString("8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f", 16)
	if !ok {
		panic("invalid _TYPE_HASH")
	}

	nameHash := common.LeftPadBytes(crypto.Keccak256([]byte(airDropName)), 32)
	versionHash := common.LeftPadBytes(crypto.Keccak256([]byte(airDropVersion)), 32)

	hash := crypto.Keccak256(
		common.LeftPadBytes(val.Bytes(), 32),
		nameHash,
		versionHash,
		common.LeftPadBytes(big.NewInt(int64(chainId)).Bytes(), 32),
		common.LeftPadBytes(common.HexToAddress(airdropAddr).Bytes(), 32),
	)
	return common.LeftPadBytes(hash, 32) //crypto.Keccak256Hash([]byte("AirDrop")).String()
}
