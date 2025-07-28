package resourcetypemodel

import "example.com/m/internal/common"

type ResourceType struct {
	common.MySqlModel
	Name string `json:"name" gorm:"column:name"`
}

func (ResourceType) TableName() string {
	return "resources_type"
}

// func (s *ResourceType) Mark() {
// 	s.GenUID(common.DBTypeResourceType)
// }
