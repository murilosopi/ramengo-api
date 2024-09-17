package controllers

import (
	"net/http"
	"ramengo/application/dtos/auth"
	"ramengo/application/services"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type AuthController struct {
	AuthService services.AuthService
	validate    *validator.Validate
}

func NewAuthController(as services.AuthService) AuthController {
	return AuthController{as, validator.New()}
}

func (ac AuthController) Login(c echo.Context) error {

	loginDTO := new(auth.LoginDTO)

	if err := c.Bind(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ac.validate.Struct(loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "malformed fields"})
	}

	token, err := ac.AuthService.Login(loginDTO)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
