package restful

import (
	"net/http"

	"github.com/gogotsenghsien/simple-rate-limit/src/api/restful/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	postHandler *handlers.PostHandler
}

func (s *Server) Run() {
	e := echo.New()

	// add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// add routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome Here")
	})
	e.POST("/post", s.postHandler.AddPost)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}

func NewServer(postHandler *handlers.PostHandler) *Server {
	return &Server{
		postHandler: postHandler,
	}
}
