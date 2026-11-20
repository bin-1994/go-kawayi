package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strings"
	"time"
)

type PassportController struct {
	Service service.IUserService
}

func NewPassportController(service service.IUserService) *PassportController {
	return &PassportController{Service: service}
}

const refresh_secret = "68d33f3c48149a45afd5f76d22c3875c"
const access_secret = "58df9b68b41e798923fe3ab883860b6f"

func (controller *PassportController) Login(ctx *gin.Context) {
	email := strings.TrimSpace(ctx.PostForm("email"))
	password := strings.TrimSpace(ctx.PostForm("password"))
	fmt.Println(email)
	resp := utils.NewResponse(ctx)

	user, err := controller.Service.Login(email, password)
	if err != nil {
		resp.Error(err.Error())
		return
	}

	// 创建7天刷新token
	refresh_token := utils.CreateToken(refresh_secret, &utils.UserClaims{
		Uid:   user.Uid,
		Email: user.Email,
	}, time.Hour*time.Duration(24)*7)

	// 创建2小时权限token
	access_token := utils.CreateToken(access_secret, &utils.UserClaims{
		Uid:   user.Uid,
		Email: user.Email,
	}, time.Hour*time.Duration(24)*7)

	resp.Success(map[string]any{
		"token": map[string]any{
			"refresh_token": refresh_token,
			"access_token":  access_token,
		},
		"info": map[string]any{
			"uid":      user.Uid,
			"email":    user.Email,
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
		},
	})
}

func (controller *PassportController) Register(ctx *gin.Context) {
	email := strings.TrimSpace(ctx.PostForm("email"))
	password := strings.TrimSpace(ctx.PostForm("password"))
	nickname := strings.TrimSpace(ctx.PostForm("nickname"))
	regCode := strings.TrimSpace(ctx.PostForm("regCode"))

	user, err := controller.Service.Register(email, password, nickname, regCode, ctx.ClientIP())
	resp := utils.NewResponse(ctx)
	if err != nil {
		resp.Error(err.Error())
		return
	}

	resp.Success(user)
}

func (controller *PassportController) CheckName(ctx *gin.Context) {
	nickname := ctx.Query("nickname")
	resp := utils.NewResponse(ctx)
	if err := controller.Service.CheckNickname(nickname); err != nil {
		resp.Error(err.Error())
		return
	}
	resp.Success(nil)
}
