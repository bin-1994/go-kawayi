package service

import (
	"errors"
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/admin/bangumi/repository"
	"go-biu/library/database/orm"
)

type StyleService struct {

}

func (this *StyleService) Find() (list []model.Style) {
	list = new(repository.StyleRepository).Find(orm.Db)
	return list
}

func (this *StyleService) FindById(id string) (data *model.Style) {
	data = new(repository.StyleRepository).FindById(orm.Db, id)
	return data
}

func (this *StyleService) Create(data model.Style) error {
	if err := new(repository.StyleRepository).Create(orm.Db, &data); err != nil {
		return errors.New("创建失败")
	}
	return nil
}

func (this *StyleService) Update(data model.Style) error {
	if err := new(repository.StyleRepository).Update(orm.Db, &data); err != nil {
		return errors.New("修改失败")
	}
	return nil
}