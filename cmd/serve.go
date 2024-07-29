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
}
