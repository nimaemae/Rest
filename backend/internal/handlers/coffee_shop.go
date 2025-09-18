package handlers

import (
	"net/http"
	"strconv"

	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"
	"coffee-shop-platform/internal/utils"

	"github.com/labstack/echo/v4"
)

type CoffeeShopHandler struct{}

func NewCoffeeShopHandler() *CoffeeShopHandler {
	return &CoffeeShopHandler{}
}

func (h *CoffeeShopHandler) GetCoffeeShops(c echo.Context) error {
	tenantID, err := strconv.ParseUint(c.Param("tenantId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid tenant ID",
		})
	}

	var coffeeShops []models.CoffeeShop
	if err := database.DB.Where("tenant_id = ?", uint(tenantID)).Find(&coffeeShops).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve coffee shops",
		})
	}

	return c.JSON(http.StatusOK, coffeeShops)
}

func (h *CoffeeShopHandler) CreateCoffeeShop(c echo.Context) error {
	tenantID, err := strconv.ParseUint(c.Param("tenantId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid tenant ID",
		})
	}

	var req models.CoffeeShopCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	coffeeShop := models.CoffeeShop{
		TenantID:     uint(tenantID),
		Name:         req.Name,
		Location:     req.Location,
		Phone:        req.Phone,
		InstagramURL: req.InstagramURL,
		LogoURL:      req.LogoURL,
		HeroImageURL: req.HeroImageURL,
		Description:  req.Description,
		IsActive:     true,
	}

	if err := database.DB.Create(&coffeeShop).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create coffee shop",
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Coffee shop created successfully",
		Data:    coffeeShop,
	})
}

func (h *CoffeeShopHandler) GetCoffeeShop(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid coffee shop ID",
		})
	}

	var coffeeShop models.CoffeeShop
	if err := database.DB.Preload("Tenant").Preload("Admins").First(&coffeeShop, uint(id)).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Coffee shop not found",
		})
	}

	return c.JSON(http.StatusOK, coffeeShop)
}

func (h *CoffeeShopHandler) UpdateCoffeeShop(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid coffee shop ID",
		})
	}

	var req models.CoffeeShopUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var coffeeShop models.CoffeeShop
	if err := database.DB.First(&coffeeShop, uint(id)).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Coffee shop not found",
		})
	}

	if req.Name != nil {
		coffeeShop.Name = *req.Name
	}
	if req.Location != nil {
		coffeeShop.Location = *req.Location
	}
	if req.Phone != nil {
		coffeeShop.Phone = *req.Phone
	}
	if req.InstagramURL != nil {
		coffeeShop.InstagramURL = *req.InstagramURL
	}
	if req.LogoURL != nil {
		coffeeShop.LogoURL = *req.LogoURL
	}
	if req.HeroImageURL != nil {
		coffeeShop.HeroImageURL = *req.HeroImageURL
	}
	if req.Description != nil {
		coffeeShop.Description = *req.Description
	}
	if req.IsActive != nil {
		coffeeShop.IsActive = *req.IsActive
	}

	if err := database.DB.Save(&coffeeShop).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to update coffee shop",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Coffee shop updated successfully",
		Data:    coffeeShop,
	})
}

func (h *CoffeeShopHandler) DeleteCoffeeShop(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid coffee shop ID",
		})
	}

	if err := database.DB.Delete(&models.CoffeeShop{}, uint(id)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete coffee shop",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Coffee shop deleted successfully",
	})
}

func (h *CoffeeShopHandler) CreateShopAdmin(c echo.Context) error {
	shopID, err := strconv.ParseUint(c.Param("shopId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid shop ID",
		})
	}

	var req models.ShopAdminCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to hash password",
		})
	}

	admin := models.ShopAdmin{
		CoffeeShopID: uint(shopID),
		Username:     req.Username,
		PasswordHash: passwordHash,
		IsActive:     true,
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create shop admin",
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Shop admin created successfully",
		Data:    admin,
	})
}
