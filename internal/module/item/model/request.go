package itemmodel

type CreateItemRequest struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ItemTypeId int    `json:"itemTypeId"`
}
type UpdateItemRequest struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ItemTypeId *int   `json:"itemTypeId"`
}

type ListItemRequest struct {
	Search string `json:"search" form:"search"`
}
