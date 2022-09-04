package service

import "gilab.com/pragmaticreviews/golang-gin-poc/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos)

	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
