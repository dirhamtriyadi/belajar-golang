package usecase

import (
	"belajar-golang/internal/entity"
	"belajar-golang/internal/repository"
)

// UserUsecase is a contract
type UserUsecase interface {
	RegisterUser(user *entity.User) (*entity.UserResponse, error)
	FindAllUser() ([]*entity.UserResponse, error)
	FindUserByID(userID int) (*entity.UserResponse, error)
	FindUserByUsername(username string) (*entity.UserResponse, error)
	UpdateUser(userID int, user *entity.User) (*entity.UserResponse, error)
	DeleteUser(userID int) error
}

// userUsecase is a struct to represent user usecase
type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository}
}

func (u *userUsecase) RegisterUser(user *entity.User) (*entity.UserResponse, error) {
	return u.userRepository.CreateUser(user)
}

func (u *userUsecase) FindAllUser() ([]*entity.UserResponse, error) {
	return u.userRepository.FindAllUser()
}

func (u *userUsecase) FindUserByID(userID int) (*entity.UserResponse, error) {
	return u.userRepository.FindUserByID(userID)
}

func (u *userUsecase) FindUserByUsername(username string) (*entity.UserResponse, error) {
	return u.userRepository.FindUserByUsername(username)
}

func (u *userUsecase) UpdateUser(userID int, user *entity.User) (*entity.UserResponse, error) {
	return u.userRepository.UpdateUser(userID, user)
}

func (u *userUsecase) DeleteUser(userID int) error {
	return u.userRepository.DeleteUser(userID)
}
