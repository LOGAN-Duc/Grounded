package itemmodel

type CreateItemRequest struct {
	Name       string `json:"name" form:"name"`
	Code       string `json:"code" form:"code"`
	ItemTypeId int    `json:"itemTypeId" form:"itemTypeId"`
	UrlImage   string `json:"urlImage" form:"urlImage"`
}

type UpdateItemRequest struct {
	Name       string `json:"name" form:"name"`
	Code       string `json:"code" form:"code"`
	ItemTypeId *int   `json:"itemTypeId" form:"itemTypeId"`
	UrlImage   string `json:"urlImage" form:"urlImage"`
}

type ListItemRequest struct {
	Search string `json:"search" form:"search"`
}
