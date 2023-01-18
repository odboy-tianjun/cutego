package refs

import (
	redisTool "cutego/pkg/redispool"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// X 全局DB
var (
	SqlDB   *xorm.Engine
	RedisDB *redisTool.RedisClient
	CoolGin *gin.Engine
)
