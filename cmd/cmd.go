package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute the application.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

var rootCmd = &cobra.Command{
	Short: "Hello, I'm Exodia Library",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("version"))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config/default.yaml)")
}

var configFile string

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("default")
		viper.AddConfigPath("./config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatal("Unable to read config from file")
	}
}
