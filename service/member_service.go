package service

import (
	"context"
	"fmt"
	models_rep "linebot/models/repository"
	models_svc "linebot/models/service"
	rep "linebot/repository/interface"
	svc_interface "linebot/service/interface"
	"linebot/utils/errs"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	cancelTimeout time.Duration = 3 // default 3 second
)

type MemberSvc struct {
	MemberRep   rep.IMemberRep
	ILinebotExt rep.ILinebotExt
}

func NewMemberSvc(iMemberRep rep.IMemberRep, iLinebotExt rep.ILinebotExt) svc_interface.IMemberSvc {
	return &MemberSvc{
		MemberRep:   iMemberRep,
		ILinebotExt: iLinebotExt,
	}
}

// webhook event
func (svc *MemberSvc) Webhook(events []*linebot.Event) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	var err error
	for _, event := range events {
		if event.Type == linebot.EventTypeFollow {
			// push userid to new follower
			err = svc.ILinebotExt.PushMessage(event.Source.UserID, linebot.NewTextMessage(fmt.Sprintf("User ID: %v", event.Source.UserID)))
		}
		err = svc.MemberRep.Insert(ctx, event)
	}
	if err != nil {
		return &errs.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return nil
}

// broadcast
func (svc *MemberSvc) Broadcast(param *models_svc.BroadcastMessage) *errs.ErrorResponse {
	if err := svc.ILinebotExt.BroadcastMessage(param); err != nil {
		return &errs.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return nil
}

// get user messages
func (svc *MemberSvc) GetUserMessages(param string) (*[]models_rep.LineEvent, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	rsp, err := svc.MemberRep.FindAll(ctx, param)
	if err != nil {
		return nil, &errs.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return rsp, nil
}
