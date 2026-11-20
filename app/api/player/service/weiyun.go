package service

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func GetWeiYunUrl () string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
		Timeout: 30 * time.Second,
	}
	vid := "1109aaf3-52f5-429f-afaa-8b25b8598e08"
	dir := "0d8d1d373e6609da7fcc649f1fd9d9ca"
	pdir := "0d8d1d37f0a739ae12b58dcd423dce4a"
	url := fmt.Sprintf("https://www.weiyun.com/video_preview?videoID=%s&dirKey=%s&pdirKey=%s", vid, dir, pdir)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.weiyun.com/disk/folder/" + dir)

	req.Header.Set("Cookie", "uin=o0924683533; skey=@enWcp0SRG;")

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	//return string(body);
	reg := regexp.MustCompile(`https_download_url\":\"(.+\.mp4)\?fname`)
	src := reg.FindStringSubmatch(string(body))
	if src[1] != "" {
		return fmt.Sprintf("%s", src[1])
	}
	return ""
}
