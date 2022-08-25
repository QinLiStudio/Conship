/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 19:56:18
 * @Description: 数据库模型配置
 */
package meta

type Meta struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Url     string `json:"url"`
	Secret  string `json:"secret"`
	Content string
}
