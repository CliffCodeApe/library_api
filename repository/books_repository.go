package repository

import (
	"library_api/contract"
	"library_api/entity"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func implBookRepository(db *gorm.DB) contract.BookRepository {
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func (b *bookRepo) GetBookByID(id uint64) (*entity.Book, error) {
	var book entity.Book
	if err := b.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *bookRepo) InsertBook(book *entity.Book) error {
	if err := b.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}
