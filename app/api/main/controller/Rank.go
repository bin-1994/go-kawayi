package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strconv"
)

type RankController struct {
	Service service.IRankService
}

func NewRankController(service service.IRankService) *RankController {
	return &RankController{Service: service}
}

func (controller *RankController) Create(ctx *gin.Context) {
	resp := utils.NewResponse(ctx)

	media_id := ctx.PostForm("media_id")
	ep_id := ctx.PostForm("ep_id")

	//numbers := utils.DecodeHashId(id)
	//if len(numbers) == 0 {
	//	resp.Error("参数错误")
	//	return
	//}
	//media_id := numbers[0]
	//ep_id := numbers[1]

	ip := ctx.ClientIP()
	controller.Service.Create(media_id, ep_id, ip)

	resp.Success(nil)
}

func (controller *RankController) FindAll(ctx *gin.Context) {
	resp := utils.NewResponse(ctx)

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "9"))

	list := controller.Service.FindAll(int64(page), int64(size))

	resp.Success(map[string]any{
		"list": list,
	})
}
