package controller

//func FetchDanmaku(c *gin.Context) {
//	ep_id := c.Query("ep_id")
//	var p utils.Paging
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "参数错误",
//		})
//		return
//	}
//	start, end := p.SetPage(p.Page, p.Pagesize)
//	list, total := new(service.DanmakuService).GetList(ep_id, start, end)
//
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": map[string]interface{}{
//			"list": list,
//			"page": gin.H{
//				"num": p.Page,
//				"size": p.Pagesize,
//				"total": total,
//			},
//		},
//	})
//}

//func AddDanmaku(c *gin.Context)  {
//	uid, _ := c.Get("uid")
//	var dm model.DanmakuModel
//	dm.Ip = c.ClientIP()
//	if err := c.ShouldBind(&dm); err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	if err := new(service.DanmakuService).Create(int64(uid.(float64)), dm); err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "发送弹幕成功",
//	})
//}
