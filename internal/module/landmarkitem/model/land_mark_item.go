package landmarkitemmodel

import (
	"example.com/m/internal/common"
	itemmodel "example.com/m/internal/module/item/model"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type LandMarkItem struct {
	common.MySqlModel
	LandmarkID *int                     `json:"landmark_id" gorm:"column:landmark_id"` // phải đúng chữ hoa I
	ItemID     *int                     `json:"item_id" gorm:"column:item_id"`
	Item       *itemmodel.Item          `json:"item" gorm:"foreignKey:ItemID"`
	ResourceID *int                     `json:"resource_id" gorm:"column:resource_id"`
	Resource   *resourcemodel.Resources `json:"resource" gorm:"foreignKey:ResourceID"`
}

func (LandMarkItem) TableName() string {
	return "landmark_items"
}
