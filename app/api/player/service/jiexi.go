package service

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func GetResolutionUrl(url string) (src string) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
		Timeout: 30 * time.Second,
	}
	// https://json.m3u8.tv:7788/json.php?url=
	// https://json.pangujiexi.com:7788/json.php?url=
	req, _ := http.NewRequest("GET", "https://json.m3u8.tv:7788/json.php?url=" + url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	type result struct {
		Url string `json:"url"`
		Type string `json:"type"`
		Player string `json:"player"`
		Code int `json:"code"`
		Success int `json:"success"`
		Msg string `json:"msg"`
		Title string `json:"title"`
	}
	var data result
	if err := json.Unmarshal(body, &data); err != nil {
		return ""
	}
	if data.Code != 200 {
		return ""
	}
	return data.Url
}
