package middleware

import (
	"bytes"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/madsilver/template-bff/internal/infra/server/utils"
)

func BodySanitize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var body []byte
			if c.Request().Body != nil {
				body, _ = io.ReadAll(c.Request().Body)

				utils.NewSanitize().Do(&body)

				c.Request().Body = io.NopCloser(bytes.NewBuffer(body))
			}
			return next(c)
		}
	}
}
