package service

import "github.com/MohamedAbdelrazeq/gin-video-server/entity"

type VideoService interface {
	Save(entity.Video) (entity.Video, error)
	FindAll() ([]entity.Video, error)
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (service *videoService) Save(video entity.Video) (entity.Video, error) {
	service.videos = append(service.videos, video)
	return video, nil
}

func (service *videoService) FindAll() ([]entity.Video, error) {
	return service.videos, nil
}
