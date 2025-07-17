package controller

import (
	"github.com/gin-gonic/gin"
)

// load web router
func LoadWebRoutes(router *gin.Engine) {
	//globals.PromMgr.SetGinMidAndRouter(router) // add mid and router

	new(WebApiController).Init(router)
}
