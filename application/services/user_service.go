package services

import (
	"errors"
	"ramengo/application/dtos/order"
	"ramengo/application/dtos/user"
	"ramengo/domain/models"
	"ramengo/domain/repositories"
	"ramengo/infrastructure/security"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return UserService{repo}
}

func (us *UserService) Save(dto *user.CreateUserDTO) (token string, err error) {
	address := &models.AddressModel{Id: dto.AddressID}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)

	if err != nil {
		return "", err
	}

	if !us.userRepo.VerifyEmailAvailable(dto.Email) {
		return "", errors.New("e-mail already in use")
	}

	user := models.UserModel{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: passwordHashed,
		Address:  address,
	}

	if us.userRepo.Save(&user) {
		token = security.GenerateTokenJWT(user.Id, security.User)
	} else {
		err = errors.New("failure user data storage")
	}

	return
}

func (us *UserService) OrderHistory(userId int) (result []order.GetOrderDTO) {
	orders := us.userRepo.OrderHistory(userId)

	for _, o := range orders {
		result = append(result, order.GetOrderDTO{
			Id:        o.Id,
			KitchenId: o.Kitchen.Id,
			Date:      o.Date.Format(time.DateOnly),
			UserId:    o.User.Id,
			UserName:  o.User.Name,
			Status:    o.Status.String(),
			StatusId:  int(o.Status),
		})
	}

	return
}
