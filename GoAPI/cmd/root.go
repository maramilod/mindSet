package cmd

import (
	"fmt"
	"mind-set/cmd/command"
	"mind-set/cmd/cron"
	"mind-set/cmd/server"
	"mind-set/config"
	"mind-set/internal/utils"
	"mind-set/internal/utils/logger"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "pos",
		Short:        "pos",
		SilenceUsage: true,
		Long: `Gin framework is used as the core of this project to build a scaffold, 
based on the project can be quickly completed business development, out of the box 📦`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config.InitConfig(configPath)
			logger.InitLogger()
		},
		Run: func(cmd *cobra.Command, args []string) {
			if printVersion {
				fmt.Println(utils.Version)
				return
			}

			fmt.Printf("%s\n", "Welcome to pos. Use -h to see more commands")
		},
	}
	configPath   string
	printVersion bool
)

func init() {

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "The absolute path of the configuration file")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", true, "GetUserInfo version info")
	rootCmd.AddCommand(server.Cmd)
	rootCmd.AddCommand(cron.Cmd)
	rootCmd.AddCommand(command.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
