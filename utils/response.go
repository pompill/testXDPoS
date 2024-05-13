package utils

import "github.com/beego/beego/v2/server/web/context"

type SuccessResp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func Success(ctx *context.Context, data interface{}, total int64) {
	resp := SuccessResp{
		Code:  200,
		Msg:   "success",
		Data:  data,
		Total: total,
	}
	ctx.Output.SetStatus(200)
	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Output.JSON(resp, true, true)
}

func Error(ctx *context.Context, errCode int, errMsg string) {
	resp := SuccessResp{
		Code:  errCode,
		Msg:   errMsg,
		Data:  nil,
		Total: 0,
	}
	ctx.Output.SetStatus(400)
	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Output.JSON(resp, true, true)
}
