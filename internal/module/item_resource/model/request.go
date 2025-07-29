package itemresourcemodel

type CreateItemRrsourceRequest struct {
	ResourceId int `json:"resource_id" gorm:"column:resource_id;primaryKey"`
	Quantity   int `json:"quantity" gorm:"column:quantity"`
}
type ListItemRrsourceRequest struct {
	Search string `json:"search" form:"search"`
	ItemId *int   `json:"itemId" form:"itemId"`
}
type UpdateItemRrsourceRequest struct {
	ResourceId int  `json:"resource_id" gorm:"column:resource_id;primaryKey"`
	Quantity   *int `json:"quantity" gorm:"column:quantity"`
}
