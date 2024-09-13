package repositories

import (
	"fiber-crud-auth/databases"
	"fiber-crud-auth/models"
)

type UserRepository interface {
    Create(user *models.User) error
    FindByEmail(email string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
    return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
    return databases.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := databases.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}
