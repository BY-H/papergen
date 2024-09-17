package global

import (
	"cyclopropane/config"
	"cyclopropane/pkg/utils"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

var (
	Env     string       = "test" // 工作环境（main, test）
	WorkDir string                // 工作路径
	Conf    *config.Conf          // 配置文件
	DB      *gorm.DB              // 数据库连接
	JWTKey  string                // JWT密钥，每次重新启动的时候随机生成
	Logger  *zap.Logger           // 全局日志
)

func init() {
	WorkDir, _ = os.Getwd()
	Logger, _ = zap.NewProduction()
	Conf, _ = config.FromYaml(fmt.Sprintf("%s/%s-", WorkDir, Env))
	JWTKey, _ = utils.GenerateRandomString(32)
}
