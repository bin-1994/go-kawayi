package service

import (
	"errors"
	"go-biu/app/admin/model"
	"go-biu/app/admin/repository"
)

type IBangumiService interface {
	Create(*model.BangumiModel, []string) error
	Update(*model.BangumiModel, []string) error
	//Delete()
	//FindById(int) model.BangumiResponse
	//FindAll(model.BangumiQuery, int, int) ([]model.BangumiModel, int64)
	//UpdateEpisode(bangumi model.BangumiModel) error
}

type bangumiService struct {
	Repository repository.IBangumiRepository
}

func NewBangumiService(bangumiRepository repository.IBangumiRepository) IBangumiService {
	return &bangumiService{Repository: bangumiRepository}
}

func (service *bangumiService) Create(bangumi *model.BangumiModel, styles []string) error {
	if err := service.Repository.Create(bangumi); err != nil {
		return errors.New("创建失败")
	}

	if bangumi.Id > 0 {

	}

	return nil
}

func (service *bangumiService) Update(bangumi *model.BangumiModel, styles []string) error {
	if err := service.Repository.Update(bangumi); err != nil {
		return errors.New("编辑失败")
	}
	return nil
}
