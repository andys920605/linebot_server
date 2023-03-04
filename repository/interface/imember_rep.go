package repository_interface

import (
	"context"
	models_rep "linebot/models/repository"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

//go:generate mockgen -destination=../../test/mock/imember_mock_repository.go -package=mock linebot/repository/interface IMemberRep
type IMemberRep interface {
	Insert(context.Context, *linebot.Event) error
	FindAll(context.Context, string) (*[]models_rep.LineEvent, error)
}
