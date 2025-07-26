package itemresourcemodel

import "example.com/m/internal/common"

type ItemResource struct {
	common.MySqlModel
	ItemId     int `json:"item_id" gorm:"column:item_id;primaryKey"`
	ResourceId int `json:"resource_id" gorm:"column:resource_id;primaryKey"`
	Quantity   int `json:"quantity" gorm:"column:quantity"`
}

func (ItemResource) TableName() string {
	return "item_resources"
}
