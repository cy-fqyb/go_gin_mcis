package db

import (
	"fmt"
	"go_gin_mcis/pkg/logger"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	// 数据库配置
	server := "127.0.0.1"
	port := 11433
	user := "sa"
	password := "Cy2466579213!"
	database := "fiber_test"

	// 构造连接字符串
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		user, password, server, port, database)

	// 连接数据库
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
	})
	if err != nil {
		logger.Fatalf("❌ 连接数据库失败: %v", err)
	}

	logger.Info("✅ 成功连接到 SQL Server")
}

func InitOrcalDB() {
	// 数据库配置
	server := "192.168.1.241"
	port := 1521
	user := "sys"
	password := "123456"
	database := "orcl"

	// 构造连接字符串
	dsn := fmt.Sprintf("oracle://%s:%s@%s:%d/%s",
		user, password, server, port, database)

	// 连接数据库
	var err error
	DB, err = gorm.Open(oracle.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
	})
	if err != nil {
		logger.Fatalf("❌ 连接数据库失败: %v", err)
	}

	logger.Info("✅ 成功连接到 Oracle 数据库")
}
