package db

import (
	"fmt"
	"go_gin_mcis/config"
	"go_gin_mcis/pkg/logger"
	"runtime"

	"github.com/oracle-samples/gorm-oracle/oracle"
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

func InitOracleDB() error {
	dsn := fmt.Sprintf(
		`user="%s" password="%s" connectString="%s:%d/%s"`,
		config.GetConf().OracleDb.User,
		config.GetConf().OracleDb.Password,
		config.GetConf().OracleDb.Host,
		config.GetConf().OracleDb.Port,
		config.GetConf().OracleDb.Sid,
	)

	// macOS / Windows 才加 libDir
	if runtime.GOOS != "linux" {
		dsn += fmt.Sprintf(` libDir="%s"`, config.GetConf().OracleDb.InstantClientDir)
	}

	var err error
	DB, err = gorm.Open(
		oracle.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)

	if err != nil {
		logger.Errorf("❌ Oracle 连接失败: %v", err)
		return err
	}

	logger.Info("✅ Oracle 数据库连接成功")
	return nil
}
