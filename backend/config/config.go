package config

import (
	"go_gin_mcis/pkg/logger"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置结构体
type Config struct {
	Server        ServerConfig
	Database      DatabaseConfig
	OracleDb      *OracleDbConfig
	App           AppConfig
	Features      FeaturesConfig
	NestedMap     map[string]map[string]map[string]string
	EndCaseServer *EndCaseServerConfig
}

type ServerConfig struct {
	Port int
	Mode string
}
type OracleDbConfig struct {
	User             string
	Password         string
	Host             string
	Port             int
	Sid              string
	InstantClientDir string
	Enable           bool
}
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type AppConfig struct {
	Name        string
	Version     string
	DisplayName string
	Description string
	Tags        []string
}

type FeaturesConfig struct {
	EnableCache bool
	RetryCount  int
}

type EndCaseServerConfig struct {
	URL       string
	Cron      string
	LoginNo   string
	Password  string
	BeginDate string
	EndDate   string
	Enable    bool
}

// Cfg 全局变量
var Cfg *Config
var mu sync.RWMutex // 热更新时保证并发安全
// OnChange 回调函数列表
var onChangeCallbacks []func()

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config") // ✅ 真正的配置目录
	//viper.AddConfigPath(".")        // 兜底

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("❌ 读取配置文件失败: %v", err)
	}

	loadConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("🔄 配置文件变更: %s", e.Name)
		loadConfig()
		for _, cb := range onChangeCallbacks {
			cb()
		}
	})
}

// loadConfig 使用 Unmarshal 自动映射到结构体
func loadConfig() {
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		// panic(fmt.Errorf("Unable to decode config into struct: %v", err))
		logger.Fatalf("❌ 解析配置文件失败: %v", err)
	}
	mu.Lock()
	Cfg = &conf
	mu.Unlock()
	logger.Infof("✅ 加载配置文件成功: %v", Cfg)
}

// GetConf 并发安全访问全局配置
func GetConf() *Config {
	mu.RLock()
	defer mu.RUnlock()
	return Cfg
}

// RegisterOnChange 注册配置变更回调
func RegisterOnChange(cb func()) {
	onChangeCallbacks = append(onChangeCallbacks, cb)
}
