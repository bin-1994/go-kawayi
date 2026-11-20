package service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-biu/library/config"
	"go-biu/library/database/cache"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	proxy *url.URL
	timeout = 30 * time.Second
)

func GetYunpanUrl (vid string) string {
	proxy, _ = url.Parse(config.YamlConf.Base.Proxy)

	var token string
	yunpanToken, err := cache.Get("yunpan_token")
	json.Unmarshal(yunpanToken, &token)
	if err != nil {
		step_login()
	}

	if err == nil {
		yunpanKey := fmt.Sprintf("yunpan_url_%s_%s", token, vid)
		yunpanUrl, err := cache.Get(yunpanKey)
		var url string
		json.Unmarshal(yunpanUrl, &url)
		if err != nil {
			url = step_yunpan_url(vid)
			expire := 600
			if url != "" {
				expire = 3600
			}
			cache.Set(yunpanKey, url, expire)
		}
		return url
	}
	return ""
}

func step_login () {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: 				http.ProxyURL(proxy),
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
		Jar: jar,
		Timeout: timeout,
	}
	token := step_token()
	if token == "" {
		return
	}
	post := fmt.Sprintf("src=pcw_eyun&from=pcw_eyun&charset=UTF-8&requestScema=https&o=sso&m=login&lm=0&captFlag=1&rtype=data&validatelm=1&isKeepAlive=&captchaApp=i360&userName=15220554696&smDeviceId=&type=normal&account=15220554696&password=%s&captcha=&token=%s&proxy=https://yunpan.360.cn/psp_jump.html&&callback=QiUserJsonp3084175&func=QiUserJsonp3084175", "a70945417c73329acc42fe23abea9571", token)
	loginReq, _ := http.NewRequest("POST", "https://login.360.cn/", strings.NewReader(post))

	loginReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	loginReq.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	loginReq.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	loginReq.Header.Set("Cache-Control", "max-age=0")
	loginReq.Header.Set("Connection", "keep-alive")
	loginReq.Header.Set("Referer", "https://yunpan.360.cn/")

	loginRes, err := client.Do(loginReq)

	loginBody, _ := ioutil.ReadAll(loginRes.Body)
	decodeurl, err := url.QueryUnescape(string(loginBody))
	if err != nil {
		panic(err)
	}

	var data = struct {
		LoginTime string `json:"loginTime"`
	}{}
	reg := regexp.MustCompile(`({.+}})&`)
	json.Unmarshal([]byte(reg.FindStringSubmatch(decodeurl)[1]), &data)

	url := fmt.Sprintf("https://yunpan.360.cn/user/login?u=https://yunpan.360.cn/file/index&st=%s&sid=&keepalive=0", data.LoginTime)
	tokenReq, _ := http.NewRequest("GET", url, nil)

	tokenReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	tokenReq.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	tokenReq.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	tokenReq.Header.Set("Cache-Control", "max-age=0")
	tokenReq.Header.Set("Connection", "keep-alive")
	tokenReq.Header.Set("Referer", "https://yunpan.360.cn/")

	tokenRes, err := client.Do(tokenReq)

	var yunpanCookie string
	for _, c := range jar.Cookies(tokenReq.URL) {
		yunpanCookie += fmt.Sprintf("%s=%s; ", c.Name, c.Value)
		// token存入redis 12小时
		if c.Name == "token" {
			cache.Set("yunpan_token", c.Value, 3600 * 12)
		}
	}
	// cookie存入redis 12小时
	cache.Set("yunpan_cookie", yunpanCookie, 3600 * 12)

	defer loginRes.Body.Close()
	defer tokenRes.Body.Close()
}

func step_token () string {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: 				http.ProxyURL(proxy),
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
		Timeout: timeout,
	}
	url := fmt.Sprintf("https://login.360.cn?func=jQuery111106153758925958073_1541003084121&src=pcw_eyun&from=pcw_eyun&charset=UTF-8&requestScema=https&o=sso&m=getToken&userName=15220554696&_=%d", time.Now().Unix() * 1000)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://yunpan.360.cn/")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var data = struct {
		Error int `json:"error"`
		Errmsg string `json:"errmsg"`
		Token string `json:"token"`
	}{}
	// jQuery111109565141896488636_1553410851807({"errno":0,"errmsg":"","token":"6422a8c05d3d969e"})
	reg := regexp.MustCompile(`{(.+)}`)
	json.Unmarshal(reg.Find(body), &data)
	if data.Error == 0 {
		return data.Token
	}
	return ""
}

func step_yunpan_url (vid string) string {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: 				http.ProxyURL(proxy),
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
		Timeout: timeout,
	}
	url := fmt.Sprintf("https://yunpan.360.cn/videoPlayer/index?nid=%s&type=1&owner_qid=0", vid)
	req, _ := http.NewRequest("GET", url, nil)

	var ck, token string
	cookie, _ := cache.Get("yunpan_cookie")
	json.Unmarshal(cookie, &ck)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://yunpan.360.cn/")
	req.Header.Set("Cookie", ck)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	redisToken, _ := cache.Get("yunpan_token")
	json.Unmarshal(redisToken, &token)
	reg := regexp.MustCompile(`downloadUrl:\s\'(.+)&',`)
	video := reg.FindStringSubmatch(string(body))
	return fmt.Sprintf("%s&token=%s", video[1], token)
}
