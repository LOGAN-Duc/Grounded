package resourcemodel

import (
	"example.com/m/internal/common"
	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

type Resources struct {
	common.MySqlModel
	Name           string                          `json:"name" gorm:"column:name"`
	Code           string                          `json:"code" gorm:"column:code"`
	ResourceTypeId int                             `json:"-" gorm:"column:resource_type_id"`
	ResourceType   *resourcetypemodel.ResourceType `json:"resource_type" gorm:"foreignKey:ResourceTypeId"`

	Url string `json:"url" gorm:"column:url"`
}

func (Resources) TableName() string {
	return "resources"
}
