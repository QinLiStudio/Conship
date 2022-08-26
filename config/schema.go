/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 19:08:21
 * @Description: 配置模型
 */
package config

type Config struct {
	Mode     Mode
	Http     Http
	Redis    Redis
	Postgres Postgres
	Limit    Limit
	Cors     Cors
}

type Mode struct {
	RunMode string
}

type Http struct {
	Url string
	Host string
	Port string
}

type Redis struct {
	Addr     string
	Password string
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (r Postgres) Dsn() string {
	return "host=" + r.Host + " user=" + r.User + " password=" + r.Password + " dbname=" + r.DBName + " port=" + r.Port + " sslmode=disable TimeZone=Asia/Shanghai"
}

type Limit struct {
	Request int64
	Content int
}

type Cors struct {
	Enable bool

	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
}
