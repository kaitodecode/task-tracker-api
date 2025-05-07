package cmd

import (
	"github.com/kaitodecode/task-tracker-api/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "root",
	Short: "To run the project",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init()
		config.InitDatabase()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		panic(err)
	}
}