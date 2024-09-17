package services

import (
	"errors"
	"ramengo/application/dtos/auth"
	"ramengo/domain/repositories"
	"ramengo/infrastructure/security"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return AuthService{repo}
}

func (ks *AuthService) Login(dto *auth.LoginDTO) (token string, err error) {

	user := ks.authRepo.FindUserByEmail(dto.Email)

	// user not found
	if user == nil {
		return "", errors.New("user e-mail not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		return "", errors.New("invalid credential")
	}

	claims := security.JWTClaims{
		Id:   user.Id,
		Type: security.User,
	}

	// validate kitchen access
	if dto.Kitchen != 0 {
		kitchen := ks.authRepo.GetUserKitchenByID(user.Id, dto.Kitchen)

		if kitchen == nil {
			return "", errors.New("kitchen access unauthorized")
		}

		claims.Id = kitchen.Id
		claims.Type = security.Kitchen
	}

	token = security.GenerateTokenJWT(claims.Id, claims.Type)
	return
}
