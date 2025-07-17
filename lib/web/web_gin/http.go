package web_gin

import (
	"errors"
	"github.com/Iori372552686/GoOne/lib/api/logger"
	"github.com/Iori372552686/GoOne/lib/api/logger/zap"
	"github.com/Iori372552686/GoOne/lib/web/rest"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
* @Description: Run gin start the server
* @param: http_port
* @param: mode
* @param: session_name
* @param: load_routers
* @return: error
* @Author: Iori
* @Date: 2022-02-28 11:27:27
**/
func RunGin(conf Config, load_routers func(router *gin.Engine)) error {
	if conf.Port <= 0 {
		return errors.New("port args err!")
	}
	router := gin.New()
	router.Use(rest.Cors())

	//mode
	switch conf.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 使用 Zap 替换默认日志中间件
	router.Use(
		ginzap.Ginzap(zap.ZapLoger, time.RFC3339, true), // 记录请求日志 [1,4](@ref)
		ginzap.RecoveryWithZap(zap.ZapLoger, true),      // 替换 gin.Recovery() [4](@ref)
	)

	//loadRoutes
	load_routers(router)
	go router.Run(conf.IP + ":" + strconv.Itoa(conf.Port))
	logger.Infof("------ Http Gin Server Running by %v:%v ------", conf.IP, conf.Port)
	return nil
}
