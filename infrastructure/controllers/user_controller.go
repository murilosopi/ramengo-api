package controllers

import (
	"net/http"
	"ramengo/application/dtos/address"
	"ramengo/application/dtos/user"
	"ramengo/application/services"
	"ramengo/infrastructure/security"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type UserController struct {
	userService    services.UserService
	addressService services.AddressService
	validator      *validator.Validate
}

func NewUserController(us services.UserService, as services.AddressService) UserController {
	return UserController{
		userService:    us,
		addressService: as,
		validator:      validator.New(),
	}
}

func (uc UserController) Save(c echo.Context) error {
	// temporary helper struct
	var dtoHelper struct {
		Name      string                    `json:"name"`
		Email     string                    `json:"email"`
		Password  string                    `json:"password"`
		Address   *address.CreateAddressDTO `json:"address"`
		AddressID int                       `json:"addressID"`
	}

	// decode json to struct
	if err := c.Bind(&dtoHelper); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// return bad request if no address sent
	if dtoHelper.AddressID == 0 && dtoHelper.Address == nil {
		return c.JSON(http.StatusBadRequest, false) // fail response
	}

	// instance of dto user creationn
	userDTO := user.CreateUserDTO{
		Name:     dtoHelper.Name,
		Email:    dtoHelper.Email,
		Password: dtoHelper.Password,
	}

	if err := uc.validator.Struct(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "malformed or empty user required fields",
		})
	}

	addressId := dtoHelper.AddressID
	if addressId == 0 {
		if err := uc.validator.Struct(*dtoHelper.Address); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "malformed or empty address required fields",
			})
		}

		_, addressId = uc.addressService.Save(dtoHelper.Address)
	}

	userDTO.AddressID = addressId

	// call saving user service
	if tokenJWT, err := uc.userService.Save(&userDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()}) // fail response
	} else {
		return c.JSON(http.StatusCreated, echo.Map{
			"token": tokenJWT,
		}) // success response
	}
}

func (uc UserController) OrderHistory(c echo.Context) error {
	payload := c.Get(string(security.Kitchen)).(*security.JWTClaims)

	userId := payload.Id

	result := uc.userService.OrderHistory(userId)

	status := http.StatusOK
	if len(result) == 0 {
		status = http.StatusNotFound
	}

	return c.JSON(status, result)

}
