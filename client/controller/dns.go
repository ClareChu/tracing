package controller

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

type Dns struct {
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (d *Dns) Start(ctx context.Context) (rep *Response) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "startService")
	span.LogKV("message", "dns start...")
	defer span.Finish()
	return &Response{
		Code:    0,
		Message: "start",
	}
}

func (d *Dns) Done(ctx context.Context) (rep *Response) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "startService")
	defer span.Finish()
	span.LogKV("message", "dns Done...")
	return &Response{
		Code:    0,
		Message: "done",
	}
}
