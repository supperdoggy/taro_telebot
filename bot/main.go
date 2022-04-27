package main

import (
	"context"
	"github.com/supperdoggy/taro-pizda/bot/internal/config"
	db2 "github.com/supperdoggy/taro-pizda/bot/internal/db"
	handlers2 "github.com/supperdoggy/taro-pizda/bot/internal/handlers"
	service2 "github.com/supperdoggy/taro-pizda/bot/internal/service"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	cfg := config.GetConfig()
	ctx := context.Background()

	db, err := db2.NewDB(logger, cfg.DBUrl, cfg.DBName,
		cfg.WarningCollectionName, cfg.AdviceCollectionName,
		cfg.PicCollectionName, cfg.RuLocCollectionName, cfg.DailyTaroCollectionName, ctx)
	if err != nil {
		logger.Fatal("error setting db", zap.Error(err), zap.Any("cfg", cfg))
	}
	service := service2.NewService(&db, logger)
	handlers := handlers2.NewHandlers(&service, logger)

	bot, err := telebot.NewBot(telebot.Settings{
		Token: cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 1*time.Millisecond},
	})
	if err != nil {
		logger.Fatal("error connecting to bot", zap.Error(err))
	}
	bot.Handle("/start", handlers.GetRandomDailyTaro)


	bot.Start()
}
