/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-07-08 20:55:49
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 21:04:52
 * @FilePath: \Conship\pkg\utils\refraction_struct\web_struct.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package refraction_struct

type Article struct {
	ID      uint   //ID
	Content string //内容
	Secret  string //密钥
}

//接口化
type ArticleService struct{}
