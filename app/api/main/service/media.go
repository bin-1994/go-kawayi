package service

type MediaService struct {
}

//func (this *MediaService) GetDetail(id int) model.MediaDetail {
//	db := orm.Db
//	media := new(repository.MediaRepository).Get(db, id)
//
//	return model.MediaDetail{
//		MediaId:    media.MediaId,
//		SeasonId:   media.SeasonId,
//		Type:       media.Type,
//		Version:    media.Version,
//		Title:      media.Title,
//		Keyword:    media.Keyword,
//		Story:      media.Story,
//		Bd:         media.Bd,
//		IsJump:     media.IsJump,
//		IsLimit:    media.IsLimit,
//		Level:      media.Level,
//		Cover:      media.Cover,
//		VideoCover: media.VideoCover,
//		Url:        media.Url,
//		Download:   media.Download,
//		Staff:      media.Staff,
//		Cast:       media.Cast,
//		Publish:    model.MediaPublish{EpIndex: media.EpIndex, EpTotal: media.EpTotal, IsFinish: media.IsFinish, PubTime: media.PubTime, PubString: media.PubString, OnAir: media.OnAir, Tips: media.Tips},
//		Count:      model.MediaCount{Follows: media.Follows, SeriesFollow: media.SeriesFollow, Views: media.Views},
//		Rating:     model.MediaRating{Score: media.Score, ScoreTotal: media.ScoreTotal},
//		Seasons:    new(repository.MediaRepository).GetSeries(db, media.SeasonId),
//		Related:    new(repository.MediaRepository).GetRelated(db, media.Tags, media.SeasonId),
//		Styles:     new(repository.StyleRepository).Get(db, media.MediaId),
//		Reviews:    new(repository.ReviewRepository).GetList(db, media.MediaId),
//		Episodes:   new(repository.EpisodeRepository).Get(db, media.MediaId, media.IsLimit),
//	}
//}

//func (this *MediaService) GetMediaList(params *model.MediaIndexParams, offset, limit int) (list []model.MediaIndex, total int64) {
//	// 综合排序 最近更新 最新上线 追番人数 人气指数
//	orders := []string{"follows", "pub_time", "update_time", "pub_time", "series_follow", "views", ""}
//	sorts := []string{"DESC", "ASC"}
//
//	orderBy := fmt.Sprintf("%s %s, media_id DESC", orders[params.Order], sorts[params.Sort])
//	db := orm.Db.Where("is_limit <= 1 AND level < 18").Order(orderBy)
//	if params.Category != 0 {
//		db = db.Where("type = ?", params.Category)
//	}
//	if params.Version != 0 {
//		if params.Version == 3 {
//			db = db.Where("version NOT IN ?", []int64{1, 2})
//		} else {
//			db = db.Where("version = ?", params.Version)
//		}
//	}
//	if params.State != 0 {
//		db = db.Where("is_finish = ?", params.State)
//	}
//	if params.Area != 0 {
//		switch params.Area {
//		case 1:
//			db = db.Where("area IN (?)", []int64{201, 101201, 201301})
//		case 2:
//			// 中国地区
//			db = db.Where("area IN (?)", []int64{101, 102, 103, 101201, 101301})
//		case 3:
//			// 美国地区
//			db = db.Where("area IN (?)", []int64{301, 101301, 201301})
//		case 4:
//			// 其他地区
//			db = db.Where("area NOT IN (?)", []int64{101, 102, 103, 201, 301, 101201, 101301})
//		}
//	}
//	if params.Month != 0 {
//		switch params.Month {
//		case 1:
//			// 冬季 1月-3月
//			db = db.Where("FROM_UNIXTIME(pub_time, '%c') IN (1,2,3)")
//		case 4:
//			// 春季 4月-6月
//			db = db.Where("FROM_UNIXTIME(pub_time, '%c') IN (4,5,6)")
//		case 7:
//			// 夏季 5月-9月
//			db = db.Where("FROM_UNIXTIME(pub_time, '%c') IN (7,8,9)")
//		case 10:
//			// 秋季 10月-12月
//			db = db.Where("FROM_UNIXTIME(pub_time, '%c') IN (10,11,12)")
//		}
//	}
//	if params.Year != "0" {
//		switch params.Year {
//		case "更早":
//			db = db.Where("FROM_UNIXTIME(pub_time, '%Y') < 1980")
//		case "80年代":
//			db = db.Where("FROM_UNIXTIME(pub_time, '%Y') BETWEEN 1980 AND 1989")
//		case "90年代":
//			db = db.Where("FROM_UNIXTIME(pub_time, '%Y') BETWEEN 1990 AND 1999")
//		default:
//			matched, _ := regexp.MatchString(`[\d+]-[\d+]`, params.Year)
//			if matched {
//				arrYear := strings.Split(params.Year, "-")
//				db = db.Where("FROM_UNIXTIME(pub_time, '%Y') BETWEEN ? AND ?", arrYear[1], arrYear[0])
//			} else {
//				db = db.Where("FROM_UNIXTIME(pub_time, '%Y') = ?", params.Year)
//			}
//		}
//	}
//	if params.Style != 0 {
//		var ids []int64
//		orm.Db.Table("media_style").Where("style_id = ?", params.Style).Group("media_id").Pluck("media_id", &ids)
//		db = db.Where("media_id in (?)", ids)
//	}
//	field := "media_id,title,cover,subtitle,type,is_finish,bd,pub_time,pub_string,ep_index,ep_total,score,series_follow,follows,update_time,views"
//	if params.Order == 4 {
//		field = field + ",COUNT(season_id) AS series_count"
//		db = db.Group("season_id")
//	}
//	list, total = new(repository.MediaRepository).GetList(db, field, offset, limit)
//	return
//}

