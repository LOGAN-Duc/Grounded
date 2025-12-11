package landmarkitemmodel

import "example.com/m/internal/common"

type LandMarkItem struct {
	common.MySqlModel
	LandmarkId int `json:"landmark_id" gorm:"column:landmark_id"`
	ItemId     int `json:"item_id" gorm:"column:item_id"`
	ResourceId int `json:"resource_id" gorm:"column:resource_id"`
}

func (LandMarkItem) TableName() string {
	return "landmark_items"
}
