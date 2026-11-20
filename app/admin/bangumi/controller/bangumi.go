package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/service"
)

var NewBangumiController = newBangumiController

func newBangumiController() *bangumiController {
	return &bangumiController{}
}

type BangumiController interface {
	Get()
}

type bangumiController struct {
	bangumiService service.BangumiService
}

func (this *bangumiController) Get(ctx *gin.Context) {

}

func (this *bangumiController) Create() {

}

func (this *bangumiController) Edit() {

}
