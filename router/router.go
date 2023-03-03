package router

import (
	models_rep "linebot/models/repository"
	svc "linebot/service/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

// interface to instance gin.Engine
type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	MemberSvc svc.IMemberSvc
}

func NewRouter(ICommentSvc svc.IMemberSvc) IRouter {
	return &Router{
		MemberSvc: ICommentSvc,
	}
}

// set router
func (router *Router) InitRouter() *gin.Engine {
	r := gin.Default()
	g1 := r.Group("/v1/")
	g1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Comment
	g1.POST("/comment", router.createComment)
	g1.GET("/comment/:uuid", router.getComment)
	g1.PUT("/comment/:uuid", router.updateComment)
	g1.DELETE("/comment/:uuid", router.deleteComment)
	return r
}

// region CRUD Comment
func (router *Router) createComment(c *gin.Context) {
	var payload models_rep.Member
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": err.Error()})
		return
	}
	// result, errRsp := router.CommentSvc.CreateComment(&payload)
	// if errRsp != nil {
	// 	c.JSON(errRsp.StatusCode, errRsp)
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
func (router *Router) getComment(c *gin.Context) {
	// uuid := c.Param("uuid")
	// result, errRsp := router.CommentSvc.GetComment(uuid)
	// if errRsp != nil {
	// 	c.JSON(errRsp.StatusCode, errRsp)
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
func (router *Router) updateComment(c *gin.Context) {
	// uuid := c.Param("uuid")
	// var payload models_rep.Comment
	// payload.Uuid = uuid
	// if err := c.ShouldBind(&payload); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"description": err.Error()})
	// 	return
	// }
	// result, errRsp := router.CommentSvc.UpdateComment(&payload)
	// if errRsp != nil {
	// 	c.JSON(errRsp.StatusCode, errRsp)
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
func (router *Router) deleteComment(c *gin.Context) {
	// uuid := c.Param("uuid")
	// errRsp := router.CommentSvc.DeleteComment(uuid)
	// if errRsp != nil {
	// 	c.JSON(http.StatusInternalServerError, errRsp)
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"description": "ok"})
}

// endregion
