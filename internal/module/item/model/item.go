package itemmodel

import "example.com/m/internal/common"

type Item struct {
	common.MySqlModel
	Name       string `json:"name" gorm:"column:name;not null"`
	Code       string `json:"code,omitempty" gorm:"column:code"`
	ItemTypeId int    `json:"item_type_id" gorm:"column:item_type_id"`
}

func (Item) TableName() string {
	return "items"
}
