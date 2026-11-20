package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/system/model"
	"go-biu/app/admin/system/service"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	MAX_SIZE = 2 * 1024 * 1024
	ROOT     = "./uploads"
)

var exts = [4]string{".jpg", ".jpeg", ".png", ".gif"}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	dir := c.PostForm("dir")
	year := c.PostForm("year")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "上传失败",
		})
		return
	}
	filename := file.Filename
	//filesize := file.Size
	//
	md5 := utils.Md5(filename)
	sha1 := utils.Sha1(filename)
	size := file.Size
	ext := path.Ext(filename)

	if !checkImageExt(ext) {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "图片格式不正确",
		})
		return
	}

	if size > MAX_SIZE {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "图片大小不能超过2MB",
		})
		return
	}

	var fileModel model.UploadModel
	if result := orm.Db.Where("`md5` = ? AND `sha1` = ? AND `size` = ?", md5, sha1, size).First(&fileModel); result.RowsAffected == 0 {
		// 查询不到 上传
		utils.IsNotExistMkDir(ROOT + "/" + dir + "/" + year)
		name := utils.Md5(fmt.Sprintf("%s:%s", filename, time.Now().UnixNano()))
		path := fmt.Sprintf("/%s/%s/%s%s", dir, year, name, ext)
		fmt.Println(path)
		if err := c.SaveUploadedFile(file, ROOT+path); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "上传失败",
			})
			return
		}
		var data = model.UploadModel{
			Dir:        dir,
			Md5:        md5,
			Sha1:       sha1,
			Name:       filename,
			Size:       size,
			Path:       path,
			CreateTime: time.Now().Unix(),
			Ext:        ext,
		}
		orm.Db.Create(&data)
		fileModel = data
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "上传成功",
		"result":  fileModel,
	})
}

func checkImageExt(ext string) bool {
	for _, allowExt := range exts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func FetchAllUpload(c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	var (
		list  []model.UploadModel
		total int64
	)
	db := orm.Db

	keyword := c.Query("keyword")
	if keyword != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	dir := c.Query("dir")
	db = db.Where("dir = ?", dir)
	db.Table("upload").Count(&total).Offset(offset).Limit(length).Order("id DESC").Find(&list)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":   0,
		"result": list,
		"total":  total,
	})
}

func DeleteUpload(c *gin.Context) {
	path := c.PostForm("path")

	tx := orm.Db.Begin()
	if err := tx.Exec("DELETE FROM `upload` WHERE `path` = ?", path).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "删除失败",
		})
		return
	}

	err := os.Remove(ROOT + path)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "删除失败",
		})
		return
	}

	//if err := os.Remove(service.WEBP_PATH + strings.Replace(path, "jpg", "webp", -1)); err != nil {
	//	tx.Rollback()
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    1,
	//		"message": "删除失败",
	//	})
	//	return
	//}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}

func CreateWebp(ctx *gin.Context) {
	path := ctx.PostForm("path")
	if err := new(service.WebpService).ToWebp(path); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建成功",
	})
}

func MoveImg(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	start := (page - 1) * size

	var list = []struct {
		Id      int    `json:"id"`
		MediaId int    `json:"media_id"`
		Path    string `json:"path"`
		PubTime int64  `json:"pub_time"`
	}{}
	orm.Db.Table("upload").
		Select("upload.path AS path,media.pub_time AS pub_time, upload.id AS id, media.media_id AS media_id").
		Where("upload.dir = ?", "video").
		Joins("INNER JOIN media ON media.b_video_cover = upload.path").
		Order("id ASC").
		Offset(start).
		Limit(size).
		Find(&list)

	for _, v := range list {
		year := time.Unix(v.PubTime, 0).Year()
		if year < 2000 {
			year = 1900
		}
		dir := fmt.Sprintf("/video/%d", year)
		//path := strings.Replace(v.Path, "bangumi", dir, -1)
		if !utils.IsExist("./uploads" + v.Path) {
			fmt.Println("文件不存在")
			continue
		}
		utils.IsNotExistMkDir(ROOT + dir)
		newPath := strings.Replace(v.Path, "/video", dir, -1)
		err := os.Rename(ROOT+v.Path, ROOT+newPath)
		if err != nil {
			log.Fatal(err)
		}

		orm.Db.Table("upload").Where("id = ?", v.Id).Update("path", newPath)
		orm.Db.Table("media").Where("media_id = ?", v.MediaId).Update("b_video_cover", newPath)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page": page,
		"size": size,
		"list": list,
	})
}
