package common

type MySqlModel struct {
	Id uint `json:"id" gorm:"column:id"`
	// FakeId *UID   `json:"id" gorm:"-"`
	Status string `json:"status,omitempty" gorm:"column:status;default:1"`
}

// func (m *MySqlModel) GenUID(dbType int) {
// 	uid := NewUID(uint32(m.Id), dbType, 1)
// 	m.FakeId = &uid
// }
