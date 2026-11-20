package controller

// var db *ipdb.City

// func init () {
// 	db, _ = ipdb.NewCity("./ipip/ipipfree.ipdb")
// }

// func BanIp (c *gin.Context) {

// 	ip := c.Query("ip")
// 	//ip := c.ClientIP()
// 	fmt.Println(ip)
// 	// 匹配到ipv4 通过
// 	r, _ := regexp.Compile(`(\d+\.\d+\.\d+\.\d+)`)
// 	if match := r.MatchString(ip); !match {
// 		return
// 	}

// 	area, err := db.Find(ip, "CN")
// 	if err != nil {

// 	}
// 	ban := false
// 	if match, _ := regexp.MatchString("中国", area[0]); !match {
// 		ban = false
// 	}

// 	if match, _ := regexp.MatchString("香港|台湾|澳门", area[1]); match {
// 		ban = true
// 	}

// 	if match, _ := regexp.MatchString("北京|上海|广州|深圳|成都|杭州|重庆|武汉|西安|苏州|天津|南京|长沙|郑州|东莞|青岛|沈阳|宁波|昆明|合肥", area[2]); match {
// 		ban = true
// 	}

// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"code": 0,
// 		"result": ban,
// 		"area": area[0]+ area[1] + area[2],
// 	})
// }
