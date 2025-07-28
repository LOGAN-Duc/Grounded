package resourcemodel

import "example.com/m/internal/common"

type Resources struct {
	common.MySqlModel
	Name           string `json:"name" gorm:"column:name"`
	Code           string `json:"code" gorm:"column:code"`
	ResourceTypeId int    `json:"resourceTypeId" gorm:"column:resource_type_id"`
	// ResourceType
}

func (Resources) TableName() string {
	return "resources"
}

// func (r *Resources) Mark() {
// 	r.MySqlModel.GenUID(common.DBTypeResource)
// }
