package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/converter"
)

func (c *UserController) SaveUserSettings(e echo.Context) error {
	user := e.Get("u").(domain.User)

	var r contracts.UserSettingsRequest
	err := e.Bind(&r)
	if err != nil {
		return err
	}
	err = ValidateStruct(r)
	if err != nil {
		return err
	}
	fmt.Println()
	err = c.userService.SaveSettings(e.Request().Context(),
		converter.FromUserSettingsRequestToUserSettings(r), user.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

type UserController struct {
	userService UserService
}
