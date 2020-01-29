package controller

import (
	"context"
	"github.com/ClareChu/tracing/client/grpc"
	"github.com/ClareChu/tracing/proto"
	"github.com/opentracing/opentracing-go"
)

type Student struct {
	Id int64 `json:"id"`
}

func (s *Student) Get(ctx context.Context) *proto.BaseResponse {
	span, ctx := opentracing.StartSpanFromContext(ctx, "controllerStudent")
	span.LogKV("message", "get grpc student...")
	resp, err := grpc.Get(ctx, s.Id)
	if err != nil {
		span.LogKV("error", err.Error())
	}
	defer span.Finish()
	return resp
}
