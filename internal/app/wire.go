/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 22:09:51
 * @FilePath: \Conship\internal\app\wire.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"github.com/google/wire"
)

// func InitializeEvent() Event {
//     wire.Build(NewEvent, NewGreeter, NewMessage)
//     return Event{}
// }

//这个是简写，直接用 panic 包裹起来即可，因为不需要检查，所以不需要反馈
func InitializeEvent() Event {
	panic(wire.Build(NewEvent, NewMessage, NewGreeter))
}

func InitializeEvent() Event {
	panic(wire.Build(NewEvent, NewMessage, NewGreeter))
}

//写两个 一个Mysql 一个 redis 暂时无封装方法
