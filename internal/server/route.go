package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) routeApiV1(r *echo.Echo) {

	auth := r.Group("/api/v1/")
	{
		auth.POST("auth/sign-up", s.handlers.university.SignUp)
		auth.POST("auth/sign-in", s.handlers.university.SignIn)
	}

}

func (s *Server) routeSwagger(r *echo.Echo) {
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
