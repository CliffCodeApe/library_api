package entity

type User struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Username  string `gorm:"column:username;type:varchar(255);not null;unique"`
	Email     string `gorm:"column:email;type:varchar(255);not null;unique"`
	Password  string `gorm:"column:password;type:varchar(255);not null"`
	Role      string `gorm:"column:role;type:varchar(50);not null;default:'member'"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;not null;default:now()"`
}

func (e *User) TableName() string {
	return "users"
}
