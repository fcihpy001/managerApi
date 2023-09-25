package api

import (
	"ManagerApi/model"
	"ManagerApi/service"
	"ManagerApi/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Create(ctx *gin.Context) {
	//	要获取团队名、数量、过期时间
	code := model.CodeRequest{}
	if err := ctx.ShouldBind(&code); err != nil {
		ErrorResp(ctx, 400, "缺少必须的参数", nil)
		return
	}

	models := []model.ActiveCode{}
	for i := 0; i < int(code.Count); i++ {
		c := model.ActiveCode{
			Code:       utils.RandStringAndNumber(6),
			GroupName:  code.GroupName,
			Status:     0,
			Address:    "",
			Expiration: time.Now().Add(time.Duration(code.Days*24) * time.Hour),
		}
		models = append(models, c)
	}

	//	将数据插入数据库中
	result := service.DB.Create(&models)

	//	返回结果
	SuccessResp(ctx, "创建成功", gin.H{
		"success": result.RowsAffected,
		"failure": int64(code.Count) - result.RowsAffected,
	})
}

func List(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var codes []model.ActiveCode
	service.DB.Order("id desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&codes)

	// 记录的总条数
	var total int64
	service.DB.Model(model.ActiveCode{}).Count(&total)

	//	返回结果
	SuccessResp(ctx, "success", gin.H{"count": total, "list": codes})
}

func Update(ctx *gin.Context) {
	code := ctx.PostForm("code")
	address := ctx.PostForm("address")
	if len(code) == 0 || len(code) < 6 {
		ErrorResp(ctx, 400, "code 有误", nil)
		return
	}

	if len(address) == 0 {
		ErrorResp(ctx, 400, "地址无效", nil)
		return
	}
	c := model.ActiveCode{
		Code:    code,
		Address: address,
	}

	//	更新入库
	result := service.DB.Where("code = ?", code).Save(&c)
	if result.RowsAffected > 0 {
		//	返回结果
		SuccessResp(ctx, "success", nil)
		return
	}
	ErrorResp(ctx, 400, "code无效", nil)

}