//func (this *MediaService) GetFilterList() model.FilterModel {
//	var category model.FilterList
//	category.Title = "分类"
//	category.Key = "category"
//	category.Items = append(category.Items, model.FilterItem{Value: "0", Name: "全部"})
//	category.Items = append(category.Items, model.FilterItem{Value: "1", Name: "动画"})
//	category.Items = append(category.Items, model.FilterItem{Value: "2", Name: "电视剧"})
//
//	var version model.FilterList
//	version.Title = "版本"
//	version.Key = "version"
//	version.Items = append(version.Items, model.FilterItem{Value: "0", Name: "版本"})
//	version.Items = append(version.Items, model.FilterItem{Value: "1", Name: "正片"})
//	version.Items = append(version.Items, model.FilterItem{Value: "2", Name: "电影"})
//	version.Items = append(version.Items, model.FilterItem{Value: "3", Name: "其他"})
//
//	var area model.FilterList
//	area.Title = "地区"
//	area.Key = "area"
//	area.Items = append(area.Items, model.FilterItem{Value: "0", Name: "全部"})
//	area.Items = append(area.Items, model.FilterItem{Value: "1", Name: "日本"})
//	area.Items = append(area.Items, model.FilterItem{Value: "2", Name: "中国"})
//	area.Items = append(area.Items, model.FilterItem{Value: "3", Name: "欧美"})
//	area.Items = append(area.Items, model.FilterItem{Value: "4", Name: "其他"})
//
//	var state model.FilterList
//	state.Title = "状态"
//	state.Key = "state"
//	state.Items = append(state.Items, model.FilterItem{Value: "0", Name: "全部"})
//	state.Items = append(state.Items, model.FilterItem{Value: "1", Name: "连载"})
//	state.Items = append(state.Items, model.FilterItem{Value: "2", Name: "完结"})
//
//	var month model.FilterList
//	month.Title = "季度"
//	month.Key = "month"
//	month.Items = append(month.Items, model.FilterItem{Value: "0", Name: "全部"})
//	month.Items = append(month.Items, model.FilterItem{Value: "1", Name: "1月"})
//	month.Items = append(month.Items, model.FilterItem{Value: "4", Name: "4月"})
//	month.Items = append(month.Items, model.FilterItem{Value: "7", Name: "7月"})
//	month.Items = append(month.Items, model.FilterItem{Value: "10", Name: "10月"})
//
//	var year model.FilterList
//	year.Title = "时间"
//	year.Key = "year"
//	year.Items = append(year.Items, model.FilterItem{Value: "0", Name: "全部"})
//
//	nowYear, nowMonth, _ := time.Now().Date()
//	if nowMonth >= 11 {
//		nowYear++
//	}
//	index := 0
//	for i := nowYear; i >= 2000; i-- {
//		if index < 2 {
//			year.Items = append(year.Items, model.FilterItem{Value: strconv.Itoa(i), Name: strconv.Itoa(i)})
//		} else {
//			if index >= 2 && i%5 == 0 {
//				str := fmt.Sprintf("%s-%s", strconv.Itoa(i+4), strconv.Itoa(i))
//				year.Items = append(year.Items, model.FilterItem{Value: str, Name: str})
//			}
//		}
//		if i%5 == 0 {
//			index++
//		}
//	}
//	year.Items = append(year.Items, model.FilterItem{Value: "90年代", Name: "90年代"})
//	year.Items = append(year.Items, model.FilterItem{Value: "80年代", Name: "80年代"})
//	year.Items = append(year.Items, model.FilterItem{Value: "更早", Name: "更早"})
//
//	var style model.FilterList
//	style.Title = "风格"
//	style.Key = "style"
//	style.Items = append(style.Items, model.FilterItem{Value: "0", Name: "全部"})
//	//styles := new(repository.StyleRepository).GetList(orm.Db)
//	//for _, item := range styles {
//	//	style.Items = append(style.Items, model.FilterItem{Value: strconv.Itoa(item.StyleId), Name: item.Title})
//	//}
//
//	var order model.FilterList
//	order.Title = "排序"
//	order.Items = append(order.Items, model.FilterItem{Value: "1", Name: "开播时间", Sort: "0,1"})
//	order.Items = append(order.Items, model.FilterItem{Value: "2", Name: "更新时间", Sort: "0,1"})
//	order.Items = append(order.Items, model.FilterItem{Value: "3", Name: "最高评分", Sort: "0"})
//	order.Items = append(order.Items, model.FilterItem{Value: "4", Name: "追番人数", Sort: "0,1"})
//	order.Items = append(order.Items, model.FilterItem{Value: "5", Name: "人气指数", Sort: "0,1"})
//	//order.Items = append(order.Items, model.FilterItem{Value: "5", Name: "神作指数", Sort: "0"})
//
//	var list model.FilterModel
//	list.Filter = append(list.Filter, category, version, area, state, month, year, style)
//	list.Order = order
//
//	return list
//}

