package controller

import (
	"context"
	"github.com/ClareChu/tracing/client/grpc"
	"github.com/ClareChu/tracing/proto"
	"github.com/opentracing/opentracing-go"
)

type Student struct {
}

func (s *Student) Get(ctx context.Context) *proto.BaseResponse {
	span, ctx := opentracing.StartSpanFromContext(ctx, "controllerStudent")
	span.LogKV("message", "get grpc student...")
	resp, err := grpc.Get(ctx)
	if err != nil {
		span.LogKV("error", err.Error())
	}
	defer span.Finish()
	return resp
}
