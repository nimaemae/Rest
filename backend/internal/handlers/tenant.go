package handlers

import (
	"net/http"
	"strconv"

	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"

	"github.com/labstack/echo/v4"
)

type TenantHandler struct{}

func NewTenantHandler() *TenantHandler {
	return &TenantHandler{}
}

func (h *TenantHandler) GetTenants(c echo.Context) error {
	var tenants []models.Tenant
	if err := database.DB.Find(&tenants).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve tenants",
		})
	}

	return c.JSON(http.StatusOK, tenants)
}

func (h *TenantHandler) CreateTenant(c echo.Context) error {
	var req models.TenantCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	tenant := models.Tenant{
		Subdomain: req.Subdomain,
		Name:      req.Name,
		IsActive:  true,
	}

	if err := database.DB.Create(&tenant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create tenant",
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Tenant created successfully",
		Data:    tenant,
	})
}

func (h *TenantHandler) GetTenant(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid tenant ID",
		})
	}

	var tenant models.Tenant
	if err := database.DB.Preload("CoffeeShops").First(&tenant, uint(id)).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Tenant not found",
		})
	}

	return c.JSON(http.StatusOK, tenant)
}

func (h *TenantHandler) UpdateTenant(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid tenant ID",
		})
	}

	var req models.TenantUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var tenant models.Tenant
	if err := database.DB.First(&tenant, uint(id)).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Tenant not found",
		})
	}

	if req.Name != nil {
		tenant.Name = *req.Name
	}
	if req.IsActive != nil {
		tenant.IsActive = *req.IsActive
	}

	if err := database.DB.Save(&tenant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to update tenant",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Tenant updated successfully",
		Data:    tenant,
	})
}

func (h *TenantHandler) DeleteTenant(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid tenant ID",
		})
	}

	if err := database.DB.Delete(&models.Tenant{}, uint(id)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete tenant",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Tenant deleted successfully",
	})
}
