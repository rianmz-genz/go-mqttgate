package service

import (
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"
	"crypto/sha256"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewAuthService(UserRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: UserRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.LoginRequest) string {
	userRequest := domain.User{Email: request.Email}
	userResponse, err := service.UserRepository.GetUserByEmail(ctx, service.DB, userRequest)
	if err != nil {
		return "user not found"
	}
	
	hash := sha256.New()
	hash.Write([]byte(request.Password))
	encryptedRequestPassword := hash.Sum(nil)

	if userResponse.Email != string(encryptedRequestPassword) {
		return "wrong password"
	}

	println(userResponse.Name)

	return "token"
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.RegisterRequest) {

}
