package usecase

import (
	"errors"
	"gopos/internal/domain"
	"gopos/pkg/utils"
	"time"
)

type AuthUsecase interface {
	Register(user *domain.RegisterRequest) error
	Login(email, password string) (*domain.LoginResponse, error)
	GetProfile(userID uint) (*domain.User, error)
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
func (u *authUsecase) Register(user *domain.RegisterRequest) error {
	existingUser, _ := u.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	return u.userRepo.Save(newUser)
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

	// Generate token
	token, expireAt, err := utils.GenerateToken(user.ID)
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

// GetProfile returns user data by ID
func (u *authUsecase) GetProfile(userID uint) (*domain.User, error) {
	return u.userRepo.FindByID(userID)
}
