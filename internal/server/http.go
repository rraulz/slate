package server

import (
	"context"
	"slate/internal/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func NewHTTPServer(ctx context.Context, apiRouter *router.APIRouter) *Server {

	g := gin.Default()
	g.HTMLRender = &TemplRender{}
	// rand.Static("/assets", "assets")

	apiRouter.RegisterRouters(g)

	return &Server{g}
}

func (s *Server) Start() {
	s.Gin.Run(":42069")
}
