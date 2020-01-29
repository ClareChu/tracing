package grpc

import (
	"context"
	"github.com/ClareChu/tracing/proto"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":7575"
)

type server struct {
	proto.StudentServiceServer
} //服务对象

// GetStudent 实现服务的接口 在proto中定义的所有服务都是接口
func (s *server) Get(ctx context.Context, in *proto.StudentDTO) (br *proto.BaseResponse, err error) {
	br = &proto.BaseResponse{}
	return
}

func NewConfig(tracer opentracing.Tracer) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize the gRPC server.
	s := grpc.NewServer(
	grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	/*s := grpc.NewServer(
	grpc.StreamInterceptor(StreamServerInterceptor),
	grpc.UnaryInterceptor(UnaryServerInterceptor)) //起一个服务*/
	proto.RegisterStudentServiceServer(s, &server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before handling. Info: %+v", info)
	resp, err := handler(ctx, req)
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

// StreamServerInterceptor is a gRPC server-side interceptor that provides Prometheus monitoring for Streaming RPCs.
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before handling. Info: %+v", info)
	err := handler(srv, ss)
	log.Printf("after handling. err: %v", err)
	return err
}
