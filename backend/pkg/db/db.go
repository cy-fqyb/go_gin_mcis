package db

import (
	"context"
	"database/sql"
	"fmt"
	"go_gin_mcis/pkg/logger"
	"time"

	_ "github.com/godror/godror"
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

var OracleDB *sql.DB

func InitOracleDB() {
	server := "172.31.66.145"
	port := 1521
	user := "mcis_kezhou"
	password := "mcis_kezhou"
	service := "orcl"

	dsn := fmt.Sprintf("%s/%s@%s:%d/%s", user, password, server, port, service)
	logger.Info("Oracle DSN: ", dsn)

	db, err := sql.Open("godror", dsn)
	if err != nil {
		logger.Fatalf("❌ 连接 Oracle 失败: %v", err)
	}

	// 连接池配置
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		logger.Fatalf("❌ Ping Oracle 失败: %v", err)
	}

	OracleDB = db
	logger.Info("✅ 成功连接到 Oracle 数据库")
}
