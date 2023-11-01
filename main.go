package main

import (
	"ManagerApi/model"
	"ManagerApi/router"
	"ManagerApi/service"
	"ManagerApi/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//=======================================================
	// 1. 读取配置文件
	utils.InitConfig()

	//makeContract()

	router.InitRouter()

}

func makeContract() {
	n1 := model.Contract{
		Address: "0x950Ae2649f36983b62593Bd880986d2D0C0C3403",
		Name:    "member",
		Desc:    "会员NFT",
	}
	n2 := model.Contract{
		Address: "0x4BF5baDdA03059b087bF3028220B00a8FD757592",
		Name:    "medal",
		Desc:    "勋章NFT",
	}

	n3 := model.Contract{
		Address: "0x50D27f4506100DDcdb305C222328eB6510d4B7A5",
		Name:    "genesis",
		Desc:    "创世NFT",
	}

	n4 := model.Contract{
		Address: "0xa2880614540B40e3c95e47f09671aAe6F03D7B9A",
		Name:    "community",
		Desc:    "社区NFT",
	}
	n5 := model.Contract{
		Address: "0x50D27f4506100DDcdb305C222328eB6510d4B7A5",
		Name:    "fellow",
		Desc:    "伙伴勋章NFT",
	}
	n6 := model.Contract{
		Address: "0x17cAb4b89f42091c9AA53656dBbA38265b38BD59",
		Name:    "TPL",
		Desc:    "toplink代币",
	}
	n7 := model.Contract{
		Address: "0x5c3bc9Ab21c730047173a16071d4b4d5e12c1135",
		Name:    "deposit",
		Desc:    "质押合约",
	}
	n8 := model.Contract{
		Address: "0x4fC9E7E947Bb2269E0B356B038c3543f6b2f7f0C",
		Name:    "airdrop",
		Desc:    "空投合约",
	}
	n9 := model.Contract{
		Address: "0x6e5F2cAe6185A8023BE3De3b038Cac86310A8842",
		Name:    "lp",
		Desc:    "swap合约",
	}
	n10 := model.Contract{
		Address: "0x514CB68CC8eA0a2C4BAbB02Dc1A234E98525F858",
		Name:    "coupon",
		Desc:    "团购合约",
	}

	models := []model.Contract{}
	models = append(models, n1)
	models = append(models, n2)
	models = append(models, n3)
	models = append(models, n4)
	models = append(models, n5)
	models = append(models, n6)
	models = append(models, n7)
	models = append(models, n8)
	models = append(models, n9)
	models = append(models, n10)
	service.GetDB().Create(models)
}
