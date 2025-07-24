package itemtypemodel

import "example.com/m/internal/common"

type ItemType struct {
	common.MySqlModel
	Name string `json:"name" gorm:"column:name"`
}

func (ItemType) TableName() string {
	return "item_type"
}

func (s *ItemType) Mark() {
	s.GenUID(common.DBTypeResourceType)
}
