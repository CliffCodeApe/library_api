package service

import "library_api/contract"

type AuthService struct {
	userRepository contract.UserRepository
}

func implAuthService(repo *contract.Repository) contract.AuthService {
	return &AuthService{
		userRepository: repo.User,
	}
}
