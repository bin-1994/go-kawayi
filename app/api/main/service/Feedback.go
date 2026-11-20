package service

import (
	"go-biu/app/api/main/model"
	"go-biu/utils"
	"strconv"
)

type IFeedbackService interface {
	Create(model.FeedbackModel) error
}

type FeedbackService struct {
}

func NewFeedbackService() IFeedbackService {
	return &FeedbackService{}
}

func (service *FeedbackService) Create(feedback model.FeedbackModel) error {
	id := strconv.Itoa(feedback.EpId)
	if err := utils.NewRedis().HSet("feedback", id, feedback); err != nil {
		return err
	}

	return nil
}
