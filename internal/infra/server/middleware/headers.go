package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/madsilver/template-bff/internal/domain/dto"
)

func HeaderMapper() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tid := c.Request().Header.Get("x-tid")
			if tid == "" {
				tid = uuid.New().String()
				c.Request().Header.Set("x-tid", tid)
			}

			c.Response().Header().Set("x-tid", tid)
			
			ctx := context.WithValue(c.Request().Context(), dto.RequestContextKey, dto.RequestContext{
				Tid:   tid,
				Roles: extractRoles(c.Request().Header),
			})

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func extractRoles(header http.Header) []string {
	sRoles := strings.Split(header.Get("x-roles"), ",")
	roles := make([]string, 0)
	for _, role := range sRoles {
		roles = append(roles, strings.TrimSpace(role))
	}
	return roles
}
