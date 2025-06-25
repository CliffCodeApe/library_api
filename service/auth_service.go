package service

import (
	"context"
	"fmt"
	"library_api/contract"
	"library_api/dto"
	"library_api/entity"
	"library_api/pkg/bcrypt"
	"library_api/pkg/errs"
	"library_api/pkg/helpers"
	"net/http"
	"regexp"
	"strings"
	"time"

	token2 "library_api/pkg/token"
)

type AuthService struct {
	userRepository contract.UserRepository
}

func implAuthService(repo *contract.Repository) contract.AuthService {
	return &AuthService{
		userRepository: repo.User,
	}
}

func (a *AuthService) Register(ctx context.Context, payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, validPayload
	}
	var missingFields []string
	if payload.Username == "" {
		missingFields = append(missingFields, "Username")
	}

	if payload.Password == "" {
		missingFields = append(missingFields, "Password")
	}

	if payload.Email == "" {
		missingFields = append(missingFields, "Email")
	} else if !isValidEmail(payload.Email) {
		return nil, errs.NewBadRequest("Format email salah!")
	}

	if len(missingFields) > 0 {
		return nil, errs.NewBadRequest(fmt.Sprintf("Field berikut harus diisi: %s", strings.Join(missingFields, ", ")))
	}

	hashPassword, err := bcrypt.Generate(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("bcrypt.Generate fail: %w", err)

	}

	exists, err := a.userRepository.IsEmailExists(payload.Email)
	if err != nil {
		return nil, fmt.Errorf("userRepo.IsEmailExists fail: %w", err)
	}
	if exists {
		return nil, errs.NewBadRequest("Email sudah terdaftar")
	}

	user := entity.User{
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	err = a.userRepository.InsertUser(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("userRepo.InsertUser fail: %w", err)
	}

	response := &dto.RegisterResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil register",
		Data: dto.AuthData{
			Username: payload.Username,
			Email:    payload.Email,
		},
	}

	return response, nil
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func (a *AuthService) Login(ctx context.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	var response *dto.LoginResponse
	// var loginErr error

	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, fmt.Errorf("helpers.ValidateStruct fail: %w", validPayload)
	}

	var missingFields []string
	if payload.Email == "" {
		missingFields = append(missingFields, "Email")
	}

	if payload.Password == "" {
		missingFields = append(missingFields, "Password")
	}

	if len(missingFields) > 0 {
		return nil, errs.NewBadRequest(fmt.Sprintf("Field berikut harus diisi: %s", strings.Join(missingFields, ", ")))
	}

	user, err := a.userRepository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, errs.ErrLoginFailed
	}

	verify := bcrypt.Verify(user.Password, payload.Password)

	if !verify {
		return nil, errs.ErrLoginFailed
	}

	accessToken, err := token2.GenerateAccessToken(&token2.UserAuthToken{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	})

	if err != nil {
		return nil, fmt.Errorf("token2.GenerateAccessToken fail: %w", err)
	}

	refreshToken, err := token2.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("token2.GenerateRefreshToken fail: %w", err)
	}

	response = &dto.LoginResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil Login",
		Data: dto.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}
