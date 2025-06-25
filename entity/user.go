package entity

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Username  string    `gorm:"column:username;type:varchar(255);not null;unique"`
	Email     string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Password  string    `gorm:"column:password;type:varchar(255);not null"`
	Role      string    `gorm:"column:role;type:varchar(50);not null;default:'member'"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (e *User) TableName() string {
	return "users"
}
