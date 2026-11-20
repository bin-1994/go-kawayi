package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func GetTimeUrl(reqUrl string) string {
	//proxy, _ := url.Parse("http://60.208.44.228:80")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
			//Proxy: http.ProxyURL(proxy),
			//Dial: (&net.Dialer{
			//	Timeout: 2 * time.Second,
			//	Deadline: time.Now().Add(3 * time.Second),
			//	KeepAlive: 2 * time.Second,
			//}).Dial,
		},
		Timeout: 180 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, _ := http.NewRequest("GET", reqUrl, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "quan.qq.com")
	//req.Header.Set("Referer", "")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	reg, _ := regexp.Compile(`(\S*)1098`)
	str := reg.ReplaceAllString(res.Header.Get("Location"), "")
	return "https://apd-vliveachy.apdcdn.tc.qq.com/vmtt.tc.qq.com/1098" + str
}
