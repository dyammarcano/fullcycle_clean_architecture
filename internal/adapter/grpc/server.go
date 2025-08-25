package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/dyammarcano/fullcycle_clean_architecture/internal/usecase"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/config"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/grpc/pb"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/logger"
	"google.golang.org/grpc"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer

	UseCase *usecase.OrderUseCase
}

func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.UseCase.ListOrders()
	if err != nil {
		return nil, err
	}

	grpcOrders := make([]*pb.Order, len(orders))
	for _, order := range orders {
		grpcOrders = append(grpcOrders, &pb.Order{
			Id:     int32(order.ID),
			Item:   order.Item,
			Amount: order.Amount,
		})
	}

	return &pb.ListOrdersResponse{Orders: grpcOrders}, nil
}

func NewGrpcOrderServer(useCase *usecase.OrderUseCase) *OrderServer {
	orderServer := &OrderServer{UseCase: useCase}
	return orderServer
}

func (s *OrderServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.G.Grpc.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, s)

	logger.Log(slog.LevelInfo, "gRPC server is running on port", slog.String("port", lis.Addr().String()))

	return grpcServer.Serve(lis)
}
