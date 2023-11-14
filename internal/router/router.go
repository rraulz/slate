package router

import (
	"context"
	"slate/internal/controller"
	"slate/templates/login"

	"github.com/labstack/echo/v4"
)

type APIRouter struct {
	userController *controller.UserController
}

func NewAPIRouter(userController *controller.UserController) *APIRouter {
	return &APIRouter{
		userController: userController,
	}
}

func (a *APIRouter) RegisterRouters(e *echo.Echo, ctx context.Context) {

	// Routes
	e.GET("/", func(c echo.Context) error {
		component := login.Page()
		return component.Render(ctx, c.Response().Writer)
	})

	e.POST("/login", a.userController.LoginUser)

	// Start server
}
