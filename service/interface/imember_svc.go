package service_interface

import (
	models_rep "linebot/models/repository"
	"linebot/utils/errs"
)

//go:generate mockgen -destination=../../test/mock/imember_mock_service.go -package=mock linebot/service/interface IMemberSvc
type IMemberSvc interface {
	CreateMember(*models_rep.Member) (*models_rep.Member, *errs.ErrorResponse)
	GetMember(string) (*models_rep.Member, *errs.ErrorResponse)
	UpdateMember(*models_rep.Member) (*models_rep.Member, *errs.ErrorResponse)
	DeleteMember(string) *errs.ErrorResponse
}
