package grpc

import (
	"context"
	"github.com/ClareChu/tracing/proto"
	"github.com/opentracing/opentracing-go"
	"time"
)

func Get(ctx context.Context, id int64) (*proto.BaseResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	conn, err := NewConfig(span.Tracer())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := proto.NewStudentServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	dto := &proto.StudentDTO{
		Id: id,
	}
	response, err := c.Get(ctx, dto)
	return response, err
}
