package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
)

type AnswerController struct {
	Service service.IAnswerService
}

func NewAnswerController(service service.IAnswerService) *AnswerController {
	return &AnswerController{Service: service}
}

func (controller *AnswerController) FindAll(ctx *gin.Context) {
	resp := utils.NewResponse(ctx)

	list, countdown, now := controller.Service.FindAll(ctx.ClientIP())

	resp.Success(map[string]interface{}{
		"list": list,
		"timestamp": map[string]interface{}{
			"end":       countdown - (int(now.Unix())),
			"countdown": countdown,
		},
	})
}

func (controller *AnswerController) Find(ctx *gin.Context) {

}
