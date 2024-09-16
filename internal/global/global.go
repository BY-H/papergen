package global

import (
	"cyclopropane/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB    // 数据库连接
	JWTKey string      // JWT密钥，每次重新启动的时候随机生成
	Logger *zap.Logger // 全局日志
)

func init() {
	Logger, _ = zap.NewProduction()
	JWTKey, _ = utils.GenerateRandomString(32)

}
