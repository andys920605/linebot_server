package service_test

import (
	"fmt"
	models_rep "linebot/models/repository"
	models_svc "linebot/models/service"
	svc "linebot/service"
	svc_interface "linebot/service/interface"
	"linebot/test/mock"
	"linebot/utils"
	"linebot/utils/errs"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
)

// mock repositories struct
type mockReps struct {
	mockiMemberRep  *mock.MockIMemberRep
	mockiLinebotExt *mock.MockILinebotExt
}

func TestMemberService_Webhook(t *testing.T) {
	t.Parallel()
	input := []*linebot.Event{}
	event := &linebot.Event{
		Type: "follow",
		Source: &linebot.EventSource{
			UserID: utils.RandomString(20),
		},
	}
	input = append(input, event)
	lineMessage := linebot.NewTextMessage(fmt.Sprintf("User ID: %v", event.Source.UserID))
	type args struct {
		arg []*linebot.Event
	}
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *errs.ErrorResponse
	}{
		{
			name: "webhook follow situation",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiLinebotExt.EXPECT().PushMessage(event.Source.UserID, lineMessage).Return(nil),
					f.mockiMemberRep.EXPECT().Insert(gomock.Any(), event).Return(nil),
				)
			},
			args: args{arg: input},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			memberSvc := newMemberService(f)
			got := memberSvc.Webhook(tt.args.arg)
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("MemberService.Webhook() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMemberService_Broadcast(t *testing.T) {
	t.Parallel()
	input := &models_svc.BroadcastMessage{
		Message: utils.RandomString(20),
	}
	type args struct {
		arg *models_svc.BroadcastMessage
	}
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *errs.ErrorResponse
	}{
		{
			name: "Broadcast",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiLinebotExt.EXPECT().BroadcastMessage(input).Return(nil),
				)
			},
			args: args{arg: input},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			memberSvc := newMemberService(f)
			got := memberSvc.Broadcast(tt.args.arg)
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("MemberService.Broadcast() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMemberService_GetUserMessages(t *testing.T) {
	t.Parallel()
	input := utils.RandomString(20)
	type args struct {
		arg string
	}
	rsp := []models_rep.LineEvent{}
	event := models_rep.LineEvent{
		Id:         utils.RandomString(20),
		ReplyToken: utils.RandomString(20),
		Type:       "message",
		Mode:       "active",
		Timestamp:  time.Now(),
		Source: &linebot.EventSource{
			Type:   "user",
			UserID: utils.RandomString(20),
		},
		Message: linebot.TextMessage{
			ID:   utils.RandomString(20),
			Text: utils.RandomString(30),
		},
	}
	rsp = append(rsp, event)
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *[]models_rep.LineEvent
	}{
		{
			name: "Get User Messages",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiMemberRep.EXPECT().FindAll(gomock.Any(), input).Return(&rsp, nil),
				)
			},
			args: args{arg: input},
			want: &rsp,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			memberSvc := newMemberService(f)
			if got, _ := memberSvc.GetUserMessages(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.UpdateComment() = %v, want = %v", got, tt.want)
			}
		})
	}
}

// region private function
// create mock repositories
// @param mockCtl controller
// @result MockReps model
func getMockReps(mockCtl *gomock.Controller) mockReps {
	return mockReps{
		mockiMemberRep:  mock.NewMockIMemberRep(mockCtl),
		mockiLinebotExt: mock.NewMockILinebotExt(mockCtl),
	}
}

// create new member service
// @param cfg config
// @param mockReps mock reps
// @result member service model
func newMemberService(mockReps mockReps) svc_interface.IMemberSvc {
	// iCommentSvc := svc.NewCommentSvc(mockReps.mockiCommentRep)
	return svc.NewMemberSvc(mockReps.mockiMemberRep, mockReps.mockiLinebotExt)
}

// endregion
