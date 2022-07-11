//Author: lxk20021217
//Date: 2022-06-17 16:49:45
//LastEditTime: 2022-07-02 15:26:58
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\service\service.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:49:45
//LastEditTime: 2022-07-11 15:29:15
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\service\service.go
//是谁总是天亮了才睡

package service

import (
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/pkg/error"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, file *multipart.FileHeader, filename string, i schema.Item) {
	if err := configs.POSTGRESDB.Select("url", "secret").Create(&i).Error; err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件上传失败。", err)
		return
	}

	if err := c.SaveUploadedFile(file, filename); err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件上传失败。", err)
		return
	}
	error.Response(c, error.OK, gin.H{"url": i.Url, "secret": i.Secret}, "配置文件上传成功。")
}

func Delete(c *gin.Context, s string) {
	var i schema.Item

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Find(&i).Error; err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件删除失败。", err)
		return
	}

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Delete(&i).Error; err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件删除失败。", err)
		return
	}

	if err := os.Remove("./file/" + i.Url); err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件删除失败。", err)
		return
	}

	error.Response(c, error.OK, gin.H{"url": i.Url, "secret": i.Secret}, "配置文件删除成功。")
}

func Update(c *gin.Context, s string, file *multipart.FileHeader, filename string, url string) {
	var i schema.Item

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Find(&i).Error; err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件修改失败。", err)
		return
	}

	if err := os.Remove("./file/" + i.Url); err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件修改失败。", err)
		return
	}

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Update("url", url).Error; err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件修改失败。", err)
		return
	}

	if err := c.SaveUploadedFile(file, filename); err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件修改失败。", err)
		return
	}
	error.Response(c, error.OK, gin.H{"url": url, "secret": s}, "配置文件修改成功。")
}

func SearchUrl(c *gin.Context, u string) {
	filename := u
	file, err := os.Open("./file" + filename)
	if err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件查找失败。", err)
		file.Close()
		return
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件查找失败。", err)
		return
	}
	error.Response(c, error.OK, gin.H{"content": string(content)}, "配置文件查找成功。")
}
func SearchSecret(c *gin.Context, s string) {
	var i schema.Item
	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Find(&i).Error; err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件查找失败。", err)
		return
	}

	filename := i.Url
	file, err := os.Open("./file/" + filename)
	if err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件查找失败。", err)
		file.Close()
		return
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件查找失败。", err)
		return
	}
	error.Response(c, error.OK, gin.H{"content": string(content)}, "配置文件查找成功。")
}
