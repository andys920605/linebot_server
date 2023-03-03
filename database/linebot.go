package database

import (
	"linebot/infras"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// new linebot instance
func NewLinebot(Opt *infras.Options) (*linebot.Client, error) {
	bot, err := linebot.New(Opt.Config.Linebot.ChannelSecret, Opt.Config.Linebot.ChannelAccessToken)
	if err != nil {
		Opt.Logger.Errorf("Error new linebot instance: %s", err)
	}

	return bot, nil
}
