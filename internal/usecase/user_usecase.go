package usecase

import (
	"gopos/internal/domain"
	appErr "gopos/pkg/errors"
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
		return nil, appErr.Get(appErr.ErrUserDetail, err)
	}
	return foundUser, nil
}

func (u *userUsecase) Create(user *domain.User) (*domain.User, error) {
	existingUser, _ := u.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, appErr.Get(appErr.ErrEmailExist, nil)
	}

	existingUser, _ = u.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return nil, appErr.Get(appErr.ErrUsernameExist, nil)
	}

	if !utils.IsValidUsername(user.Username) {
		return nil, appErr.Get(appErr.ErrUsernameFormat, nil)
	}

	if !utils.IsEmail(user.Email) {
		return nil, appErr.Get(appErr.ErrEmailFormat, nil)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, appErr.Get(appErr.ErrHashPassword, err)
	}

	newUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	savedUser, err := u.userRepo.Save(newUser)
	if err != nil {
		return nil, appErr.Get(appErr.ErrUserCreate, err)
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
		return nil, appErr.Get(appErr.ErrEmailExist, nil)
	}

	// Cek username jika diubah
	existingUserByUsername, _ := u.userRepo.FindByUsername(user.Username)
	if existingUserByUsername != nil && existingUserByUsername.ID != user.ID {
		return nil, appErr.Get(appErr.ErrUsernameExist, nil)
	}

	if !utils.IsValidUsername(user.Username) {
		return nil, appErr.Get(appErr.ErrUsernameFormat, nil)
	}

	if !utils.IsEmail(user.Email) {
		return nil, appErr.Get(appErr.ErrEmailFormat, nil)
	}

	// Hash password jika diisi
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, appErr.Get(appErr.ErrHashPassword, err)
		}
		user.Password = hashedPassword
	}

	// Update data
	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		return nil, appErr.Get(appErr.ErrUserUpdate, err)
	}

	return updatedUser, nil
}

func (u *userUsecase) Delete(id uint) error {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return appErr.Get(appErr.ErrUserDelete, err)
	}
	return u.userRepo.Delete(user)
}
