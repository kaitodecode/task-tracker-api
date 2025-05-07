package cmd

import (
	"github.com/kaitodecode/task-tracker-api/config"
	"github.com/kaitodecode/task-tracker-api/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/gin-gonic/gin"
)


var rootCmd = &cobra.Command{
	Use: "root",
	Short: "To run the project",
	Run: func(cmd *cobra.Command, args []string) {
		
		config.Init()
		db := database.InitDatabase()
		database.Migrate(db)

		
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		panic(err)
	}
}