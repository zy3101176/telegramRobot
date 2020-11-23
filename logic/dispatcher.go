package logic

import "telegram_robot/logic/activity"

type DispatcherService struct {
	center ConfigCenter
	activityMap map[string]*activity.Activity
	userStatusMap map[int64]int64
}

func NewDispatcherService(center ConfigCenter) *DispatcherService {
	return &DispatcherService{
		center: center,
		activityMap:   make(map[string]*activity.Activity),
		userStatusMap: make(map[int64]int64),
	}
}

func (d *DispatcherService) Init() {

}