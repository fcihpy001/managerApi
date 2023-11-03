package api

import (
	_const "ManagerApi/const"
	"ManagerApi/middleware"
	"ManagerApi/model"
	"ManagerApi/service"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func SetupActiveCodeRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/code")
	v1Router.Use(gin.Logger())
	v1Router.POST("create", create)
	v1Router.GET("list", list)
	v1Router.POST("enable", middleware.VerifyHeader(), enable)
}

func create(ctx *gin.Context) {
	//	要获取团队名、数量、过期时间
	code := model.CodeRequest{}
	if err := ctx.ShouldBind(&code); err != nil {
		ErrorResp(ctx, 400, _const.ErrorBodyMsg, nil)
		return
	}

	var models []model.ActiveCode
	var codes []string
	for i := 0; i < int(code.Count); i++ {
		code_str := utils.RandStringAndNumber(6)
		c := model.ActiveCode{
			Code:       code_str,
			GroupName:  code.GroupName,
			Status:     0,
			NFT:        code.NFT,
			Expiration: time.Now().Add(time.Duration(code.Days*24) * time.Hour),
		}
		models = append(models, c)
		codes = append(codes, code_str)
	}

	//	将数据插入数据库中
	result := service.GetDB().Create(&models)

	//	返回结果
	SuccessResp(ctx, "创建成功", gin.H{
		"success": result.RowsAffected,
		"failure": int64(code.Count) - result.RowsAffected,
		"codes":   codes,
	})
}

func list(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pagenum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "20"))

	// 分页
	var codes []model.ActiveCode
	service.GetDB().Where("wallet_address = ''").Order("id desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&codes)

	// 记录的总条数
	var total int64
	service.GetDB().Model(model.ActiveCode{}).Count(&total)

	//	返回结果
	SuccessResp(ctx, "success", gin.H{"count": total, "list": codes})
}

func enable(ctx *gin.Context) {
	//	根据code,找到nft
	code := ctx.PostForm("code")
	address := ctx.PostForm("wallet")
	if len(code) < 6 {
		ErrorResp(ctx, 400, "code 有误", nil)
		return
	}

	if len(address) < 40 {
		ErrorResp(ctx, 400, "地址无效", nil)
		return
	}

	var result struct {
		Code string `json:"code"`
		NFT  string `json:"nft"`
		Sign string `json:"sign"`
	}

	// 获取NFT地址
	service.GetDB().Table("active_code").
		Select("active_code.code AS code,contract.address AS NFT").
		Joins("INNER JOIN contract ON contract.name = active_code.nft").
		Where("active_code.code = ?", code).
		Scan(&result)

	//返回NFT合约地址和签名数据
	hash, err := utils.Eip712Digest(result.NFT, result.Code)
	if err != nil {
		fmt.Println("加密错误", err)
	}
	fmt.Println("digest:", hash.String())

	sign, err := utils.Eip712Sign(result.NFT, result.Code)
	fmt.Println("signature:", sign)
	result.Sign = sign

	//	将钱包地址与code进行关联, 并更新入库
	service.GetDB().Model(&model.ActiveCode{Code: code}).Update("wallet_address", address)

	SuccessResp(ctx, "", result)
}
