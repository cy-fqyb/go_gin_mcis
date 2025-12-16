package bootstrap

import (
	"go_gin_mcis/config"
	"go_gin_mcis/pkg/logger"
)

type ServiceSwitch struct {
	Name   string
	Enable func(conf *config.Config) bool
	Start  func(conf *config.Config) error
}

func StartServices() {
	conf := config.GetConf()
	if conf == nil {
		logger.Fatal("config not initialized")
		return
	}

	services := []ServiceSwitch{
		{
			Name: "EndCaseScheduler",
			Enable: func(c *config.Config) bool {
				return c.EndCaseServer != nil && c.EndCaseServer.Enable
			},
			Start: startEndCaseScheduler,
		},
		{
			Name: "OracleDB",
			Enable: func(c *config.Config) bool {
				return c.OracleDb != nil && c.OracleDb.Enable
			},
			Start: startDb,
		},
	}

	for _, s := range services {
		if !s.Enable(conf) {
			logger.Infof("⏭ 服务未启用，跳过: %s", s.Name)
			continue
		}

		logger.Infof("🚀 启动服务: %s", s.Name)
		if err := s.Start(conf); err != nil {
			logger.Errorf("❌ 服务启动失败 [%s]: %v", s.Name, err)
			continue // 继续其他服务
		}
	}
}
