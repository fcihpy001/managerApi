package main

import (
	"ManagerApi/model"
	"ManagerApi/router"
	"ManagerApi/service"
	"ManagerApi/utils"
)

func main() {
	//=======================================================
	// 1. 读取配置文件
	utils.InitConfig()

	//=======================================================
	// 2. 初始化数据库
	service.InitDB()
	//makeNFTInfo()

	router.InitRouter()
}

func makeNFTInfo() {
	n1 := model.NFT{
		ContractAddress: "0x950Ae2649f36983b62593Bd880986d2D0C0C3403",
		Type:            "member",
	}
	n2 := model.NFT{
		ContractAddress: "0x4BF5baDdA03059b087bF3028220B00a8FD757592",
		Type:            "medal",
	}

	n3 := model.NFT{
		ContractAddress: "0x50D27f4506100DDcdb305C222328eB6510d4B7A5",
		Type:            "genesis",
	}

	n4 := model.NFT{
		ContractAddress: "0xa2880614540B40e3c95e47f09671aAe6F03D7B9A",
		Type:            "community",
	}
	n5 := model.NFT{
		ContractAddress: "0x50D27f4506100DDcdb305C222328eB6510d4B7A5",
		Type:            "fellow",
	}
	models := []model.NFT{}
	models = append(models, n1)
	models = append(models, n2)
	models = append(models, n3)
	models = append(models, n4)
	models = append(models, n5)
	service.DB.Create(models)
}
