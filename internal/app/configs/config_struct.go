//Author: lxk20021217
//Date: 2022-06-17 23:33:31
//LastEditTime: 2022-07-05 15:18:22
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config_struct.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 23:33:31
//LastEditTime: 2022-07-10 15:34:32
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config_struct.go
//是谁总是天亮了才睡

package configs

type Config struct {
	Mode     Mode     `mapstructure:"mode"`
	Http     Http     `mapstructure:"http"`
	Log      Log      `mapstructure:"log"`
	Redis    Redis    `mapstructure:"redis"`
	Postgres Postgres `mapstructure:"postgres"`
	Limit    Limit    `mapstructure:"limit"`
	Cors     Cors     `mapstructure:"cors"`
	Gorm     Gorm     `mapstructure:"gorm"`
}

type Mode struct {
	Runmode string `mapstructure:"runmode"`
}

type Http struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Log struct {
	Level  int    `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Outout string `mapstructure:"output"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func Dsn(r Postgres) string {
	return "host=" + r.Host + " user=" + r.User + " password=" + r.Password + " dbname=" + r.DBName + " port=" + r.Port + " sslmode=disable TimeZone=Asia/Shanghai"
}

type Limit struct {
	Limit int64 `mapstructure:"limit"`
}

type Cors struct {
	Enable           bool          `mapstructure:"enable"`
	AllowOrigins     []interface{} `mapstructure:"alloworigins"`
	AllowMethods     []interface{} `mapstructure:"allowmethods"`
	AllowHeaders     []interface{} `mapstructure:"allowheaders"`
	AllowCredentials bool          `mapstructure:"allowcredentials"`
	MaxAge           int           `mapstructure:"maxage"`
}

type Gorm struct {
	Debug  bool   `mapstructure:"debug"`
	DBType string `mapstructure:"dbtype"`
}
