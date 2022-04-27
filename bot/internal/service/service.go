package service

import (
	"bytes"
	"context"
	"fmt"
	db2 "github.com/supperdoggy/taro-pizda/bot/internal/db"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

type IService interface {
	DailyTaro(userID int64) (interface{}, error)
}

type service struct {
	db db2.IDB
	l *zap.Logger
}

func NewService(d *db2.IDB, l *zap.Logger) IService {
	return &service{
		db: *d,
		l:  l,
	}
}

func (s *service) DailyTaro(userID int64) (interface{}, error) {
	ctx := context.Background()
	taro, err := s.db.GetRandomTaro(ctx)
	if err != nil {
		s.l.Error("error getting taro card", zap.Error(err))
		return nil, err
	}

	res := telebot.Photo{
		File:    telebot.FromReader(bytes.NewReader(taro.Pic.Data)),
		Caption: fmt.Sprintf("*Карта дня*: %s\n\n*Совет дня*: %s\n\n*Предостережение дня*: %s", taro.Loc.Value, taro.Advice.Value, taro.Warning.Value),
	}

	err = s.db.SaveDailyTaro(taro.ID, userID, ctx)
	if err != nil {
		s.l.Error("error saving daily taro", zap.Error(err))
		return nil, err
	}

	return res, nil
}
