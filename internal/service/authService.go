package service

import (
	"slate/internal/errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

type LoginData struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (a *AuthService) LoginUser(loginData *LoginData) (string, error) {
	user, err := a.userService.GetUserByUsernameService(loginData.Username)
	if err != nil {
		return "", errors.NewServiceError(err, errors.ErrBadRequest)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		// Passwords don't match
		return "", errors.NewServiceError(err, errors.ErrUnauthorized)
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
		return "", errors.NewServiceError(err, errors.ErrInternalError)
	}

	return t, nil
}
