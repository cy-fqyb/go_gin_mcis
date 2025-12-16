package db

import (
	"fmt"
	"go_gin_mcis/config"
	"go_gin_mcis/pkg/logger"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	// 构造连接字符串，直接兼容 TLS
	// encrypt=disable 表示不使用 TLS 加密
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		config.GetConf().Database.User,
		config.GetConf().Database.Password,
		config.GetConf().Database.Host,
		config.GetConf().Database.Port,
		config.GetConf().Database.Name,
	)

	// 连接数据库
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
	})
	if err != nil {
		logger.Errorf("❌ 连接数据库失败: %v", err)
		return
	}

	logger.Info("✅ 成功连接到 SQL Server (TLS 已兼容)")
}
