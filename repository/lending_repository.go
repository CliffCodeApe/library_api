package repository

import (
	"library_api/contract"
	"library_api/entity"

	"gorm.io/gorm"
)

type lendingRepo struct {
	db *gorm.DB
}

func implLendingRepository(db *gorm.DB) contract.LendingRepository {
	return &lendingRepo{
		db: db,
	}
}

func (l *lendingRepo) GetLendingByID(id uint64) (*entity.Lending, error) {
	var lending entity.Lending
	if err := l.db.First(&lending, id).Error; err != nil {
		return nil, err
	}
	return &lending, nil
}

func (l *lendingRepo) GetAllLendings() ([]entity.Lending, error) {
	var lendings []entity.Lending
	if err := l.db.Find(&lendings).Error; err != nil {
		return nil, err
	}
	return lendings, nil
}

func (l *lendingRepo) MakeLending(lending *entity.Lending) error {
	if err := l.db.Create(lending).Error; err != nil {
		return err
	}
	return nil
}

func (l *lendingRepo) ChangeLendingStatus(id uint64, lending *entity.Lending) error {
	if err := l.db.Model(&entity.Lending{}).Where("id = ?", id).Updates(lending).Error; err != nil {
		return err
	}
	return nil
}

func (l *lendingRepo) GetLendingsByStatus(status string) ([]entity.Lending, error) {
	var lendings []entity.Lending
	if err := l.db.Where("status = ?", status).Find(&lendings).Error; err != nil {
		return nil, err
	}
	return lendings, nil
}

func (l *lendingRepo) SearchLendings(keyword string) ([]entity.Lending, error) {
	var lendings []entity.Lending
	likePattern := "%" + keyword + "%"
	if err := l.db.Where(
		"status ILIKE ?",
		likePattern,
	).Find(&lendings).Error; err != nil {
		return nil, err
	}
	return lendings, nil
}
