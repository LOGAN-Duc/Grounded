package common

import "strings"

type Paging struct {
	Limit int64 `json:"limit" form:"limit"`
	Page  int64 `json:"page" form:"page"`
	Total int64 `json:"total" form:"total"`
	//	Support cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor" form:"next_cursor"`
}

func (p Paging) Fulfill() Paging {
	newPg := p
	if p.Limit == 0 {
		newPg.Limit = 10
	}
	if p.Page == 0 {
		newPg.Page = 1
	}
	return newPg
}

func (p Paging) FulfillMySql() Paging {
	newPg := p
	if p.Page <= 0 {
		newPg.Page = 1
	}
	if p.Limit <= 0 {
		newPg.Limit = 10
	}

	newPg.FakeCursor = strings.TrimSpace(p.FakeCursor)
	return newPg
}

func (p Paging) GetList() Paging {
	newPg := p
	if p.Limit == 0 {
		newPg.Limit = 20
	}
	if p.Page == 0 {
		newPg.Page = 1
	}
	return newPg
}

func (p *Paging) SetLimit(limit int64) *Paging {
	p.Limit = limit
	return p
}

func (p *Paging) SetPage(page int64) *Paging {
	p.Page = page
	return p
}
