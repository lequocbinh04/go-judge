package appCommon

import "time"

type SQLModel struct {
	Id        int64      `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 27112004)
	m.FakeId = &uid
}
