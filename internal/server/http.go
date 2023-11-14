package server

import (
	"context"
	"fmt"
	"slate/internal/router"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	Ctx  context.Context
}

func NewHTTPServer(ctx context.Context, apiRouter *router.APIRouter) *Server {

	e := echo.New()
	e.Static("/assets", "assets")
	apiRouter.RegisterRouters(e, ctx)

	fmt.Println("Server initialized")

	return &Server{e, ctx}
}

func (s *Server) Start() {
	s.Echo.Logger.Fatal(s.Echo.Start(":1323"))
}
