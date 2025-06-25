package service

import (
	"library_api/contract"
	"library_api/dto"
	"library_api/entity"
	"net/http"
)

type LendingService struct {
	lendingRepository contract.LendingRepository
	bookRepository    contract.BookRepository
}

func implLendingService(repo *contract.Repository) contract.LendingService {
	return &LendingService{
		lendingRepository: repo.Lending,
		bookRepository:    repo.Book,
	}
}

func (l *LendingService) GetAllLendings() (*dto.LendingResponse, error) {
	lendings, err := l.lendingRepository.GetAllLendings()
	if err != nil {
		return nil, err
	}

	data := []dto.LendingData{}
	for _, lending := range lendings {
		data = append(data, dto.LendingData{
			ID:     lending.ID,
			BookID: lending.BookID,
			UserID: lending.UserID,
			Status: lending.Status,
		})
	}

	responses := &dto.LendingResponse{
		StatusCode: http.StatusOK,
		Message:    "List Peminjaman berhasil diambil",
		Data:       data,
	}
	return responses, nil
}

func (l *LendingService) MakeLending(userID uint64, payload *dto.LendingRequest) (*dto.MakeLendingResponse, error) {

	book, err := l.bookRepository.GetBookByURL(payload.BookURL)
	if err != nil {
		return nil, err
	}

	lending := entity.Lending{
		BookID: book.ID,
		UserID: userID,
		Status: "borrowed",
	}

	if err := l.lendingRepository.MakeLending(&lending); err != nil {
		return nil, err
	}

	if err := l.bookRepository.ChangeStock(lending.BookID, -1); err != nil {
		return nil, err
	}

	response := &dto.MakeLendingResponse{
		StatusCode: http.StatusCreated,
		Message:    "Peminjaman berhasil dibuat",
		Data: dto.LendingData{
			ID:     lending.ID,
			UserID: lending.UserID,
			BookID: lending.BookID,
			Status: lending.Status,
		},
	}
	return response, nil
}

func (l *LendingService) ReturnBook(lendingID, userID uint64, fileBytes []byte) (*dto.MakeLendingResponse, error) {
	lending, err := l.lendingRepository.GetLendingByID(lendingID)
	if err != nil {
		return nil, err
	}

	lending.Status = "returned"
	if err := l.lendingRepository.ChangeLendingStatus(lendingID, lending); err != nil {
		return nil, err
	}

	if err := l.bookRepository.ChangeStock(lending.BookID, 1); err != nil {
		return nil, err
	}

	response := &dto.MakeLendingResponse{
		StatusCode: http.StatusOK,
		Message:    "Status peminjaman berhasil diubah",
		Data: dto.LendingData{
			ID:     lending.ID,
			UserID: lending.UserID,
			BookID: lending.BookID,
			Status: lending.Status,
		},
	}

	return response, nil
}

func (b *LendingService) GetLendingsByStatus(status string) (*dto.LendingResponse, error) {
	lendings, err := b.lendingRepository.GetLendingsByStatus(status)
	if err != nil {
		return nil, err
	}

	data := []dto.LendingData{}
	for _, lending := range lendings {
		data = append(data, dto.LendingData{
			ID:     lending.ID,
			BookID: lending.BookID,
			UserID: lending.UserID,
			Status: lending.Status,
		})
	}

	responses := &dto.LendingResponse{
		StatusCode: http.StatusOK,
		Message:    "List Peminjaman berhasil diambil",
		Data:       data,
	}
	return responses, nil
}

func (b *LendingService) SearchLendings(keyword string) (*dto.LendingResponse, error) {
	lendings, err := b.lendingRepository.SearchLendings(keyword)
	if err != nil {
		return nil, err
	}

	data := []dto.LendingData{}
	for _, lending := range lendings {
		data = append(data, dto.LendingData{
			ID:     lending.ID,
			BookID: lending.BookID,
			UserID: lending.UserID,
			Status: lending.Status,
		})
	}

	responses := &dto.LendingResponse{
		StatusCode: http.StatusOK,
		Message:    "List Peminjaman berhasil diambil",
		Data:       data,
	}
	return responses, nil
}
