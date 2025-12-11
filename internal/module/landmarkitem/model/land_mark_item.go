package landmarkitemmodel

import "example.com/m/internal/common"

type LandMarkItem struct {
	common.MySqlModel
	LandmarkID *int `json:"landmark_id" gorm:"column:landmark_id"` // phải đúng chữ hoa I
	ItemID     *int `json:"item_id" gorm:"column:item_id"`
	ResourceID *int `json:"resource_id" gorm:"column:resource_id"`
}

func (LandMarkItem) TableName() string {
	return "landmark_items"
}
