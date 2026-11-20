package service

var NewBangumiService = newBangumiService()

type BangumiService interface {
}

func newBangumiService() bangumiService {
	return bangumiService{}
}

type bangumiService struct {
}

func (this *bangumiService) Get() {

}

func (this *bangumiService) Create() {

}

func (this *bangumiService) Edit() {

}