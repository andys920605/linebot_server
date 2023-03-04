package repository_interface

import (
	models_svc "linebot/models/service"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

//go:generate mockgen -destination=../../test/mock/ilinebot_mock_external.go -package=mock linebot/repository/interface ILinebotExt
type ILinebotExt interface {
	// sending message to user
	// @param userid
	// @param message
	// @return error
	PushMessage(string, *linebot.TextMessage) error

	// broadcast message
	// @param message
	// @return error
	BroadcastMessage(*models_svc.BroadcastMessage) error
}
