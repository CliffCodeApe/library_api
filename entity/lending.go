package entity

import "time"

type Lending struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	UserID    uint64    `gorm:"column:user_id;not null;"`
	BookID    uint64    `gorm:"column:book_id;not null;"`
	Status    string    `gorm:"column:status;not null;default:'not_returned'"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}
