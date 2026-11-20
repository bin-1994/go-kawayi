package controller

//func FetchMyBangumi(c *gin.Context) {
//	uid, _ := c.Get("uid")
//	page := c.DefaultQuery("page", "1")
//	pagesize := c.DefaultQuery("pagesize", "15")
//	state, exists := c.GetQuery("state")
//
//	start, end := utils.GetPage(page, pagesize)
//
//	var (
//		bangumi []model.Media
//		total   int64
//	)
//	db := orm.Db.Table("media_follow")
//
//	if exists && state != "0" {
//		db = db.Where("media_follow.follow_state = ?", state)
//	}
//
//	db = db.Where("uid = ?", int64(uid.(float64)))
//
//	db.Select("media.media_id AS media_id, media.season_id AS season_id,media.title AS title, media.cover AS cover, media.video_cover AS video_cover, media.story AS story, media.is_finish AS is_finish, media.ep_index AS ep_index, media.ep_total AS ep_total, media.update_time AS update_time, media_follow.last_time AS last_time, media_follow.ep_id AS ep_id, media_follow.ep_index AS watch_ep, media_follow.progress AS progress,media_follow.follow_state AS follow_state").
//		Joins("INNER JOIN media ON media_follow.media_id = media.media_id").
//		Offset(start).Limit(end).
//		Order("GREATEST(media.update_time, media_follow.create_time) DESC").Find(&bangumi)
//
//	db.Count(&total)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "success",
//		"result": map[string]interface{}{
//			"list":  bangumi,
//			"total": total,
//		},
//	})
//}
