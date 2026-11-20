package controller

import (
	"fmt"
	"go-biu/app/admin/system/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-biu/app/admin/system/service"
)

type HlsController interface {
	Upload(*gin.Context)
	List(*gin.Context)
	Get(*gin.Context)
	Update(ctx *gin.Context)
}

type controller struct {
	service service.HlsService
}

func NewHlsController(service service.HlsService) HlsController {
	return &controller{service: service}
}

func (c *controller) Upload(ctx *gin.Context) {
	media_id, _ := strconv.Atoi(ctx.PostForm("media_id"))
	if media_id == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	data, err := c.service.Upload(files, media_id, ctx)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "上传成功",
		"result": gin.H{
			"list": data,
		},
	})
}

func (c *controller) List(ctx *gin.Context) {
	media_id, _ := strconv.Atoi(ctx.Query("media_id"))
	//page := ctx.Query("page")
	//size := ctx.Query("size")
	list := c.service.List(media_id)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询成功",
		"result": gin.H{
			"list": list,
		},
	})
}

func (c *controller) Get(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	fmt.Println(id)
	data, content, err := c.service.Get(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询成功",
		"result": gin.H{
			"info":    data,
			"content": content,
		},
	})
}

func (c *controller) Update(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	content := ctx.PostForm("content")
	data := model.HlsModel{
		Id:   id,
		Path: ctx.PostForm("path"),
	}
	if err := c.service.Update(&data, content); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "修改成功",
	})
}
