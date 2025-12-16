package main

import (
	"fmt"
	"go_gin_mcis/cmd"
	"go_gin_mcis/config"
	"os"
	"path/filepath"
	"sync"

	"github.com/kardianos/service"
)

func init() {
	// 切换工作目录到 exe 所在路径
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	os.Chdir(dir)
	fmt.Println("📁 工作目录:", dir)
}

type program struct {
	sync.Mutex
}

func (p *program) Start(s service.Service) error {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🔥 服务 panic recovered: %v\n", r)
			}
		}()
		cmd.Start()
	}()
	fmt.Println("✅ 服务主逻辑启动成功")
	return nil
}

func (p *program) Stop(s service.Service) error {
	fmt.Println("🛑 服务停止")
	return nil
}

func main() {
	// 提前初始化配置
	config.Init() // 现在配置在 cmd.Start 里也会再次调用，可以去掉重复调用
	cfg := config.GetConf()
	svcConfig := &service.Config{
		Name:        cfg.App.Name,
		DisplayName: cfg.App.DisplayName,
		Description: cfg.App.Description,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("❌ 创建服务失败:", err)
		return
	}

	// 有参数时执行控制命令
	if len(os.Args) > 1 {
		action := os.Args[1]
		fmt.Println("⚙️  执行操作:", action)
		err := service.Control(s, action)
		if err != nil {
			fmt.Println("❌ 操作失败:", err)
		} else {
			fmt.Println("✅ 操作成功:", action)
		}
		return
	}

	fmt.Println("🚀 运行服务中...")
	err = s.Run()
	if err != nil {
		fmt.Println("❌ 服务运行失败:", err)
	}
}
