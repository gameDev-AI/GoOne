package main

import (
	"fmt"
	"github.com/Iori372552686/GoOne/common/gamedata"
	"github.com/Iori372552686/GoOne/common/gconf"
	"github.com/Iori372552686/GoOne/lib/api/logger"
	"github.com/Iori372552686/GoOne/lib/service/sensitive_words"
	"github.com/Iori372552686/GoOne/lib/util/marshal"
	"github.com/Iori372552686/GoOne/lib/web/web_gin"
	"github.com/Iori372552686/GoOne/module/misc"
	"github.com/Iori372552686/GoOne/src/web_svr/cmd_handler"
	"github.com/Iori372552686/GoOne/src/web_svr/controller"
	"github.com/Iori372552686/GoOne/src/web_svr/globals"

	"log"
	"net/http"
	"runtime"
)

// gameSvr mainloop struct
type AppSvrImpl struct{}

/**
* @Description:init
* @return: error
* @Author: Iori
* @Date: 2022-04-27 21:04:30
**/
func (self *AppSvrImpl) OnInit() error {
	//-- set sys args
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)

	//-- load cfg
	err := self.OnReload()
	if err != nil {
		logger.Fatalf("Failed to load config | %v", err)
		return err
	}

	// init zap logger
	if _, err = logger.InitLogger(gconf.WebSvrCfg.LogDir, gconf.WebSvrCfg.LogLevel, "websvr"); err != nil {
		return err
	}

	//-- open pprof
	if gconf.WebSvrCfg.Pprof {
		go func() {
			logger.Infof("pprof listen on :81%02d", misc.ServerType_WebSvr)
			log.Println(http.ListenAndServe(fmt.Sprintf(":81%02d", misc.ServerType_WebSvr), nil))
		}()
	}

	//-- init redis
	err = globals.RedisMgr.InitAndRun(gconf.WebSvrCfg.DbInstances)
	if err != nil {
		logger.Errorf("RedisMgr InitAndRun error !! err=%v", err)
		return err
	}

	//-- init orm cache in some table
	/*	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), define.MaxOrmLruCacheLimitNum)
		err = globals.OrmMgr.GetOrmEngine().MapCacher(&define.MallRoleInfo{}, cacher)
		if err != nil {
			logger.Errorf("init orm cache error !! err | %v ", err)
			return err
		}*/

	//-- init Sign Mgr
	globals.SignMgr.InitAndRun(gconf.WebSvrCfg.HTTPSigns)
	//-- init RestApi mgr
	globals.RestMgr.Init(gconf.WebSvrCfg.RestApiConf, globals.SignMgr)
	//-- init CmdHandler
	cmd_handler.CmdHandler = cmd_handler.NewRegCmdHandler()
	//-- init sensitive words
	sensitive_words.Init(gconf.WebSvrCfg.SensitiveWordsFile)

	//-- init http server
	err = web_gin.RunGin(gconf.WebSvrCfg.HttpServer, controller.LoadWebRoutes)
	if err != nil {
		logger.Errorf("Http Serivce Start error !! err=%v", err)
		return err
	}

	return err
}

/**
* @Description:  reload
* @return: error
* @Author: Iori
* @Date: 2022-04-27 21:04:41
**/
func (self *AppSvrImpl) OnReload() error {
	// load start_conf, game_xlc_cfg_data..
	err := marshal.LoadConfFile(*gconf.SvrConfFile, &gconf.WebSvrCfg)
	if err != nil {
		logger.Fatalf("Failed to load server config | %s", err)
		return err
	}
	logger.Infof("svr_conf: %+v", gconf.WebSvrCfg)

	//local loading gameconf
	if gconf.WebSvrCfg.GameDataDir != "" {
		logger.Infof("Loading local file by gameconf_dir: %v ", gconf.WebSvrCfg.GameDataDir)
		gamedata.InitLocal(gconf.WebSvrCfg.GameDataDir)
	}

	return nil
}

/**
* @Description:  proc
* @return: bool
* @Author: Iori
* @Date: 2022-04-27 21:05:01
**/
func (self *AppSvrImpl) OnProc() bool {
	// mainloop  proc
	return true
}

/**
* @Description: mainloop tick
* @param: lastMs
* @param: nowMs
* @Author: Iori
* @Date: 2022-04-27 21:04:53
**/
func (self *AppSvrImpl) OnTick(lastMs, nowMs int64) {
}

/**
* @Description: main exit
* @Author: Iori
* @Date: 2022-04-27 21:05:07
**/
func (self *AppSvrImpl) OnExit() {
	// game exit todo something
	logger.Infof("web_service exit, right now !")
	logger.Infof("================== AppSvrImpl Stop =========================")
	logger.Flush()
}
