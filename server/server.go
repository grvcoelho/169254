package server

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Server struct {
	app  *iris.Application
	port string
}

func New() *Server {
	app := iris.New()

	server := &Server{
		app:  app,
		port: "8080",
	}

	server.initMiddlewares()
	server.initHandlers()

	return server
}

func (s *Server) initMiddlewares() {
	s.app.Logger().SetLevel("debug")
	s.app.Use(recover.New())
	s.app.Use(logger.New())
}

func (s *Server) initHandlers() {
	s.app.Get("/_health_check", func(ctx iris.Context) {
		ctx.StatusCode(200)
		ctx.WriteString("ok")
	})
}

func (s *Server) Start() {
	s.app.Run(iris.Addr(":8080"))
}
