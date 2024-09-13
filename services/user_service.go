package services

import (
	"errors"
	"fiber-crud-auth/models"
	"fiber-crud-auth/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (string, error)
}

type authService struct{
	repo repositories.UserRepository
	jwtSecret string
}

func NewAuthService(repo repositories.UserRepository, jwtSecret string) AuthService{
	return &authService{repo,jwtSecret}
}

func (s *authService) RegisterUser(user *models.User)error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}

func (s *authService) LoginUser(email,password string) (string,error){
	user,err := s.repo.FindByEmail(email)
	if err != nil{
		return "", errors.New("user not found")
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));
	err !=nil {
		return "",errors.New("Invalid Credentials")
	}

	//Buat Token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":user.ID,
		"exp":time.Now().Add(time.Hour * 72).Unix(),
	})

	return	token.SignedString([]byte(s.jwtSecret))
}