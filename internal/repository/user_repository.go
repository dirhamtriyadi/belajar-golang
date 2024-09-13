package repository

import (
	"belajar-golang/internal/entity"

	"gorm.io/gorm"
)

// UserRepository is a contract
type UserRepository interface {
	CreateUser(user entity.User) (entity.UserResponse, error)
	FindAllUser() ([]entity.User, error)
	FindUserByID(userID int) (entity.User, error)
	FindUserByUsername(username string) (entity.User, error)
	UpdateUser(userID int, user entity.User) (entity.UserResponse, error)
	DeleteUser(userID int) error
}

// userRepository is a struct to represent user repository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.User) (entity.UserResponse, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return entity.UserResponse{}, err
	}

	return entity.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *userRepository) FindAllUser() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindUserByID(userID int) (entity.User, error) {
	var user entity.User

	if err := r.db.First(&user, userID).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByUsername(username string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(userID int, user entity.User) (entity.UserResponse, error) {
	if err := r.db.Model(&entity.User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		return entity.UserResponse{}, err
	}

	user, err := r.FindUserByID(userID)
	if err != nil {
		return entity.UserResponse{}, err
	}

	return entity.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *userRepository) DeleteUser(userID int) error {
	if err := r.db.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}

	return nil
}
