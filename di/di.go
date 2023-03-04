package di

import (
	"context"
	"linebot/database"
	"linebot/infras"
	"linebot/infras/logger"
	"linebot/models/commons"
	ext "linebot/repository/externals"
	rep "linebot/repository/mongodb"
	"linebot/router"
	svc "linebot/service"

	"linebot/app"
)

// injection here
func CreateLinebotServer(ctx context.Context, info *commons.SystemInfo) (*app.LinebotServer, error) {
	config := infras.ProvideConfig()
	apiLogger := logger.NewApiLogger(config)
	options := &infras.Options{
		Ctx:    ctx,
		Info:   info,
		Config: config,
		Logger: apiLogger,
	}
	// Linebot
	linebot, _ := database.NewLinebot(options)
	// MongoDB
	db, _ := database.NewDb(options)
	// Repository
	memberRep := rep.NewMemberRep(db)
	// Externals
	linebotExt := ext.NewLinebotExt(linebot)
	// Service
	memberSvc := svc.NewMemberSvc(memberRep, linebotExt)
	// Router
	router := router.NewRouter(memberSvc, linebot)
	// Server
	q1Server := app.NewLineBotServer(options, router)
	return q1Server, nil
}
