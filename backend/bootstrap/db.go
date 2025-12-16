package bootstrap

import (
	"go_gin_mcis/config"
	"go_gin_mcis/pkg/db"
	"go_gin_mcis/pkg/logger"
)

func startDb(conf *config.Config) error {
	err := db.InitOracleDB()
	if err != nil {
		return err
	}
	config.RegisterOnChange(func() {
		logger.Infof("🔄 配置变更，重启 OracleDb")
		if err = db.InitOracleDB(); err != nil {
			logger.Errorf("OracleDb reload failed: %v", err)
		}
	})
	return nil
}
