package utils

import "github.com/gin-gonic/gin"

const statusOK = 200

const (
	OK_CODE        = 0
	ERROR_CODE     = 1
	DATAEXIST_CODE = 4003
)

type Response struct {
	Ctx     *gin.Context `json:"-"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Result  interface{}  `json:"result,omitempty"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (resp *Response) Error(message string) {
	resp.Code = 1
	resp.Message = message
	resp.Ctx.JSON(statusOK, resp)
}

func (resp *Response) Success(data any) {
	resp.Code = 0
	resp.Message = "成功"
	if data != nil {
		resp.Result = data
	}
	resp.Ctx.JSON(statusOK, resp)
}

func (resp *Response) List(list any, num int, size int, total int64) {
	resp.Success(gin.H{
		"list": list,
		"page": map[string]any{
			"num":   num,
			"size":  size,
			"total": total,
		},
	})
}
