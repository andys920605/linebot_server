package externals

import (
	models_svc "linebot/models/service"
	rep_interface "linebot/repository/interface"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LinebotExt struct {
	LinebotClient *linebot.Client
}

func NewLinebotExt(linebotClient *linebot.Client) rep_interface.ILinebotExt {
	return &LinebotExt{
		LinebotClient: linebotClient,
	}
}

// push message to user
func (ext *LinebotExt) PushMessage(userid string, param *linebot.TextMessage) error {
	if _, err := ext.LinebotClient.PushMessage(userid, param).Do(); err != nil {
		return err
	}
	return nil
}

// push message to user
func (ext *LinebotExt) BroadcastMessage(param *models_svc.BroadcastMessage) error {
	_, err := ext.LinebotClient.BroadcastMessage(linebot.NewTextMessage(param.Message)).Do()
	if err != nil {
		return err
	}
	return nil
}
