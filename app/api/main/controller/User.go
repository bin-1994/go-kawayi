package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strconv"
)

type UserController struct {
	Service service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{Service: service}
}

func (controller *UserController) FindById(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	resp := utils.NewResponse(ctx)
	user, err := controller.Service.FindById(uid.(int64))
	if err != nil {
		resp.Error(err.Error())
		return
	}

	resp.Success(user)
}

func (controller *UserController) FindState(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	media_id, _ := strconv.Atoi(ctx.Query("media_id"))
	resp := utils.NewResponse(ctx)

	state, err := controller.Service.FindState(uid.(int64), media_id)
	if err != nil {
		resp.Error("查询出错")
		return
	}

	resp.Success(state)
}
