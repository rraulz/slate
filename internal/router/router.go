package router

import (
	"net/http"
	"slate/internal/controller"
	"slate/templates/login"

	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	userController *controller.UserController
}

func NewAPIRouter(userController *controller.UserController) *APIRouter {
	return &APIRouter{
		userController: userController,
	}
}

func (a *APIRouter) RegisterRouters(g *gin.Engine) {

	// Routes
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", login.Page())
	})

	// g.POST("/login", a.userController.LoginUserController)
	g.POST("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", login.LoginForm("aasd", "aas", "aasd"))
	})

}
