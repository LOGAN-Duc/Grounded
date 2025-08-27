package boimesmodel

import (
	"example.com/m/internal/common"
	itemmodel "example.com/m/internal/module/item/model"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type Biome struct {
	common.MySqlModel
	Name        string `json:"name" gorm:"column:name;not null"`
	Url         string `json:"url" gorm:"column:url"`
	Description string `json:"description" gorm:"column:description"`

	Items     []*itemmodel.Item          `json:"items" gorm:"foreignKey:BiomeId"`
	Resources []*resourcemodel.Resources `json:"resources" gorm:"many2many:biome_resources;joinForeignKey:BiomeId;JoinReferences:ResourceId"`
}
type BiomeResource struct {
	BiomeId    *int `json:"biomeId" gorm:"column:biome_id"`
	ResourceId *int `json:"-" gorm:"column:resource_id;"`
}

func (Biome) TableName() string {
	return "biomes"
}
func (BiomeResource) TableName() string {
	return "biome_resources"
}
