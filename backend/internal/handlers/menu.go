package handlers

import (
	"net/http"
	"strconv"

	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"

	"github.com/labstack/echo/v4"
)

type MenuHandler struct{}

func NewMenuHandler() *MenuHandler {
	return &MenuHandler{}
}

func (h *MenuHandler) GetPublicMenuItems(c echo.Context) error {
	tenantID := c.Get("tenant_id")
	if tenantID == nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Tenant not found",
		})
	}

	var menuItems []models.MenuItem
	if err := database.DB.Preload("Category").Where("coffee_shop_id IN (SELECT id FROM coffee_shops WHERE tenant_id = ?) AND is_available = ?", tenantID, true).Order("order_index ASC").Find(&menuItems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve menu items",
		})
	}

	return c.JSON(http.StatusOK, menuItems)
}

func (h *MenuHandler) GetMenuItems(c echo.Context) error {
	shopID := c.Get("shop_id").(uint)
	
	var menuItems []models.MenuItem
	if err := database.DB.Preload("Category").Where("coffee_shop_id = ?", shopID).Order("order_index ASC").Find(&menuItems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve menu items",
		})
	}

	return c.JSON(http.StatusOK, menuItems)
}

func (h *MenuHandler) CreateMenuItem(c echo.Context) error {
	shopID := c.Get("shop_id").(uint)
	
	var req models.MenuItemCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	menuItem := models.MenuItem{
		CoffeeShopID:   shopID,
		CategoryID:     req.CategoryID,
		Name:           req.Name,
		Price:          req.Price,
		PricePremium:   req.PricePremium,
		HasDualPricing: req.HasDualPricing,
		ImageURL:       req.ImageURL,
		OrderIndex:     req.OrderIndex,
		IsAvailable:    req.IsAvailable,
	}

	if err := database.DB.Create(&menuItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create menu item",
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Menu item created successfully",
		Data:    menuItem,
	})
}

func (h *MenuHandler) GetMenuItem(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid menu item ID",
		})
	}

	shopID := c.Get("shop_id").(uint)
	
	var menuItem models.MenuItem
	if err := database.DB.Preload("Category").Where("id = ? AND coffee_shop_id = ?", uint(id), shopID).First(&menuItem).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Menu item not found",
		})
	}

	return c.JSON(http.StatusOK, menuItem)
}

func (h *MenuHandler) UpdateMenuItem(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid menu item ID",
		})
	}

	shopID := c.Get("shop_id").(uint)
	
	var req models.MenuItemUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var menuItem models.MenuItem
	if err := database.DB.Where("id = ? AND coffee_shop_id = ?", uint(id), shopID).First(&menuItem).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Menu item not found",
		})
	}

	if req.Name != nil {
		menuItem.Name = *req.Name
	}
	if req.CategoryID != nil {
		menuItem.CategoryID = *req.CategoryID
	}
	if req.Price != nil {
		menuItem.Price = *req.Price
	}
	if req.PricePremium != nil {
		menuItem.PricePremium = req.PricePremium
	}
	if req.HasDualPricing != nil {
		menuItem.HasDualPricing = *req.HasDualPricing
	}
	if req.ImageURL != nil {
		menuItem.ImageURL = *req.ImageURL
	}
	if req.OrderIndex != nil {
		menuItem.OrderIndex = *req.OrderIndex
	}
	if req.IsAvailable != nil {
		menuItem.IsAvailable = *req.IsAvailable
	}

	if err := database.DB.Save(&menuItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to update menu item",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Menu item updated successfully",
		Data:    menuItem,
	})
}

func (h *MenuHandler) DeleteMenuItem(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid menu item ID",
		})
	}

	shopID := c.Get("shop_id").(uint)
	
	if err := database.DB.Where("id = ? AND coffee_shop_id = ?", uint(id), shopID).Delete(&models.MenuItem{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete menu item",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Menu item deleted successfully",
	})
}

func (h *MenuHandler) GetShopSettings(c echo.Context) error {
	tenantID := c.Get("tenant_id")
	if tenantID == nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Tenant not found",
		})
	}

	var coffeeShop models.CoffeeShop
	if err := database.DB.Where("tenant_id = ?", tenantID).First(&coffeeShop).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Coffee shop not found",
		})
	}

	return c.JSON(http.StatusOK, coffeeShop)
}

func (h *MenuHandler) UpdateShopSettings(c echo.Context) error {
	shopID := c.Get("shop_id").(uint)
	
	var req models.CoffeeShopUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var coffeeShop models.CoffeeShop
	if err := database.DB.First(&coffeeShop, shopID).Error; err != nil {
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
			Error: "Failed to update shop settings",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Shop settings updated successfully",
		Data:    coffeeShop,
	})
}
