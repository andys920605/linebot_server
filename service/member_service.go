package service

import (
	models_rep "linebot/models/repository"
	rep "linebot/repository/interface"
	svc_interface "linebot/service/interface"
	"linebot/utils/errs"
)

var (
//cancelTimeout time.Duration = 3 // default 3 second
)

type MemberSvc struct {
	MemberRep rep.IMemberRep
}

func NewMemberSvc(IMemberRep rep.IMemberRep) svc_interface.IMemberSvc {
	return &MemberSvc{
		MemberRep: IMemberRep,
	}
}

func (svc *MemberSvc) CreateMember(param *models_rep.Member) (*models_rep.Member, *errs.ErrorResponse) {
	// ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	// defer cancel()
	// param.Uuid = uuid.NewV4().String()
	// if errRsp := svc.MemberRep.Insert(ctx, param); errRsp != nil {
	// 	return nil, &errs.ErrorResponse{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    errRsp.Error(),
	// 	}
	// }
	return param, nil
}
func (svc *MemberSvc) GetMember(uuid string) (*models_rep.Member, *errs.ErrorResponse) {
	// ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	// defer cancel()
	// result, errRsp := svc.MemberRep.Find(ctx, uuid)
	// if errRsp != nil {
	// 	return nil, &errs.ErrorResponse{
	// 		StatusCode: http.StatusNotFound,
	// 		Message:    errRsp.Error(),
	// 	}
	// }
	return nil, nil
}
func (svc *MemberSvc) UpdateMember(param *models_rep.Member) (*models_rep.Member, *errs.ErrorResponse) {
	// ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	// defer cancel()
	// errRsp := svc.MemberRep.Updates(ctx, param)
	// if errRsp != nil {
	// 	return nil, &errs.ErrorResponse{
	// 		StatusCode: http.StatusNotFound,
	// 		Message:    errRsp.Error(),
	// 	}
	// }
	// result, errRsp := svc.MemberRep.Find(ctx, param.Uuid)
	// if errRsp != nil {
	// 	return nil, &errs.ErrorResponse{
	// 		StatusCode: http.StatusNotFound,
	// 		Message:    errRsp.Error(),
	// 	}
	// }
	return nil, nil
}
func (svc *MemberSvc) DeleteMember(uuid string) *errs.ErrorResponse {
	// ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	// defer cancel()
	// errRsp := svc.MemberRep.Delete(ctx, uuid)
	// if errRsp != nil {
	// 	return &errs.ErrorResponse{
	// 		StatusCode: http.StatusNotFound,
	// 		Message:    errRsp.Error(),
	// 	}
	// }
	return nil
}
