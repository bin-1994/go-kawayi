package service

import (
	"fmt"
	"go-biu/app/admin/bangumi/repository"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"time"
)

//type MediaParams struct {
//	model.Media
//	Offset int
//	Limit int
//	Keyword string
//}
//
//func (m *MediaParams) GetAll () interface{} {
//	var media []model.Media
//	db := orm.Db
//	if m.Keyword != "" {
//		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", m.Keyword))
//	}
//	db.Limit(m.Limit).Offset(m.Offset).Order("id DESC").Find(&media)
//
//	//orm.Db.Where("title LIKE ?", "%海贼王%").Limit(m.Limit).Offset(m.Offset).Order("id DESC").Find(&media)
//	return media
//}
//
//func (m *MediaParams) Count () int64 {
//	var count int64
//	db := orm.Db
//	if m.Keyword != "" {
//		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", m.Keyword))
//	}
//	db.Model(&model.Media{}).Count(&count)
//	return count
//}

type MediaService struct {
}

func (this *MediaService) UpdateTime(id int) error {
	new(repository.MediaRepository).Update(orm.Db, id, map[string]interface{}{"update_time": time.Now().Unix()})
	return nil
}

func (this *MediaService) DeleteCaceh(id string) {
	utils.NewRedis().Del(fmt.Sprintf("bangumi:%s", id))
}
