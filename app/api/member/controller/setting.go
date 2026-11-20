package controller

//func FetchSetting (c *gin.Context) {
//	uid := c.GetInt("uid")
//	var user model.User
//	if err := orm.Db.Select("nickname,mobile,birthday,gender,sign").Where("uid = ?", uid).First(&user).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "失败",
//		})
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": user,
//	})
//}

//func UpdateSetting (c *gin.Context) {
//	uid := c.GetInt("uid")
//	user := model.User{
//		Nickname: c.PostForm("nickname"),
//		//Mobile: c.PostForm("mobile"),
//		//Birthday: c.PostForm("birthday"),
//		//Gender: c.PostForm("gender"),
//		//Sign: c.PostForm("sign"),
//	}
//
//	if utf8.RuneCountInString(user.Nickname) < 2 || utf8.RuneCountInString(user.Nickname) > 8 {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "昵称长度2-8个字符",
//		})
//		return
//	}
//
//	if result := orm.Db.Where("nickname = ? AND uid <> ?", user.Nickname, uid).First(&user); result.RowsAffected > 0 {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "昵称已存在",
//		})
//		return
//	}
//
//	if err := orm.Db.Table("user").Where("uid = ?", uid).Updates(&user).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "更新信息失败",
//		})
//		return
//	}
//
//	// 删除缓存
//	cache.Delete(fmt.Sprintf("user_%s", uid))
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//	})
//}

//func UpdatePassword (c *gin.Context) {
//	uid := c.GetInt("uid")
//	password := c.PostForm("password")
//	oldpassword := c.PostForm("oldpassword")
//	repassword := c.PostForm("repassword")
//
//	if checkPass, _ := regexp.MatchString(`^[A-Za-z0-9]{6,16}$`, password); !checkPass {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "密码长度6-16字符",
//		})
//		return
//	}
//
//	if password != repassword {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "密码不一致",
//		})
//		return
//	}
//
//	var user model.User
//	if orm.Db.Table("user").Where("uid = ?", uid).First(&user); user.Password != utils.UserMd5(oldpassword) {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "旧密码不正确",
//		})
//		return
//	}
//
//	if err := orm.Db.Exec("UPDATE user SET password = ? WHERE uid = ?", utils.UserMd5(password), uid).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "修改密码失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "修改密码成功",
//	})
//}
