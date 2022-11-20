package v2

import (
	"github.com/chamhaw/go-tools/gf-tool/v2/gerror"
	"github.com/chamhaw/go-tools/utils"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"net/http"
	"reflect"
)

type DefaultHandlerResponse struct {
	Success    bool        `json:"success" dc:"是否成功，和 http status 保持一致"`
	Code       string      `json:"code" dc:"文字描述的 code"`
	Message    string      `json:"message" dc:"错误信息"`
	Properties interface{} `json:"properties" dc:"额外信息"`
}

func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	if err != nil {
		er, ok := err.(gerror.BizError)
		if ok {
			r.Response.Status = utils.GetOrDefault(er.Code().HttpStatus(), http.StatusBadRequest)
			r.Response.WriteJson(DefaultHandlerResponse{
				Code:       er.Code().Code(),
				Message:    er.Message(),
				Properties: er.Detail(),
			})
		} else {
			r.Response.Status = http.StatusInternalServerError
			glog.Error(gctx.New(), "Runtime error.", err)
			r.Response.WriteJson(DefaultHandlerResponse{
				Code:    gerror.CodeInternalError.Code(),
				Message: err.Error(),
			})
		}
	} else if res != nil && !reflect.ValueOf(res).IsNil() {
		r.Response.WriteJson(res)
	} else {
		r.Response.WriteJson(DefaultHandlerResponse{
			Success: r.Response.Status >= http.StatusOK && r.Response.Status < http.StatusMultipleChoices,
			Code:    http.StatusText(r.Response.Status),
			Message: http.StatusText(r.Response.Status),
		})
	}

}

func JsonUtf8Middleware(r *ghttp.Request) {
	r.Middleware.Next()
	// 中间件处理逻辑
	r.Response.Header().Set("Content-Type", "application/json;charset=utf-8")
}
