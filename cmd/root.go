package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/parameters"
	"github.com/inovacc/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fullcycle_clean_architecture",
	Short: "Fullcycle Clean Architecture",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "config file (default is $binary_path/config.yaml)")
}

// initConfig reads in the config file and ENV variables if set.
func initConfig() {
	cfgFile, err := rootCmd.Flags().GetString("config")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stdout, "Error getting config file")
		os.Exit(1)
	}

	if err := config.InitServiceConfig(&parameters.Service{}, cfgFile); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
