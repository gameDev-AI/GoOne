package controller

import (
	"errors"
	"fmt"
	"github.com/Iori372552686/GoOne/common/gconf"
	"github.com/Iori372552686/GoOne/lib/api/http_sign"
	"github.com/Iori372552686/GoOne/lib/api/logger"
	"github.com/Iori372552686/GoOne/lib/web/rest"
	"github.com/Iori372552686/GoOne/src/web_svr/cmd_handler/dispach"
	define "github.com/Iori372552686/GoOne/src/web_svr/common"
	"github.com/Iori372552686/GoOne/src/web_svr/globals"
	"github.com/gin-gonic/gin"
	"io"
)

/*
*  WebApiController
*  @Description:
 */
type WebApiController struct {
	rest.Controller
}

/**
* @Description:  AOP pre
* @receiver: self
* @return: gin.HandlerFunc
* @Author: Iori
* @Date: 2022-01-28 11:40:41
**/
func (self *WebApiController) before() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyBytes, _ := io.ReadAll(ctx.Request.Body)
		ctx.Set("body", bodyBytes)

		// sign
		if self.Auth {
			ok, err, _ := globals.SignMgr.GetSignIns().CheckSign(http_sign.UriParam2Map(ctx.Request.URL.RawQuery), bodyBytes, "")
			if ok {
				ctx.Next()
			} else {
				logger.Errorf("CheckSign fail, url_args | %v | err: %v", ctx.Request.RequestURI, err.Error())
				rest.ResultFail(ctx, fmt.Sprintf("Invalid signature ! err | %v", err.Error()))
				ctx.Abort()
			}
		}

		return
	}
}

/**
* @Description: init
* @Author: Iori
* @Date: 2022-07-26 17:50:15
**/
func (self *WebApiController) Init(router *gin.Engine) error {
	if router == nil {
		return errors.New("gin router is nil!")
	}

	self.Auth = gconf.WebSvrCfg.HttpServer.AuthEnable
	self.router(router)
	return nil
}

/**
* @Description: add Router
* @receiver: self
* @param: router
* @Author: Iori
* @Date: 2022-01-28 11:40:28
**/
func (self *WebApiController) router(router *gin.Engine) {
	rg := router.Group(define.RestApi_SafeMsg_Dir)
	rg.Use(self.before())

	//post
	rg.POST("/:cmd", dispatch.CmdHandlerProcess)
	rg.GET("/:cmd", dispatch.CmdHandlerProcess)
}
