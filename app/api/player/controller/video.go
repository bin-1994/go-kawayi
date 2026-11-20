package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go-biu/app/api/player/service"
	"go-biu/library/database/cache"
)

func Player(c *gin.Context) {
	value := c.Param("value")
	arr := strings.Split(value, ".")
	enbyte, err := base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		panic(err)
	}

	var data = struct {
		Expire  int    `json:"expire"`
		EpId    int    `json:"ep_id"`
		Url     string `json:"url"`
		JumpUrl string `json:"jump_url"`
		Source  string `json:"source"`
	}{}

	json.Unmarshal(enbyte, &data)

	if int(time.Now().Unix()) > data.Expire {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "已过期",
		})
		return
	}
	data.Url, _ = url.QueryUnescape(data.Url)
	var url string
	if data.Source == "jx" && data.JumpUrl != "" {
		url = getJiexiUrl(data.JumpUrl)
	} else if matched, _ := regexp.MatchString("hoopchina", data.Url); matched {
		//cacheKey := `hupu:` + data.Url
		//cacheUrl, _ := cache.Get(cacheKey)
		//json.Unmarshal(cacheUrl, &url)
		//if url == "" {
		//	url = service.GetHupuUrl(data.Url)
		//	cache.Set(cacheKey, url, 3600 * 2)
		//}
		url = service.GetHupuUrl(data.Url)
	} else {
		url = getCdnUrl(data.Url)
	}

	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请求链接失败",
		})
		return
	}
	c.Header("Content-Security-Policy", "upgrade-insecure-requests")
	c.Redirect(302, url)
}

func getJiexiUrl(url string) (src string) {
	src = service.GetResolutionUrl(url)

	return src
}

func getCdnUrl(url string) (src string) {
	if matched, _ := regexp.MatchString("1098_", url); matched {
		u, err := cache.Get(url)
		json.Unmarshal(u, &src)
		if err != nil {
			src = service.GetTimeUrl(url)
			cache.Set(url, src, 1800)
		}
		return src
	}
	return url
}

func WeiYun(c *gin.Context) {
	id := c.Query("id")
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "12"))

	if id == "" {
		c.JSON(200, map[string]interface{}{
			"message": "数据错误",
		})
		return
	}

	var arr []string

	for i := offset; i <= limit; i++ {
		str := fmt.Sprintf("https://www.sbdm.net/index.php/vod/play/id/%s/sid/1/nid/%d.html", id, i)
		arr = append(arr, str)
	}

	url := service.GetAliyun(arr)

	c.JSON(200, map[string]interface{}{
		"result": url,
	})
}
