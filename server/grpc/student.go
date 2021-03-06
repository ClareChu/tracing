package grpc

import (
	"context"
	"github.com/ClareChu/tracing/proto"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func RegistryStudent(server *grpc.Server) {
	proto.RegisterStudentServiceServer(server, &StudentGrpcService{})
}

type StudentGrpcService struct {
	proto.StudentServiceServer
} //服务对象

// GetStudent 实现服务的接口 在proto中定义的所有服务都是接口
func (s *StudentGrpcService) Get(ctx context.Context, in *proto.StudentDTO) (br *proto.BaseResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpcStudentServer")
	defer span.Finish()
	span.LogKV("message", "grpc student server start...")
	if in.Id == 1 {
		br = &proto.BaseResponse{
			Code:    0,
			Message: "success",
			Data: &proto.StudentDTO{
				Id:     1,
				TbAge:  10,
				TbName: "clare",
				TbSex:  1,
			},
		}
		return
	}
	br = &proto.BaseResponse{
		Code:    500,
		Message: "fail",
		Data:    in,
	}
	return
}
