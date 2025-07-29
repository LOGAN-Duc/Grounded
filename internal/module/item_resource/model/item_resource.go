package itemresourcemodel

import (
	"example.com/m/internal/common"
	itemmodel "example.com/m/internal/module/item/model"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type ItemResource struct {
	common.MySqlModel
	ItemId     int                      `json:"-" gorm:"column:item_id;primaryKey"`
	Item       *itemmodel.Item          `json:"item" `
	ResourceId int                      `json:"-" gorm:"column:resource_id;primaryKey"`
	Resource   *resourcemodel.Resources `json:"resource"`
	Quantity   int                      `json:"quantity" gorm:"column:quantity"`
}

func (ItemResource) TableName() string {
	return "item_resources"
}
