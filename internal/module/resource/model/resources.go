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
	ResourceType   *resourcetypemodel.ResourceType `json:"resourceType"`
}

func (Resources) TableName() string {
	return "resources"
}

// func (r *Resources) Mark() {
// 	r.MySqlModel.GenUID(common.DBTypeResource)
// }
