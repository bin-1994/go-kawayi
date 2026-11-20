package controller

//func FetchMyReview (c *gin.Context) {
//	uid := c.GetInt("uid")
//	page := c.DefaultQuery("page", "1")
//	pagesize := c.DefaultQuery("pagesize", "15")
//	start, end := utils.GetPage(page, pagesize)
//
//	var (
//		list []model.Review
//		total int64
//	)
//
//	orm.Db.Preload("Media").Where("uid = ?", uid).Order("update_time DESC, review_id DESC").Offset(start).Limit(end).Find(&list)
//
//	orm.Db.Table("review").Where("uid = ?", uid).Count(&total)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": map[string]interface{}{
//			"list": list,
//			"total": total,
//		},
//	})
//}
