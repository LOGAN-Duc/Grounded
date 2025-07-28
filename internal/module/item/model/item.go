package itemmodel

import (
	"example.com/m/internal/common"
	itemtypemodel "example.com/m/internal/module/item_type/model"
)

type Item struct {
	common.MySqlModel
	Name       string                  `json:"name" gorm:"column:name;not null"`
	Code       string                  `json:"code,omitempty" gorm:"column:code"`
	ItemTypeId int                     `json:"-" gorm:"column:item_type_id"`
	ItemType   *itemtypemodel.ItemType `json:"item_Type"`
}

func (Item) TableName() string {
	return "items"
}
