package global

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"papergen/config"
	"papergen/internal/db"
)

var (
	Env     string       = "test" // 工作环境（main, test） TODO 之后得整合到环境变量中
	WorkDir string                // 工作路径
	Conf    *config.Conf          // 配置文件
	DB      *gorm.DB              // 数据库连接
	JWTKey  string                // JWT密钥
	Logger  *zap.Logger           // 全局日志
)

func init() {
	WorkDir, _ = os.Getwd()
	Logger, _ = zap.NewProduction()
	Conf, _ = config.FromYaml(fmt.Sprintf("%s/%s-", WorkDir, Env))
	JWTKey = Conf.JWTKey
	DB, _ = db.InitDB(Conf.DatabaseHost, Conf.DatabaseUser, Conf.DatabasePassword)
}
