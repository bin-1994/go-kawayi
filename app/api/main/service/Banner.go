package service

import (
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/api/main/repository"
)

type IBannerService interface {
	FindAll(int, int) []model.BannerModel
}

type BannerService struct {
	Repository repository.IBannerRepository
}

func NewBannerService(repository repository.IBannerRepository) IBannerService {
	return &BannerService{Repository: repository}
}

func (service *BannerService) FindAll(offset, limit int) []model.BannerModel {
	list := service.Repository.FindAll(offset, limit)
	return list
}
