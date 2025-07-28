package itemresourcemodel

type CreateItemRrsourceRequest struct {
	ResourceId int `json:"resource_id" gorm:"column:resource_id;primaryKey"`
	Quantity   int `json:"quantity" gorm:"column:quantity"`
}
type ListItemRrsourceRequest struct {
	Search string `json:"search"`
}
type UpdateItemRrsourceRequest struct {
	ResourceId int `json:"resource_id" gorm:"column:resource_id;primaryKey"`
	Quantity   int `json:"quantity" gorm:"column:quantity"`
}
