package di

import (
	"context"
	"linebot/database"
	"linebot/infras"
	"linebot/infras/logger"
	"linebot/models/commons"
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
	// MongoDB
	db, _ := database.NewDb(options)
	// Repository
	memberRep := rep.NewMemberRep(db)
	// Service
	memberSvc := svc.NewMemberSvc(memberRep)
	// Router
	router := router.NewRouter(memberSvc)
	// Server
	q1Server := app.NewLineBotServer(options, router)
	return q1Server, nil
}