//func (this *MediaService) GetUserState(uid int64, media_id int, ep_id int64) *model.UserState {
//	//state := &model.UserState{}
//	db := orm.Db
//
//	if media_id == 0 {
//		media_id = new(repository.EpisodeRepository).EpFindId(db, ep_id)
//	}
//
//	media := new(repository.MediaRepository).Get(db, media_id)
//	follow := new(repository.FollowRepository).Get(db, uid, media.SeasonId)
//	review := new(repository.ReviewRepository).Get(db, uid, media_id)
//	state := &model.UserState{}
//	if follow != nil {
//		state.Progress.MediaId = follow.MediaId
//		state.Progress.SeriesId = follow.SeriesId
//		state.Progress.EpId = follow.EpId
//		state.Progress.Progress = follow.Progress
//		state.Progress.FollowState = follow.FollowState
//	}
//	if review != nil {
//		state.Review.ReviewId = review.ReviewId
//		state.Review.Score = review.Score
//		state.Review.Content = review.Content
//		state.Review.IsCoin = review.IsCoin
//		state.Review.UpdateTime = review.UpdateTime
//	}
//	return state
//}

//func (this *MediaService) GetSearchList(keyword string, offset, limit int) ([]model.SearchMedia, int64) {
//
//	hour := time.Now().Hour()
//	db := orm.Db
//	if hour >= 20 || hour <= 6 {
//		db = db.Where("is_limit <= 1")
//	} else {
//		db = db.Where("is_limit = 0")
//	}
//
//	list, total := new(repository.MediaRepository).GetSearchList(db, keyword, true, offset, limit)
//	return list, total
//}

//func (this *MediaService) GetSearchSuggest(keyword string) []model.SearchMedia {
//	list, _ := new(repository.MediaRepository).GetSearchList(orm.Db, keyword, false, 0, 10)
//	return list
//}

//func (this *MediaService) AddHotword(keyword string) {
//	if keyword == "" {
//		return
//	}
//	conn := cache.Redis.Get()
//	defer conn.Close()
//	// 全部搜索词 先留着
//	//conn.Do("SADD", "search:keyword", keyword)
//
//	// 当前日期时间戳
//	year, month, day := time.Now().Date()
//	timestamp := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
//	redisKey := fmt.Sprintf("search:keyword:%v", timestamp)
//	conn.Do("ZINCRBY", redisKey, 1, keyword)
//	conn.Do("EXPIRE", redisKey, 86400*14)
//}

//func (this *MediaService) FetchHotword() (interface{}, error) {
//	conn := cache.Redis.Get()
//	defer conn.Close()
//
//	year, month, day := time.Now().Date()
//	timestamp := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
//	lastDay := timestamp - (2 * 86400)
//
//	conn.Do("ZUNIONSTORE", "search:hot", 3, fmt.Sprintf("search:keyword:%v", lastDay), fmt.Sprintf("search:keyword:%v", lastDay+86400), fmt.Sprintf("search:keyword:%v", lastDay+(2*86400)), "WEIGHTS", 1, 1, 1)
//
//	res, _ := redis.Values(conn.Do("ZREVRANGE", "search:hot", 0, 9, "WITHSCORES"))
//
//	if len(res) == 0 {
//		return nil, errors.New("数据不存在")
//	}
//
//	var list []interface{}
//	for key, value := range res {
//		if key%2 == 0 {
//			item := make(map[string]interface{}, 0)
//			item["keyword"] = string(value.([]byte))
//			list = append(list, item)
//		}
//	}
//
//	return list, nil
//}

//func (this *MediaService) FetchTimeLine() (list []model.TimeLine, err error) {
//	var timeline []model.TimeLine
//
//	if err := new(repository.MediaRepository).GetTimeLineList(orm.Db, &timeline); err != nil {
//		return nil, errors.New("查询失败")
//	}
//
//	return timeline, nil
//}
