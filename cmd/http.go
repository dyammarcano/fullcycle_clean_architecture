package cmd

import (
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/adapter/http"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/repository"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/usecase"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		orderRepo, err := repository.NewOrderPostgresRepository()
		if err != nil {
			return err
		}

		orderServer := http.NewHttpOrderServer(usecase.NewOrderUseCase(orderRepo))
		return orderServer.Start()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
