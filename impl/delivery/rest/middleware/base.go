package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORS()
}

func Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}

func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func Timeout() echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	})
}
