package service

import "github.com/hotbrandon/go-gin-tutorial/entity"

type IVideoService interface {
	Save(entity.Video) entity.Video
	GetAll() []entity.Video
}

type VideoService struct {
	videos []entity.Video
}

func (svc *VideoService) Save(vid entity.Video) entity.Video {
	svc.videos = append(svc.videos, vid)
	return vid
}

func (svc *VideoService) GetAll() []entity.Video {
	return svc.videos
}

func NewVideoService() VideoService {
	return VideoService{}
}
