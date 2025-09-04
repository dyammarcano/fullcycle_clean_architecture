package grpc

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/dyammarcano/fullcycle_clean_architecture/internal/usecase"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/grpc/pb"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/parameters"
	"github.com/inovacc/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcOrders := make([]*pb.Order, 0, len(orders))
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
	cfg, err := config.GetServiceConfig[*parameters.Service]()
	if err != nil {
		log.Fatalf("Failed to get service config: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, s)

	// Enable gRPC server reflection for tools like evans or grpcurl
	reflection.Register(grpcServer)

	slog.Info("gRPC server is running on port", slog.String("port", fmt.Sprintf(":%d", cfg.Grpc.Port)))

	return grpcServer.Serve(lis)
}
