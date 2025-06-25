package entity

import "time"

type Book struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Title       string    `gorm:"column:title;type:varchar(255);not null"`
	Author      string    `gorm:"column:author;type:varchar(255);not null"`
	Year        int       `gorm:"column:year;type:int;not null"`
	Genre       string    `gorm:"column:genre;type:varchar(100);not null"`
	Stock       int       `gorm:"column:stock;type:int;not null;default:0"`
	Description string    `gorm:"column:description;type:text;not null"`
	Publisher   string    `gorm:"column:publisher;type:varchar(255);not null"`
	ISBN        string    `gorm:"column:isbn;type:varchar(100);not null"`
	Language    string    `gorm:"column:language;type:varchar(100);not null"`
	Pages       int       `gorm:"column:pages;type:int;not null"`
	Thumbnail   string    `gorm:"column:thumbnail;type:text;not null"`
	FilePath    string    `gorm:"column:file_path;type:text;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (e *Book) TableName() string {
	return "books"
}
