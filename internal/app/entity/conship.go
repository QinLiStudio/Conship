/*
 * @Date: 2022-06-17 11:57:58
 * @LastEditors: tich425
 * @LastEditTime: 2022-06-17 19:46:18
 * @FilePath: \Conship\internal\app\entity\conship.go
 * @Description: 数据库内容实例
 */
package entity

import (
	"gorm.io/gorm"
)

type CorrespondingTable struct {
	gorm.Model

	Link   string `json:"link" binding:"max=8"`
	Secret string `json:"key" binding:"max=8"`
}
