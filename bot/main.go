package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/supperdoggy/taro-pizda/bot/internal/config"
	db2 "github.com/supperdoggy/taro-pizda/bot/internal/db"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	cfg := config.GetConfig()
	ctx := context.Background()

	db, err := db2.NewDB(logger, cfg.DBUrl, cfg.DBName, cfg.WarningCollectionName, cfg.AdviceCollectionName, cfg.PicCollectionName, cfg.RuLocCollectionName, ctx)
	if err != nil {
		logger.Fatal("error setting db", zap.Error(err), zap.Any("cfg", cfg))
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token: cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 1*time.Millisecond},
	})
	if err != nil {
		logger.Fatal("error connecting to bot", zap.Error(err))
	}

	bot.Handle("/start", func(m *telebot.Message) {
		ctx := context.Background()
		taro, err := db.GetRandomTaro(ctx)
		if err != nil {
			logger.Error("error getting taro card", zap.Error(err))
			bot.Reply(m, "error")
			return
		}

		res := telebot.Photo{
			File:    telebot.FromReader(bytes.NewReader(taro.Pic.Data)),
			Caption: fmt.Sprintf("Карта дня: %s\n\nПредсказание на день: %s\nПредостережение дня: %s", taro.Loc.Value, taro.Advice.Value, taro.Warning.Value),
		}
		m, err = bot.Reply(m, &res)
		if err != nil {
			logger.Error("error sending", zap.Error(err))
		}
	})


	bot.Start()
}
