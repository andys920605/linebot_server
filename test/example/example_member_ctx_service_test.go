package example

import (
	"context"
	"fmt"
	"linebot/database"
	"linebot/infras"
	"linebot/infras/configs"
	"linebot/infras/logger"
	models_svc "linebot/models/service"
	ext "linebot/repository/externals"
	rep_interface "linebot/repository/interface"
	rep "linebot/repository/mongodb"
	svc "linebot/service"
	"linebot/utils"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	cfg           *configs.Config
	db            *mongo.Collection
	apiLogger     *logger.ApiLogger
	LinebotClient *linebot.Client
	imemberRep    rep_interface.IMemberRep
	ilinebotExt   rep_interface.ILinebotExt
)

func ExampleMemberSvc_Webhook() {
	initCommentRep()
	svc := svc.NewMemberSvc(imemberRep, ilinebotExt)
	want := createRandomEvents()
	err := svc.Webhook(want)
	fmt.Println(err)
	// Output:
	// <nil>
}

func ExampleMemberSvc_Broadcast() {
	initCommentRep()
	svc := svc.NewMemberSvc(imemberRep, ilinebotExt)
	want := &models_svc.BroadcastMessage{
		Message: "example-test-" + utils.RandomString(5),
	}
	err := svc.Broadcast(want)
	fmt.Println(err)
	// Output:
	// <nil>
}

func ExampleMemberSvc_GetUserMessages() {
	initCommentRep()
	svc := svc.NewMemberSvc(imemberRep, ilinebotExt)
	rsp, err := svc.GetUserMessages("U078ce5dad75b4154a4f420fa6df49bd2")
	fmt.Println(rsp)
	fmt.Println(err)
	/// Output:
	// &[{6403228158e90734e5b9e3f7 3f5d5f5ca815440181735c958366d3bc message active 2023-03-04 10:50:39.229 +0000 UTC 0xc0003ac2c0 {17744697317931 111 [] <nil> <nil>  <nil>} <nil> <nil> <nil> <nil> <nil> <nil> [] <nil> <nil>}]
	// <nil>
}

// region private function
func initCommentRep() {
	utils.ConfigPath = "example"
	cfgFile, err := configs.LoadConfig(utils.GetConfigPath())
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, _ = configs.ParseConfig(cfgFile)
	apiLogger = logger.NewApiLogger(cfg)
	options := &infras.Options{
		Ctx:    context.Background(),
		Info:   nil,
		Config: cfg,
		Logger: apiLogger,
	}
	// Linebot
	LinebotClient, _ = database.NewLinebot(options)
	// MongoDB
	db, _ = database.NewDb(options)
	// Repository
	imemberRep = rep.NewMemberRep(db)
	// Externals
	ilinebotExt = ext.NewLinebotExt(LinebotClient)
}

func createRandomEvents() []*linebot.Event {
	var target []*linebot.Event
	event := linebot.Event{
		ReplyToken: "test-" + utils.RandomString(20),
		Type:       "message",
		Mode:       "active",
		Timestamp:  time.Now(),
		Source: &linebot.EventSource{
			Type:   "user",
			UserID: utils.RandomString(20),
		},
		Message: linebot.NewTextMessage(fmt.Sprintf("User ID: %v", utils.RandomString(20))),
	}
	target = append(target, &event)
	return target
}

// endregion
