package itemmodel

type CreateItemRequest struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ItemTypeId int    `json:"itemTypeId"`
	BoimeId    *int   `json:"boimeId"`
}
type UpdateItemRequest struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ItemTypeId *int   `json:"itemTypeId"`
	BoimeId    *int   `json:"boimeId"`
}

type ListItemRequest struct {
	Search string `json:"search" form:"search"`
}
