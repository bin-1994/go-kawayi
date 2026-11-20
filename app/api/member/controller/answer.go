package controller

import (
	"go-biu/app/api/member/model"
)

type cacheModel struct {
	Timestamp int            `json:"timestamp"`
	List      []model.Answer `json:"list"`
}

const countdown = 1801

//func FetchQuestion(c *gin.Context) {
//
//	cacheKey := fmt.Sprintf("answer_%s", c.ClientIP())
//
//	var question []model.Answer
//	// 倒计时30分钟
//	now := time.Now()
//
//	if exists := cache.Exists(cacheKey); !exists{
//		orm.Db.Order("RAND()").Limit(50).Find(&question)
//
//		cache.Set(cacheKey, cacheModel{Timestamp: int(now.Unix()), List: question}, countdown)
//	}
//	var list cacheModel
//	cacheList, _ := cache.Get(cacheKey)
//	json.Unmarshal(cacheList, &list)
//
//	result := make([]interface{}, 0)
//	for _, val := range list.List {
//		var option []string
//		option = append(option, val.Option1, val.Option2, val.Option3, val.Option4)
//
//		item := make(map[string]interface{})
//		item["question"] = val.Question
//		item["option"] = option
//		result = append(result, item)
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": result,
//		"timestamp": map[string]interface{}{
//			"end": countdown - (int(now.Unix()) - list.Timestamp),
//			"countdown": countdown,
//		},
//	})
//}
//
//func PostAnswer(c *gin.Context) {
//	answer := c.PostFormArray("answer[]")
//	cacheKey := fmt.Sprintf("answer_%s", c.ClientIP())
//
//	cacheList, err := cache.Get(cacheKey)
//	if err != nil || len(answer) < 50 {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "请认真答题！",
//		})
//		return
//	}
//
//	var list cacheModel
//	json.Unmarshal(cacheList, &list)
//
//	if (int(time.Now().Unix()) - list.Timestamp > countdown) {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "已经超时，答题结束",
//		})
//		return
//	}
//	score := 0
//	for key, val := range answer {
//		if val == list.List[key].Answer {
//			score = score + 2
//		}
//	}
//
//	cache.Delete(cacheKey)
//	if score >= 80 {
//		// 生成注册码
//		rand.Seed(time.Now().Unix())
//		number := rand.Intn(89999) + 10000
//		codeKey := fmt.Sprintf("kwy_reg_code:%s:%s", c.ClientIP(), strconv.Itoa(number))
//		cache.Set(codeKey, model.RegCode{Ip: c.ClientIP(), Code: strconv.Itoa(number)}, 300)
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 0,
//			"message": "恭喜你！通过考验",
//			"result": map[string]interface{}{
//				"code": number,
//				"score": score,
//			},
//		})
//	} else {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "很遗憾！请再接再厉",
//			"result": map[string]interface{}{
//				"score": score,
//			},
//		})
//	}
//}
