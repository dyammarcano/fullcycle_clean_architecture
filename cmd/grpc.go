package cmd

import (
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/adapter/grpc"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/repository"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/usecase"
	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	RunE: func(cmd *cobra.Command, args []string) error {
		orderRepo := repository.Must(repository.NewOrderPostgresRepository())
		orderServer := grpc.NewGrpcOrderServer(usecase.NewOrderUseCase(orderRepo))
		return orderServer.Start()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
