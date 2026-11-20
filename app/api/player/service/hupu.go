package service

import (
	"fmt"
	"go-biu/utils"
	"regexp"
	"strings"
	"time"
)

func GetHupuUrl(url string) string {
	reg := regexp.MustCompile(`(\/bbs-editor-web\/editor\/.*\.mp4)`)
	prefix := reg.FindAllString(url, -1)

	expired := time.Now().Unix() + (3600 * 24)
	auth_key := utils.Md5(fmt.Sprintf(`%s-%v-0-0-UqY3yu`, prefix[0], expired))

	domain := strings.Split(url, "?")

	return fmt.Sprintf(`%s?auth_key=%v-0-0-%s`, domain[0], expired, auth_key)
}
