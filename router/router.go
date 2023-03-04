package router

import (
	models_svc "linebot/models/service"
	svc "linebot/service/interface"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// interface to instance gin.Engine
type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	MemberSvc     svc.IMemberSvc
	LinebotClient *linebot.Client
}

func NewRouter(IMemberSvc svc.IMemberSvc, Linebot *linebot.Client) IRouter {
	return &Router{
		MemberSvc:     IMemberSvc,
		LinebotClient: Linebot,
	}
}

// set router
func (router *Router) InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/callback", router.webhook)
	r.POST("/broadcast", router.broadcast)
	r.GET("/user/:userId/messages", router.getUserMessages)
	return r
}

// region private function

// webhook router
func (router *Router) webhook(c *gin.Context) {
	events, err := router.LinebotClient.ParseRequest(c.Request)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if errRsp := router.MemberSvc.Webhook(events); errRsp != nil {
		c.JSON(errRsp.StatusCode, errRsp.Message)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// broadcast router
func (router *Router) broadcast(c *gin.Context) {
	var payload models_svc.BroadcastMessage
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if errRsp := router.MemberSvc.Broadcast(&payload); errRsp != nil {
		c.JSON(errRsp.StatusCode, errRsp.Message)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// get user messages router
func (router *Router) getUserMessages(c *gin.Context) {
	userId := c.Param("userId")
	rsp, errRsp := router.MemberSvc.GetUserMessages(userId)
	if errRsp != nil {
		c.JSON(errRsp.StatusCode, errRsp.Message)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// endregion
