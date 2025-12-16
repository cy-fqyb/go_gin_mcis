package bootstrap

import (
	"go_gin_mcis/config"
	"go_gin_mcis/internal/service"
	"go_gin_mcis/pkg/db"
	"go_gin_mcis/pkg/logger"
	"sync"
)

func startEndCaseScheduler(conf *config.Config) error {
	scheduler := service.NewEndCaseScheduler(
		db.DB,
		conf.EndCaseServer.URL,
	)

	var mu sync.Mutex
	start := func() error {
		mu.Lock()
		defer mu.Unlock()
		return scheduler.Reload(conf.EndCaseServer.Cron)
	}

	if err := start(); err != nil {
		return err
	}

	config.RegisterOnChange(func() {
		logger.Infof("🔄 配置变更，重启 EndCaseScheduler")
		if err := start(); err != nil {
			logger.Errorf("EndCaseScheduler reload failed: %v", err)
		}
	})

	return nil
}
