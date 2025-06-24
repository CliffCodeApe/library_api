package service

import "library_api/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		Auth: implAuthService(repo),
		// Code here
		// Example:
		// Example: implExampleService(repo),
	}
}
