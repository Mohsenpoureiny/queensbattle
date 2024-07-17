/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"queensbattle/internal/repository"
	"queensbattle/internal/repository/redis"
	"queensbattle/internal/service"
	"queensbattle/internal/telegram"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	_ = godotenv.Load()
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to redis server")
	}
	accountRepository := repository.NewAccountRedisRepository(redisClient)

	// set up app
	app := service.NewApp(
		service.NewAccountService(accountRepository),
	)
	tg, err := telegram.NewTelegram(app, os.Getenv("BOT_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to telegram server")
	}
	tg.Start()
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
