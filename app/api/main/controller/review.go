package controller

//func FetchReview (c *gin.Context) {
//	media_id := c.Query("media_id")
//	page := c.DefaultQuery("page", "1")
//	pagesize := c.DefaultQuery("pagesize", "20")
//
//	start, end := utils.GetPage(page, pagesize)
//
//	var review []model.Review
//	result := orm.Db.Preload("Author").
//		Where("media_id = ? AND content <> ''", media_id).
//		Order("update_time DESC, review_id DESC").
//		Offset(start).Limit(end).Find(&review)
//
//	if result.Error != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": -404,
//			"message": "地球上找不到该内容哦 (╥﹏╥)",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": review,
//	})
//}

//func PostReview (c *gin.Context) {
//	uid, _ := c.Get("uid")
//	var form model.ReviewForm
//	if err := c.ShouldBind(&form); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "参数错误",
//		})
//		return
//	}
//
//	review, err := new(service.ReviewService).Post(int64(uid.(float64)), &form)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "点评成功",
//		"result": review,
//	})
//}

//func DeleteReview (c *gin.Context) {
//	uid, _ := c.Get("uid")
//
//	var p model.ReviewForm
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "参数错误",
//		})
//		return
//	}
//
//	if err := new(service.ReviewService).Delete(int64(uid.(float64)), p.MediaId, p.ReviewId); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "删除点评成功",
//	})
//}
