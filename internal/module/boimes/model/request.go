package boimesmodel

type CreateBiomeRequest struct {
	Name        string `json:"name" `
	Url         string `json:"url"`
	Description string `json:"description"`
	ItemIds     []int  `json:"itemIds"`
	ResourceIds []int  `json:"resourceIds"`
}

type FilterBiomeRequest struct {
	Search string `json:"search" `
	Item   string `json:"item"`
}
type UpdateBiomeRequest struct {
	Name        string `json:"name" `
	Url         string `json:"url"`
	Description string `json:"description"`
	ItemIds     []int  `json:"itemIds"`
	ResourceIds []int  `json:"resourceIds"`
}
