package initialize

import (
	"log"
	"os"
	"time"

	"wally/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initMysql() {
	mysqlConfig := global.ServerConfig.MySQLConfig

	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlConfig.Dsn,
		DefaultStringSize: 100,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		zap.S().Panic(err)
	}

	sqlDB, err := global.DB.DB()
	if err != nil {
		zap.S().Panic(err)
	}
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.MaxConnLifeTime) * time.Second)

	if err != nil {
		zap.S().Panic(err)
	}
}
