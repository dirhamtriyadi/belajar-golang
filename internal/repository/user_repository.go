package repository

import (
	"belajar-golang/internal/entity"

	"gorm.io/gorm"
)

// UserRepository is a contract
type UserRepository interface {
	CreateUser(user entity.User) (entity.UserResponse, error)
	FindAllUser() ([]entity.UserResponse, error)
	FindUserByID(userID int) (entity.UserResponse, error)
	FindUserByUsername(username string) (entity.UserResponse, error)
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

func (r *userRepository) FindAllUser() ([]entity.UserResponse, error) {
	var users []entity.User
	var userResponses []entity.UserResponse

	if err := r.db.Find(&users).Error; err != nil {
		return userResponses, err
	}

	for _, user := range users {
		userResponses = append(userResponses, entity.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userResponses, nil
}

func (r *userRepository) FindUserByID(userID int) (entity.UserResponse, error) {
	var user entity.User

	if err := r.db.First(&user, userID).Error; err != nil {
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

func (r *userRepository) FindUserByUsername(username string) (entity.UserResponse, error) {
	var user entity.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
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

func (r *userRepository) UpdateUser(userID int, user entity.User) (entity.UserResponse, error) {
	var userToUpdate entity.User

	if err := r.db.First(&userToUpdate, userID).Error; err != nil {
		return entity.UserResponse{}, err
	}

	userToUpdate.Username = user.Username
	userToUpdate.Email = user.Email
	userToUpdate.Password = user.Password
	userToUpdate.UpdatedAt = user.UpdatedAt

	if err := r.db.Save(&userToUpdate).Error; err != nil {
		return entity.UserResponse{}, err
	}

	return entity.UserResponse{
		ID:        userToUpdate.ID,
		Username:  userToUpdate.Username,
		Email:     userToUpdate.Email,
		CreatedAt: userToUpdate.CreatedAt,
		UpdatedAt: userToUpdate.UpdatedAt,
	}, nil
}

func (r *userRepository) DeleteUser(userID int) error {
	if err := r.db.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}

	return nil
}
