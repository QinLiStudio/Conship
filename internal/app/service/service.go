//Author: lxk20021217
//Date: 2022-06-17 16:49:45
//LastEditTime: 2022-06-29 20:20:25
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\service\service.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:49:45
//LastEditTime: 2022-06-30 00:42:07
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\service\service.go
//是谁总是天亮了才睡

package service

import (
	"mime/multipart"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/pkg/error"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, file *multipart.FileHeader, filename string, i schema.Item) {
	if err := configs.POSTGRESDB.Select("url", "secret").Create(&i).Error; err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "配置文件上传失败。。", error.WithStack(err))
		return
	}

	if err := c.SaveUploadedFile(file, filename); err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "配置文件上传失败。", error.WithStack(err))
		return
	}
	error.Response(c, error.OK, gin.H{"url": i.Url, "secret": i.Secret}, "配置文件上传成功。")
}

func Delete(c *gin.Context, s string) {
	var i schema.Item

	if err := configs.POSTGRESDB.Table("item").Where("secret = ?", s).Delete(&i).Error; err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "删除配置文件失败。。", error.WithStack(err))
		return
	}
	error.Response(c, error.OK, gin.H{"url": i.Url, "secret": i.Secret}, "配置文件删除成功。")
}

func Update(c *gin.Context, s string, file *multipart.FileHeader, filename string, url string) {
	if err := configs.POSTGRESDB.Table("item").Where("secret = ?", s).Update("url", url).Error; err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "修改配置文件失败。。", error.WithStack(err))
		return
	}

	if err := c.SaveUploadedFile(file, filename); err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "配置文件修改失败。", error.WithStack(err))
		return
	}
	error.Response(c, error.OK, gin.H{"url": url, "secret": s}, "配置文件修改成功。")
}

func SearchUrl(c *gin.Context, u string) {
	filename := "./file/" + u

}
func SearchSecret(c *gin.Context, s string) {
	filename := "./file/" + u

}
