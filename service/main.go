package service

import "library_api/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		Auth:    implAuthService(repo),
		Book:    implBookService(repo),
		Lending: implLendingService(repo),
	}
}
