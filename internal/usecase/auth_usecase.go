package usecase

import (
	"errors"
	"gopos/internal/domain"
	"gopos/pkg/utils"
	"time"
)

type AuthUsecase interface {
	Register(user *domain.RegisterRequest) (*domain.User, error)
	Login(email, password string) (*domain.LoginResponse, error)
	LoginInfo(userID uint) (*domain.User, error)
}

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(userRepo domain.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

// Register new user
func (u *authUsecase) Register(user *domain.RegisterRequest) (*domain.User, error) {
	existingUser, _ := u.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	existingUser, _ = u.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return nil, errors.New("username already registered")
	}

	if !utils.IsValidUsername(user.Username) {
		return nil, errors.New("username must not be an email")
	}

	if !utils.IsEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	savedUser, err := u.userRepo.Save(newUser)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

// Login and return JWT token
func (u *authUsecase) Login(username, password string) (*domain.LoginResponse, error) {
	user, err := u.userRepo.FindByEmailOrUsername(username)
	if err != nil || user == nil {
		return nil, errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	data := map[string]interface{}{
		"user_id": user.ID,
	}
	// Generate token
	token, expireAt, err := utils.GenerateToken(data)
	if err != nil {
		return nil, err
	}

	res := &domain.LoginResponse{
		User: *user,
		Token: domain.TokenResponse{
			Token:        token,
			ExpireAt:     expireAt.Format(time.RFC3339),
			TokenType:    "Bearer",
			IssuedAt:     time.Now().Format(time.RFC3339),
			RefreshToken: "",
		},
	}

	return res, nil
}

// LoginInfo returns user data by ID
func (u *authUsecase) LoginInfo(userID uint) (*domain.User, error) {
	return u.userRepo.FindByID(userID)
}
