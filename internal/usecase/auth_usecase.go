package usecase

import (
	"gopos/internal/domain"
	appError "gopos/pkg/errors"
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
		return nil, appError.ErrEmailExist
	}

	existingUser, _ = u.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return nil, appError.ErrUsernameExist
	}

	if !utils.IsValidUsername(user.Username) {
		return nil, appError.ErrUsernameFormat
	}

	if !utils.IsEmail(user.Email) {
		return nil, appError.ErrEmailFormat
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, appError.ErrHashPassword
	}

	newUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	savedUser, err := u.userRepo.Save(newUser)
	if err != nil {
		return nil, appError.Get(appError.ErrUserCreate, err)
	}

	return savedUser, nil
}

// Login and return JWT token
func (u *authUsecase) Login(username, password string) (*domain.LoginResponse, error) {
	user, err := u.userRepo.FindByEmailOrUsername(username)
	if err != nil || user == nil {
		return nil, appError.ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, appError.ErrInvalidCredentials
	}

	data := map[string]interface{}{
		"user_id": user.ID,
	}
	// Generate token
	token, expireAt, err := utils.GenerateToken(data)
	if err != nil {
		return nil, appError.Get(appError.ErrGenerateToken, err)
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
