package service

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func GetQzoneUrl(vid string) string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
		Timeout: 30 * time.Second,
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("http://tsdmw.net/yunbo/1006.php?url=%s", vid), nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", fmt.Sprintf("http://www.tldm.net/js/player/j3.html?p=p&a=%s&b=b&w=660&h=550", vid))

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	reg := regexp.MustCompile(`var url = \'(.+)\';`)
	url := reg.FindStringSubmatch(string(body))
	if url[1] != "" {
		return url[1]
	}
	return ""
	//return fmt.Sprintf("%s", url[1])
}
