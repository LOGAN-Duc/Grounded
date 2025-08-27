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
	ItemType   *itemtypemodel.ItemType `json:"item_type" gorm:"foreignKey:ItemTypeId"`

	Url     string `json:"url" gorm:"column:url"`
	BiomeId *int   `json:"biomeId" gorm:"column:biome_id"`
}

func (Item) TableName() string {
	return "items"
}
