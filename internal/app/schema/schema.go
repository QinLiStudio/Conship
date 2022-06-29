//Author: lxk20021217
//Date: 2022-06-17 16:48:19
//LastEditTime: 2022-06-24 23:55:25
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\schema\schema.go
//是谁总是天亮了才睡


package schema

type Item struct {
	ID     int		`gorm:"primaryKey" json:"id"`
	Url    string	`json:"url"`
	Secret string	`json:"secret"`
}
