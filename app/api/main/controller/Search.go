package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strings"
)

//var seg gse.Segmenter
//func init() {
//seg.New("./plugins/sego/dictionary.txt")
//seg = gse.New("./plugins/sego/dictionary.txt", "alpha")
//seg.LoadStop("./plugins/sego/stop_tokens.txt")
//}
//var segmenter sego.Segmenter
//func init() {
//	segmenter.LoadDictionary("./plugins/sego/dictionary.txt")
//}

type SearchController struct {
	Service service.ISearchService
}

func NewSearchController(service service.ISearchService) *SearchController {
	return &SearchController{Service: service}
}

func (controller *SearchController) FindAll(ctx *gin.Context) {
	keyword := strings.Replace(strings.TrimSpace(ctx.Query("keyword")), "-", "", -1)

	resp := utils.NewResponse(ctx)

	if keyword == "" {
		resp.Error("搜索关键词不能为空")
		return
	}
	page := utils.Paging{Page: ctx.Query("page"), Size: ctx.Query("size")}
	offset, limit := page.SetPage()

	// 创建热搜词
	controller.Service.Create(keyword, 1)

	list, total := controller.Service.FindAll(keyword, true, offset, limit)

	resp.List(list, offset, limit, total)
}

func (controller *SearchController) FetchSugget(ctx *gin.Context) {
	keyword := strings.Replace(strings.TrimSpace(ctx.Query("keyword")), "-", "", -1)
	resp := utils.NewResponse(ctx)

	list, _ := controller.Service.FindAll(keyword, false, 0, 10)

	resp.Success(map[string]any{
		"list": list,
	})
}

func (controller *SearchController) FetchHotword(ctx *gin.Context) {
	resp := utils.NewResponse(ctx)

	list := controller.Service.FindHot(0, 9)

	resp.Success(map[string]any{
		"list": list,
	})
}
