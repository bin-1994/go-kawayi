package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
)

type BannerController struct {
	Service service.IBannerService
}

func NewBannerController(service service.IBannerService) *BannerController {
	return &BannerController{Service: service}
}

func (controller *BannerController) FindAll(ctx *gin.Context) {
	resp := utils.NewResponse(ctx)

	list := controller.Service.FindAll(0, 5)

	resp.Success(map[string]any{
		"list": list,
	})
}
