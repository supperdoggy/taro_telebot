package service

import (
	"bytes"
	"context"
	"fmt"

	db2 "github.com/supperdoggy/taro-pizda/bot/internal/db"
	"github.com/supperdoggy/taro-pizda/structs"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

type IService interface {
	DailyTaro(userID int64) (*telebot.Photo, error)
}

type service struct {
	db db2.IDB
	l  *zap.Logger
}

func NewService(d *db2.IDB, l *zap.Logger) IService {
	return &service{
		db: *d,
		l:  l,
	}
}

func (s *service) DailyTaro(userID int64) (*telebot.Photo, error) {
	var (
		ctx  = context.Background()
		taro structs.Taro
		err  error
	)
	can := s.db.CanGetNewDailyTaro(ctx, userID)
	if !can {
		var daily structs.DailyTaro
		daily, err = s.db.GetSavedDailyTaro(userID, ctx)
		if err != nil {
			return nil, err
		}

		taro, err = s.db.GetTaro(ctx, daily.CardID)
	} else {
		taro, err = s.db.GetRandomTaro(ctx)
		err = s.db.SaveDailyTaro(taro.ID, userID, ctx)
		if err != nil {
			s.l.Error("error saving daily taro", zap.Error(err))
			return nil, err
		}
	}

	if err != nil {
		s.l.Error("error getting taro from db", zap.Error(err))
		return nil, err
	}

	res := telebot.Photo{
		File:    telebot.FromReader(bytes.NewReader(taro.Pic.Data)),
		Caption: fmt.Sprintf("*Ваша карта дня*: %s\n\n*Совет дня*: %s\n\n*Предостережение дня*: %s", taro.Loc.Value, taro.Advice.Value, taro.Warning.Value),
	}

	return &res, nil
}
