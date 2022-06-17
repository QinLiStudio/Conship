package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	C    = new(Config)
	once sync.Once
)

func MustLoad(fpath ...string) {
	once.Do(func() {
		// 文件信息
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		// 文件位置 (文件夹)
		viper.AddConfigPath("./configs")
		// 读取配置信息
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Read config failed: ", err)
		}
		// 解析至结构体
		if err := viper.Unmarshal(&C); err != nil {
			log.Fatal("Unmarshal config failed: ", err)
		}
	})
}

type Config struct {
	RunMode     string
	StorageMode string
	HTTP        HTTP
	Log         Log
	Root        Root
	Redis       Redis
	RateLimiter RateLimiter
	CORS        CORS
	Gorm        Gorm
	Postgres    Postgres
	Storage     Storage
}

type HTTP struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	ShutdownTimeout    int    `mapstructure:"shutdown_timeout"`
	MaxContentLength   int    `mapstructure:"max_content_length"`
	MaxReqLoggerLength int    `mapstructure:"max_req_logger_length"`
}

type Log struct {
	Level int `mapstructure:"level"`
}

type Root struct {
	UserID   int    `mapstructure:"user_id"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	RealName string `mapstructure:"real_name"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type RateLimiter struct {
	Enable  bool `mapstructure:"enable"`
	Count   int  `mapstructure:"count"`
	RedisDB int  `mapstructure:"redis_db"`
}

type CORS struct {
	Enable           bool     `mapstructure:"enable"`
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

type Gorm struct {
	Debug             bool   `mapstructure:"debug"`
	DBType            string `mapstructure:"db_type"`
	MaxLifetime       int    `mapstructure:"max_lifetime"`
	MaxOpenConns      int    `mapstructure:"max_open_conns"`
	MaxIdleConns      int    `mapstructure:"max_idle_conns"`
	TablePrefix       string `mapstructure:"table_prefix"`
	EnableAutoMigrate bool   `mapstructure:"enable_auto_migrate"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Storage struct {
	Directory string `mapstructure:"directory"`
}

func (p Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		p.Host, p.Port, p.User, p.DBName, p.Password, p.SSLMode)
}
