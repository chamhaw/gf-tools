package v2

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"heguan-code.urbanic.com/infrastructure/gf-tools/v2/gerror"
	"heguan-code.urbanic.com/infrastructure/go-tools/utils"
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
	// 有错误
	if err != nil {
		er, ok := err.(gerror.BizError)
		// 如果是我们业务主动抛出，则按业务错误处理
		if ok {
			r.Response.Status = utils.GetOrDefault(er.Code().HttpStatus(), http.StatusBadRequest)
			r.Response.WriteJson(DefaultHandlerResponse{
				Code:       er.Code().Code(),
				Message:    er.Message(),
				Properties: er.Detail(),
			})
			return
		}

		// 其余的一律按系统错误处理
		r.Response.Status = http.StatusInternalServerError
		glog.Error(gctx.New(), "Runtime error.", err)
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gerror.CodeInternalError.Code(),
			Message: err.Error(),
		})
		return

	}

	// 如果 response 不为空
	if res != nil {
		v := reflect.ValueOf(res)
		if !v.IsNil() {
			r.Response.WriteJson(res)
			return
		}

		switch reflect.TypeOf(v.Interface()).Kind() {
		case reflect.Array:
			r.Response.WriteJson(g.Array{})
			return
		case reflect.Slice:
			r.Response.WriteJson(g.Slice{})
			return
		case reflect.Map:
			r.Response.WriteJson(g.Map{})
			return
		default:
			// nothing
		}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Success: r.Response.Status >= http.StatusOK && r.Response.Status < http.StatusMultipleChoices,
		Code:    http.StatusText(r.Response.Status),
		Message: http.StatusText(r.Response.Status),
	})

}

func JsonUtf8Middleware(r *ghttp.Request) {
	r.Middleware.Next()
	// 中间件处理逻辑
	r.Response.Header().Set("Content-Type", "application/json;charset=utf-8")
}
