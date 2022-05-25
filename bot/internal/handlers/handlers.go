package handlers

import (
	"github.com/supperdoggy/taro-pizda/bot/internal/service"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

type IHandlers interface {
	GetRandomDailyTaro(m *telebot.Message)
}

type handlers struct {
	service service.IService
	bot     *telebot.Bot

	logger *zap.Logger
}

func NewHandlers(s *service.IService, l *zap.Logger, b *telebot.Bot) IHandlers {
	return &handlers{
		service: *s,
		logger:  l,
		bot:     b,
	}
}

func (h *handlers) GetRandomDailyTaro(m *telebot.Message) {
	h.logger.Info("got message from", zap.Any("id", m.Sender.ID))
	result, err := h.service.DailyTaro(m.Sender.ID)
	if err != nil {
		h.logger.Error("error DailyTaro", zap.Error(err))
		h.bot.Reply(m, err.Error())
		return
	}

	m, err = h.bot.Reply(m, result, telebot.ModeMarkdown)
	if err != nil {
		h.logger.Error("error replying", zap.Error(err), zap.Any("taro", result))
	}
}
