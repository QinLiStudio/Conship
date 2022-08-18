package service

import (
	"io/ioutil"
	"os"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/pkg/error"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, p string, i schema.Item) {
	if err := configs.POSTGRESDB.Select("url", "secret").Create(&i); err.Error != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "数据库创建数据失败。", err.Error)
		return
	}

	file, err := os.Create("./file/" + i.Url)
	if err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "创建文件失败。", err)
		file.Close()
		return
	}

	if _, err := file.WriteString(p); err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "文件写入配置失败。", err)
		file.Close()
		return
	}

	file.Close()
	error.Response(c, error.Created, gin.H{"url": configs.CONFIG.Http.Host + "/" + i.Url, "secret": i.Secret}, "配置文件上传成功。")
}

func Delete(c *gin.Context, s string) {
	var i schema.Item

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Find(&i); err.Error != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "数据库请求数据失败。", err.Error)
		return
	}

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Delete(&i); err.Error != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "数据库删除数据失败。", err.Error)
		return
	}

	err := os.Remove("./file/" + i.Url)
	if err != nil {
		error.Response(c, error.NotFind, gin.H{}, "未找到配置文件。")
		return
	}

	error.Response(c, error.OK, gin.H{}, "配置文件删除成功。")
}

func Update(c *gin.Context, s string, p string) {
	var i schema.Item

	if err := configs.POSTGRESDB.Table("items").Where("secret = ?", s).Find(&i); err.Error != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "数据库请求数据失败。", err.Error)
		return
	}

	file, err := os.OpenFile("./file/"+i.Url, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		error.Response(c, error.NotFind, gin.H{}, "未找到配置文件。")
		file.Close()
		return
	}

	if _, err := file.WriteString(p); err != nil {
		error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "配置文件修改失败。", err)
		file.Close()
		return
	}

	file.Close()
	error.Response(c, error.OK, gin.H{"url": configs.CONFIG.Http.Host + "/" + i.Url, "secret": i.Secret}, "配置文件修改成功。")
}

func Search(c *gin.Context, u string) {
	a := u[1:]
	if len(a) == 7 {
		file, err := os.Open("./file/" + a)
		if err != nil {
			error.Response(c, error.NotFind, gin.H{}, "未找到配置文件。")
			file.Close()
			return
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "查询配置文件失败。", err)
			file.Close()
			return
		}

		file.Close()
		c.String(200, string(content))
		return
	}

	if len(a) == 8 {
		var i schema.Item
		if err := configs.POSTGRESDB.Table("items").Where("secret = ?", a).Find(&i); err.Error != nil {
			error.ErrResponse(c, error.ErrBadRequest, error.BadRequest, "数据库请求数据失败。", err.Error)
			return
		}

		file, err := os.Open("./file/" + i.Url)
		if err != nil {
			error.Response(c, error.NotFind, gin.H{}, "未找到配置文件。")
			file.Close()
			return
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			error.Response(c, error.NotFind, gin.H{}, "未找到配置文件。")
			file.Close()
			return
		}

		file.Close()
		c.String(200, string(content))
		return
	}
	error.Response(c, error.NotFind, gin.H{}, "未找到配置文件。")
}
