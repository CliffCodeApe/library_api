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

func (b *bookRepo) GetBooksByGenre(genre string) ([]entity.Book, error) {
	var books []entity.Book
	if err := b.db.Where("genre = ?", genre).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (b *bookRepo) InsertBook(book *entity.Book) error {
	if err := b.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookRepo) ChangeStock(id uint64, delta int) error {
	if err := b.db.Model(&entity.Book{}).Where("id = ?", id).UpdateColumn("stock", gorm.Expr("stock + ?", delta)).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookRepo) SearchBooks(keyword string) ([]entity.Book, error) {
	var books []entity.Book
	likePattern := "%" + keyword + "%"
	if err := b.db.Where(
		"title ILIKE ? OR author ILIKE ? OR genre ILIKE ?",
		likePattern, likePattern, likePattern,
	).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (b *bookRepo) GetBookByURL(url string) (*entity.Book, error) {
	var book entity.Book
	if err := b.db.Where("url = ?", url).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
