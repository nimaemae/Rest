package handlers

import (
	"net/http"

	"coffee-shop-platform/internal/config"
	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"
	"coffee-shop-platform/internal/utils"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) MainAdminLogin(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var admin models.MainAdmin
	if err := database.DB.Where("username = ? AND is_active = ?", req.Username, true).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid credentials",
		})
	}

	if !utils.CheckPasswordHash(req.Password, admin.PasswordHash) {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid credentials",
		})
	}

	cfg := c.Get("config").(*config.Config)
	token, err := utils.GenerateJWT(admin.ID, admin.Username, "main_admin", nil, cfg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
		User:  admin,
	})
}

func (h *AuthHandler) ShopAdminLogin(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var admin models.ShopAdmin
	if err := database.DB.Preload("CoffeeShop").Where("username = ? AND is_active = ?", req.Username, true).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid credentials",
		})
	}

	if !utils.CheckPasswordHash(req.Password, admin.PasswordHash) {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid credentials",
		})
	}

	cfg := c.Get("config").(*config.Config)
	token, err := utils.GenerateJWT(admin.ID, admin.Username, "shop_admin", &admin.CoffeeShopID, cfg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
		User:  admin,
	})
}
