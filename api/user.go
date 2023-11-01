package api

import (
	"ManagerApi/middleware"
	"ManagerApi/model"
	"ManagerApi/service"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SetupUserRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/user")
	v1Router.Use(gin.Logger())
	v1Router.POST("login", login)
	v1Router.POST("register", register)
	v1Router.GET("info", middleware.AuthMiddleWare(), info)
	v1Router.POST("wallet/login", middleware.VerifyHeader(), walletLogin)
}

func login(c *gin.Context) {
	requestUser := &model.User{}
	if err := c.ShouldBind(&requestUser); err != nil {
		ErrorResp(c, 405, "参数缺失", nil)
		return
	}
	//判断用户是否存在
	user := model.User{}
	service.GetDB().Where("user_name = ?", requestUser.UserName).First(&user)
	if user.ID == 0 {
		ErrorResp(c, 400, "用户不存在", nil)
		return
	}

	fmt.Println("passwod")
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
		ErrorResp(c, 400, "密码错误", nil)
		return
	}

	//	发送token
	token, err := utils.GetToken(user.ID)
	if err != nil {
		response(c, http.StatusUnprocessableEntity, 500, "token 生成失败", nil)
		return
	}

	//	返回结果
	SuccessResp(c, "登录成功", gin.H{"token": token})
}

func register(c *gin.Context) {

	var header model.HeaderData
	c.BindHeader(&header)

	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		ErrorResp(c, 405, "参数缺失", nil)
		return
	}

	//数据验证
	if len(user.UserName) < 3 {
		ErrorResp(c, http.StatusUnprocessableEntity, "用户名格式有误", nil)
		return
	}

	// 密码加密
	hashPasswd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response(c, http.StatusUnprocessableEntity, 500, "加密失败", nil)
		return
	}

	// 数据入库
	user.Password = string(hashPasswd)
	err = service.GetDB().Create(&user).Error
	if err != nil {
		response(c, http.StatusUnprocessableEntity, 500, "数据入库异常", nil)
		return
	}

	//	发送token
	token, err := utils.GetToken(user.ID)
	if err != nil {
		response(c, http.StatusUnprocessableEntity, 500, "token 生成失败", nil)
		return
	}

	//	返回结果
	SuccessResp(c, "注册成功", gin.H{"token": token})
}

func info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	SuccessResp(ctx, "success", model.ToUserDTO(user.(model.User)))
}
