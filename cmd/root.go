package cmd

import (
	"fmt"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfgFile, err := rootCmd.Flags().GetString("config")
	if err != nil {
		fmt.Println("Error getting config file")
		os.Exit(1)
	}

	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file")
		os.Exit(1)
	}

	if err = viper.Unmarshal(config.G); err != nil {
		fmt.Println("Error unmarshalling config")
		os.Exit(1)
	}

	if err = config.G.Validate(); err != nil {
		fmt.Println("Error validating config")
		os.Exit(1)
	}
}
