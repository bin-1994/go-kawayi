package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strconv"
	"time"
)

type BangumiController struct {
	Service service.IBangumiService
}

func NewBangumiController(service service.IBangumiService) *BangumiController {
	return &BangumiController{Service: service}
}

func (controller *BangumiController) FindById(ctx *gin.Context) {
	str_id := ctx.Query("id")
	id, _ := strconv.Atoi(str_id)

	if ep_id, exist := ctx.GetQuery("ep_id"); exist {
		id = controller.Service.EpToId(ep_id)
	}

	resp := utils.NewResponse(ctx)
	//if !utils.IsNumber(str_id) {
	//	number := utils.DecodeHashId(str_id)
	//	if len(number) == 0 {
	//		resp.Error("抱歉出错了")
	//		return
	//	}
	//	id = int(number[0])
	//}
	var info model.BangumiInfoModel
	//key := fmt.Sprintf("bangumi:%d", id)
	//err := utils.NewRedis().Get(key, &info)

	//if err != nil {
	info = controller.Service.FindById(id)
	//utils.NewRedis().Set(key, info, time.Minute*5)
	//}

	resp.Success(info)
}

func (controller *BangumiController) FindAll(ctx *gin.Context) {
	var queryModel model.BangumiQueryModel

	resp := utils.NewResponse(ctx)

	if err := ctx.ShouldBind(&queryModel); err != nil {
		resp.Error("参数错误")
		return
	}

	list, total := controller.Service.FindAll(queryModel)

	resp.List(list, queryModel.Page, queryModel.Size, total)
}

func (controller *BangumiController) FindFilter(ctx *gin.Context) {
	resp := utils.NewResponse(ctx)

	var list model.FilterModel
	if err := utils.NewRedis().Get("bangumi:filter", &list); err != nil {
		list = controller.Service.FindFilter()
		utils.NewRedis().Set("bangumi:filter", list, time.Hour*24*30)
	}
	resp.Success(list)
}
