package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginData struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func LoginHandler(c echo.Context) error {
	loginData := &LoginData{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	if err := c.Bind(loginData); err != nil {
		return err
	}

	// Validate login data
	if loginData.Username == "" || loginData.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid login data"})
	}

	// If login data is valid
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})
}
