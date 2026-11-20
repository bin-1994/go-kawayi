package service

import (
	"errors"
	"fmt"
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/admin/bangumi/repository"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"strconv"
	"time"
)

type EpisodeService struct {
}

func (this *EpisodeService) Get(id string, offset, limit int) (model.MediaEpRela, int64) {
	db := orm.Db

	info := new(repository.MediaRepository).First(db, id)
	list, total := new(repository.EpisodeRepository).FindById(db, id, offset, limit)

	var rela = model.MediaEpRela{
		MediaId: info.MediaId,
		Title:   info.Title,
		EpList:  list,
	}

	return rela, total
}

func (this *EpisodeService) Create(data model.EpForm, is_update string) error {
	if err := new(repository.EpisodeRepository).Create(orm.Db, &data); err != nil {
		return errors.New("添加失败")
	}

	if is_update == "1" {
		var value = map[string]interface{}{
			"ep_index":    data.EpIndex,
			"ep_id":       data.EpId,
			"update_time": time.Now().Unix(),
			"tips":        "",
		}
		new(repository.MediaRepository).Update(orm.Db, data.MediaId, value)
	}

	utils.NewRedis().Del(fmt.Sprintf("bangumi:%s", data.MediaId))
	utils.NewRedis().Del("timeline")

	return nil
}

func (this *EpisodeService) Update(data model.EpForm) error {
	if err := new(repository.EpisodeRepository).Update(orm.Db, data); err != nil {
		return errors.New("编辑失败")
	}
	return nil
}

func (this *EpisodeService) Delete(id string) error {
	if err := new(repository.EpisodeRepository).Delete(orm.Db, id); err != nil {
		return errors.New("删除失败")
	}
	return nil
}

func (this *EpisodeService) CreateEp(media_id, source string, title, ep_index, jump_url, index, vurl []string) error {
	epAfter := make(map[string]interface{})
	db := orm.Db.Begin()
	for key, _ := range index {
		if row := new(repository.EpisodeRepository).IsExistEp(db, media_id, index[key]); row > 0 {
			continue
		}
		media_id, _ := strconv.Atoi(media_id)
		i, _ := strconv.Atoi(index[key])
		epItem := model.EpForm{
			MediaId: media_id,
			Source:  source,
			Title:   title[key],
			EpIndex: ep_index[key],
			Version: 1,
			JumpUrl: jump_url[key],
			Vurl:    vurl[key],
			Index:   i,
		}
		if err := new(repository.EpisodeRepository).Create(db, &epItem); err != nil {
			db.Rollback()
			return errors.New(fmt.Sprintf("操作失败, 第%d个", key+1))
		}
		epAfter["ep_id"] = epItem.EpId
		epAfter["ep_index"] = epItem.EpIndex
	}
	epAfter["update_time"] = time.Now().Unix()
	id, _ := strconv.Atoi(media_id)
	new(repository.MediaRepository).Update(db, id, epAfter)
	db.Commit()
	return nil
}
