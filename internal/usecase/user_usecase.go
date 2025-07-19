package usecase

import (
	"errors"
	"gopos/internal/domain"
	"gopos/pkg/utils"
)

type UserUsecase interface {
	List() ([]domain.User, error)
	Detail(userId uint) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(userId uint) error
}

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) List() ([]domain.User, error) {
	return u.userRepo.List()
}

func (u *userUsecase) Detail(userId uint) (*domain.User, error) {
	foundUser, err := u.userRepo.FindByID(userId)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u *userUsecase) Create(user *domain.User) (*domain.User, error) {
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

func (u *userUsecase) Update(user *domain.User) (*domain.User, error) {
	// Cek apakah user dengan ID tersebut ada
	_, err := u.userRepo.FindByID(user.ID)
	if err != nil {
		return nil, err
	}

	// Cek email jika diubah
	existingUserByEmail, _ := u.userRepo.FindByEmail(user.Email)
	if existingUserByEmail != nil && existingUserByEmail.ID != user.ID {
		return nil, errors.New("email already registered by another user")
	}

	// Cek username jika diubah
	existingUserByUsername, _ := u.userRepo.FindByUsername(user.Username)
	if existingUserByUsername != nil && existingUserByUsername.ID != user.ID {
		return nil, errors.New("username already registered by another user")
	}

	if !utils.IsValidUsername(user.Username) {
		return nil, errors.New("username must not be an email")
	}

	if !utils.IsEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}

	// Hash password jika diisi
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	// Update data
	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *userUsecase) Delete(id uint) error {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	return u.userRepo.Delete(user)
}
