package service

import (
	"regexp"
)

func GetAliyun(arr []string) string {
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	//		DisableCompression: true,
	//	},
	//	Timeout: 30 * time.Second,
	//}
	//
	//req, _ := http.NewRequest("GET", "https://www.sbdm.net/index.php/vod/play/id/1241/sid/1/nid/2.html", nil)
	//
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cache-Control", "max-age=0")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Referer", "https://www.sbdm.net/")
	//
	//res, err := client.Do(req)
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()

	//body, _ := ioutil.ReadAll(res.Body)
	str := `var player_aaaa={\"flag\":\"play\",\"encrypt\":0,\"trysee\":0,\"points\":0,\"link\":\"\\/index.php\\/vod\\/play\\/id\\/1241\\/sid\\/1\\/nid\\/1.html\",\"link_next\":\"\\/index.php\\/vod\\/play\\/id\\/1241\\/sid\\/1\\/nid\\/3.html\",\"link_pre\":\"\\/index.php\\/vod\\/play\\/id\\/1241\\/sid\\/1\\/nid\\/1.html\",\"url\":\"onopen-x9uRg8WNZj1|63b8bec6c5ff23c7f2fd4e86a245896b971219cb\",\"url_next\":\"onopen-x9uRg8WNZj1|63b8bec9085f8745799944a8b421516375fdf249\",\"from\":\"dplayer\",\"server\":\"no\",\"note\":\"\",\"id\":\"1241\",\"sid\":1,\"nid\":2}`
	// \"url\":\"onopen-x9uRg8WNZj1|63b8bec6c5ff23c7f2fd4e86a245896b971219cb\",\"url_next\":
	//reg := regexp.MustCompile(`var url = \'(.+)\';`)
	//reg := regexp.MustCompile(`"url":"(onopen.+)",`)
	//url := reg.FindStringSubmatch(string(body))
	reg := regexp.MustCompile(`"url":(onopen.+)`)
	url := reg.FindAllString(str, -1)

	//fmt.Println(string(body))

	return url[1]
}
