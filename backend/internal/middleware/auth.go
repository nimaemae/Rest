package middleware

import (
	"net/http"
	"strings"

	"coffee-shop-platform/internal/config"
	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"
	"coffee-shop-platform/internal/utils"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
					Error: "Authorization header required",
				})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
					Error: "Bearer token required",
				})
			}

			cfg := c.Get("config").(*config.Config)
			claims, err := utils.ParseJWT(tokenString, cfg.JWT.Secret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
					Error: "Invalid token",
				})
			}

			// Store user info in context
			c.Set("user_id", (*claims)["user_id"])
			c.Set("username", (*claims)["username"])
			c.Set("user_type", (*claims)["type"])
			c.Set("shop_id", (*claims)["shop_id"])

			return next(c)
		}
	}
}

func MainAdminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userType := c.Get("user_type").(string)
			if userType != "main_admin" {
				return c.JSON(http.StatusForbidden, models.ErrorResponse{
					Error: "Main admin access required",
				})
			}
			return next(c)
		}
	}
}

func ShopAdminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userType := c.Get("user_type").(string)
			if userType != "shop_admin" {
				return c.JSON(http.StatusForbidden, models.ErrorResponse{
					Error: "Shop admin access required",
				})
			}
			return next(c)
		}
	}
}

func TenantResolver() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract subdomain from Host header
			host := c.Request().Host
			parts := strings.Split(host, ".")
			
			if len(parts) >= 2 {
				subdomain := parts[0]
				
				// Skip if it's localhost or IP
				if subdomain != "localhost" && subdomain != "127" && subdomain != "0" {
					var tenant models.Tenant
					if err := database.DB.Where("subdomain = ? AND is_active = ?", subdomain, true).First(&tenant).Error; err == nil {
						c.Set("tenant_id", tenant.ID)
						c.Set("tenant", tenant)
					}
				}
			}

			return next(c)
		}
	}
}
