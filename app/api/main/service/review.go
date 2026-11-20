package service

type ReviewService struct {
}

//func (this *ReviewService) Post(uid int64, params *model.ReviewForm) (*model.ReviewForm, error) {
//
//	if params.Score % 2 != 0 || (params.Score < 2 || params.Score > 10) {
//		return nil, errors.New("点评失败")
//	}
//
//	db := orm.Db
//	review := &model.ReviewForm{}
//	rd := repository.ReviewRepository{}
//
//	info := rd.Get(db, uid, params.MediaId)
//
//	if info.ReviewId != 0 {
//		// 存在id修改数据
//		review.Content = params.Content
//		review.Score = params.Score
//		if err := rd.Update(db, uid, review); err != nil {
//			return nil, errors.New("修改失败")
//		}
//	} else {
//		// 不存在id创建数据
//		review.Score = params.Score
//		review.Content = params.Content
//		review.MediaId = params.MediaId
//		review.Uid = uid
//		if err := rd.Create(db, review); err != nil {
//			return nil, errors.New("点评失败")
//		}
//	}
//
//	total, score := rd.Count(db, params.MediaId)
//	// 大于10人 计算评分
//	if total > 10 {
//		avg := float64(score * 10 / total) / float64(10)
//		new(repository.MediaRepository).CountAvg(db, params.MediaId, avg, total)
//	}
//
//	return review, nil
//}

//func (this *ReviewService) Delete(uid int64, media_id int, review_id int64) error {
//	db := orm.Db
//	if err := new(repository.ReviewRepository).Delete(db, uid, media_id, review_id); err != nil {
//		return errors.New("删除失败")
//	}
//
//	total, score := new(repository.ReviewRepository).Count(db, media_id)
//	// 大于10人 计算评分
//	if total > 10 {
//		avg := float64(score * 10 / total) / float64(10)
//		new(repository.MediaRepository).CountAvg(db, media_id, avg, total)
//	}
//
//	return  nil
//}
