package controller

import (
	"net/http"
	"slate/internal/service"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

type LoginData struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (uc *UserController) LoginUser(c echo.Context) error {
	loginData := &LoginData{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	user, err := uc.userService.GetUserByUsernameService()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		// Passwords don't match
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
	}

	// If login data is valid, generate JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Interanl server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}
