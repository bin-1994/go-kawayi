package service

import (
	"errors"
	"go-biu/app/api/main/repository"
	"go-biu/app/api/main/model"
	"go-biu/library/database/orm"
	"unicode/utf8"
)

type DanmakuService struct {

}

func (this *DanmakuService) GetList(ep_id string, offset, limit int) ([]model.DanmakuModel, int64) {
	list, total := new(repository.DanmakuRepository).GetList(orm.Db, ep_id, offset, limit)
	return list, total
}

func (this *DanmakuService) Create(uid int64, dm model.DanmakuModel) error {

	if utf8.RuneCountInString(dm.Text) > 30 {
		return errors.New("弹幕太长了")
	}

	pramas := &model.DanmakuModel{
		Uid: uid,
		EpId: dm.EpId,
		MediaId: dm.MediaId,
		Mode: dm.Mode,
		Color: dm.Color,
		Text: dm.Text,
		Size: dm.Size,
		Stime: dm.Stime,
		Ip: dm.Ip,
	}

	if err := new(repository.DanmakuRepository).Create(orm.Db, pramas); err != nil {
		return errors.New("发送弹幕失败")
	}

	return nil
}
