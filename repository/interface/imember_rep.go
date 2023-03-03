package repository_interface

import (
	"context"
	models_rep "linebot/models/repository"
)

//go:generate mockgen -destination=../../test/mock/imember_mock_repository.go -package=mock linebot/repository/interface IMemberRep
type IMemberRep interface {
	Insert(context.Context, *models_rep.Member) error
	Find(context.Context, string) (*models_rep.Member, error)
	Update(context.Context, *models_rep.Member) error
	Delete(context.Context, string) error
}
