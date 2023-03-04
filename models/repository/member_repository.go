package repository

import (
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Member struct {
	Name  string `json:"name,omitempty"`
	Age   string `json:"age,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type LineEvent struct {
	Id                string `json:"id" bson:"_id"`
	ReplyToken        string
	Type              linebot.EventType
	Mode              linebot.EventMode
	Timestamp         time.Time
	Source            *linebot.EventSource
	Message           linebot.TextMessage
	Joined            *linebot.Members
	Left              *linebot.Members
	Postback          *linebot.Postback
	Beacon            *linebot.Beacon
	AccountLink       *linebot.AccountLink
	Things            *linebot.Things
	Members           []*linebot.EventSource
	Unsend            *linebot.Unsend
	VideoPlayComplete *linebot.VideoPlayComplete
}
