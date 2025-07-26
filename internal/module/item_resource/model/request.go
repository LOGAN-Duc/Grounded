package itemresourcemodel

type CreateItemRrsourceRequest struct {
	ItemId     int    `json:"item_id" gorm:"column:item_id;primaryKey"`
	ResourceId int    `json:"resource_id" gorm:"column:resource_id;primaryKey"`
	Quantity   int    `json:"quantity" gorm:"column:quantity"`
	Status     string `json:"status" gorm:"column:status;default:1"`
}
