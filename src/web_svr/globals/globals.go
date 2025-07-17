// Package globals
// @Description: 全局服务管理器
package globals

import (
	"github.com/Iori372552686/GoOne/lib/api/http_sign"
	"github.com/Iori372552686/GoOne/lib/api/rest_api"
	"github.com/Iori372552686/GoOne/lib/db/redis"
	orm "github.com/Iori372552686/GoOne/lib/db/xorm"
)

var (
	OrmMgr   = orm.Orm_Mgr
	SignMgr  = http_sign.NewSignMgr()
	RestMgr  = rest_api.NewRestApiMgr()
	RedisMgr = redis.NewRedisMgr()
)
