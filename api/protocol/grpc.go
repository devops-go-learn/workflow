package protocol

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/workflow/conf"
)

// NewGRPCService todo
func NewGRPCService(interceptors ...grpc.UnaryServerInterceptor) *GRPCService {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptors...,
	)))

	return &GRPCService{
		svr: grpcServer,
		l:   zap.L().Named("GRPC Service"),
		c:   conf.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() error {
	// 装载所有GRPC服务
	app.LoadGrpcApp(s.svr)

	// 启动HTTP服务
	s.l.Infof("GRPC 开始启动, 监听地址: %s", s.c.GRPC.Addr())
	lis, err := net.Listen("tcp", s.c.GRPC.Addr())
	if err != nil {
		log.Fatal(err)
	}
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		return fmt.Errorf("start service error, %s", err.Error())
	}

	return nil
}

// Stop 停止GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Info("start graceful shutdown")

	// 优雅关闭HTTP服务
	s.svr.GracefulStop()

	return nil
}
