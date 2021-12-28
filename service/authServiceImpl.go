package service

import (
	"bank-test/dto"
	"bank-test/entity"
	"bank-test/repo"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewAuthService(authRepo *repo.AuthRepo) AuthService {
	return &AuthServiceImpl{
		AuthRepo: *authRepo,
	}
}

type AuthServiceImpl struct {
	AuthRepo repo.AuthRepo
}

func (service *AuthServiceImpl) Login(username, password string) (string, error) {
	customer, err := service.AuthRepo.Login(username, password)
	if err != nil {
		return "", err
	}

	token := service.GenerateToken(customer)

	//save token in db
	err = service.AuthRepo.SaveToken(token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthServiceImpl) GenerateToken(customer entity.Customer) string {
	credential := dto.UserDto{
		Id:       customer.GetId(),
		Username: customer.GetUsername(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credential)
	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func (service *AuthServiceImpl) Logout(token string) error {
	return service.AuthRepo.Logout(token)
}
