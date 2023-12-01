package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfig = "./config.json"
)

var (
	config string
	cmd    string
	src    string
	dest   string

	rootCmd = &cobra.Command{
		Use: "DaDah image decorator",
		Run: func(cmd *cobra.Command, args []string) {
			decorateCmd.Help()
		},
	}

	decorateCmd = &cobra.Command{
		Use:   "decorate",
		Short: "decorate the image",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("decorate")
			// fmt.Println(viper.AllSettings())

			settings, _ := json.Marshal(viper.AllSettings())
			fmt.Println(string(settings))
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(onInit)

	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is config.json)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().StringVar(&cmd, "version", "", "get version")
	viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))

	rootCmd.PersistentFlags().StringVar(&src, "source-dir", "./sources", "The folder for images ready to process")
	viper.BindPFlag("source-dir", rootCmd.PersistentFlags().Lookup("source-dir"))

	rootCmd.PersistentFlags().StringVar(&dest, "dest-dir", "./output", "The folder for processed images")
	viper.BindPFlag("dest-dir", rootCmd.PersistentFlags().Lookup("dest-dir"))

	rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.AddCommand(decorateCmd)
}

func onInit() {
	if config == "" {
		config = defaultConfig
		fmt.Printf("Use default config file: %s\n", defaultConfig)
	}

	viper.SetConfigFile(config)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
