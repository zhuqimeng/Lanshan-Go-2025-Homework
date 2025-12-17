package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Name     string `gorm:"size:64;not null;comment:代办项目"`
	Priority string `gorm:"size:32;default:低;comment:优先级"`
	Status   string `gorm:"size:32;default:未开始;comment:任务状态"`
	Version  uint   `gorm:"default:0;comment:版本号机制"`
}

type QueryRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	Token    string `json:"token"`
	Version  string `json:"version"`
}
