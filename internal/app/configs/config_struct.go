//Author: lxk20021217
//Date: 2022-06-17 23:33:31
//LastEditTime: 2022-06-19 00:34:45
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config_struct.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 23:33:31
//LastEditTime: 2022-06-19 00:42:02
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config_struct.go
//是谁总是天亮了才睡

package configs

type Config struct {
	Http    Http    `mapstructure:"http"`
	Log     Log     `mapstructure:"log"`
	Redis   Redis   `mapstructure:"redis"`
	Postgre Postgre `mapstructure:"postgre"`
}

type Http struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	Highfile string `mapstructure:"highfile"`
	Level    int    `mapstructure:"level"`
	Lowfile  string `mapstructure:"lowfile"`
}
type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type Postgre struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type Limit struct {
	Limit int `mapstructure:"limit"`
}
