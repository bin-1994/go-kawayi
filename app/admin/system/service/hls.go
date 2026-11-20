package service

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"go-biu/app/admin/system/model"
	"go-biu/app/admin/system/repository"
	"go-biu/utils"
)

type HlsService interface {
	Upload([]*multipart.FileHeader, int, *gin.Context) ([]model.HlsModel, error)
	List(int) []model.HlsModel
	Get(int64) (*model.HlsModel, string, error)
	Update(*model.HlsModel, string) error
}

type service struct {
	repository repository.HlsRepository
}

func NewHlsService(repository repository.HlsRepository) HlsService {
	return &service{
		repository: repository,
	}
}

const rootPath = "./uploads/hls"

func (s *service) Upload(files []*multipart.FileHeader, media_id int, ctx *gin.Context) ([]model.HlsModel, error) {
	hashId := utils.EncodeHashId([]int{media_id}, 4)
	if err := utils.IsNotExistMkDir(rootPath + "/" + hashId); err != nil {
		return nil, err
	}
	var data []model.HlsModel
	for _, file := range files {
		ext := path.Ext(file.Filename)
		fileData := model.HlsModel{
			Md5:  utils.Md5(file.Filename),
			Sha1: utils.Sha1(file.Filename),
			Size: file.Size,
		}

		if isExist := s.repository.IsExist(&fileData); !isExist {
			// 不存在 则创建

			fileData.Name = file.Filename
			name := utils.Md5(fmt.Sprintf("%s:%s", file.Filename, time.Now().UnixNano()))
			fileData.Path = "/" + hashId + "/" + name + ext
			fileData.MediaId = media_id
			fileData.Dir = hashId
			fileData.CreateTime = time.Now().Unix()

			if err := ctx.SaveUploadedFile(file, rootPath+fileData.Path); err != nil {
				return nil, errors.New("上传失败")
			}

			s.repository.Create(&fileData)
		}
		fmt.Println(fileData)
		data = append(data, fileData)
	}
	return data, nil
}

func (s *service) List(media_id int) []model.HlsModel {
	list := s.repository.List(&model.HlsModel{MediaId: media_id})

	return list
}

func (s *service) Get(id int64) (*model.HlsModel, string, error) {

	data := s.repository.Get(id)
	if isExist := utils.IsExist(rootPath + data.Path); !isExist {
		return nil, "", errors.New("文件不存在")
	}
	// 读取文件内容
	file, err := os.Open(rootPath + data.Path)
	if err != nil {
		return nil, "", errors.New("打开文件失败")
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", errors.New("读取文件失败")
	}
	return &data, string(content), nil
}

func (s *service) Update(data *model.HlsModel, content string) error {
	_, err := s.repository.Update(data)
	if err != nil {
		return errors.New("更新信息失败")
	}

	path := rootPath + data.Path
	if isExist := utils.IsExist(path); !isExist {
		return errors.New("文件不存在")
	}

	file, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return errors.New("打开文件失败")
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	if err != nil {
		return errors.New("写入文件失败")
	}

	return nil
}
