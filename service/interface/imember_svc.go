package service_interface

import (
	models_rep "linebot/models/repository"
	models_svc "linebot/models/service"
	"linebot/utils/errs"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

//go:generate mockgen -destination=../../test/mock/imember_mock_service.go -package=mock linebot/service/interface IMemberSvc
type IMemberSvc interface {
	Webhook([]*linebot.Event) *errs.ErrorResponse
	Broadcast(*models_svc.BroadcastMessage) *errs.ErrorResponse
	GetUserMessages(string) (*[]models_rep.LineEvent, *errs.ErrorResponse)
}
