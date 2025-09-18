package handlers

import (
	"net/http"
	"strconv"

	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct{}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// GetCategories retrieves all active categories
func (h *CategoryHandler) GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := database.DB.Where("is_active = ?", true).Order("order_index ASC").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve categories",
		})
	}

	return c.JSON(http.StatusOK, categories)
}

// GetAllCategories retrieves all categories (including inactive) for admin
func (h *CategoryHandler) GetAllCategories(c echo.Context) error {
	var categories []models.Category
	if err := database.DB.Order("order_index ASC").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to retrieve categories",
		})
	}

	return c.JSON(http.StatusOK, categories)
}

// CreateCategory creates a new category
func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var req models.CategoryCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Check if category name already exists
	var existingCategory models.Category
	if err := database.DB.Where("name = ?", req.Name).First(&existingCategory).Error; err == nil {
		return c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: "Category name already exists",
		})
	}

	category := models.Category{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Emoji:       req.Emoji,
		Color:       req.Color,
		OrderIndex:  req.OrderIndex,
		IsActive:    true,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create category",
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Category created successfully",
		Data:    category,
	})
}

// GetCategory retrieves a specific category
func (h *CategoryHandler) GetCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid category ID",
		})
	}

	var category models.Category
	if err := database.DB.Preload("MenuItems").First(&category, uint(id)).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Category not found",
		})
	}

	return c.JSON(http.StatusOK, category)
}

// UpdateCategory updates a category
func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid category ID",
		})
	}

	var req models.CategoryUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	var category models.Category
	if err := database.DB.First(&category, uint(id)).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Category not found",
		})
	}

	// Check if new name conflicts with existing category
	if req.Name != nil && *req.Name != category.Name {
		var existingCategory models.Category
		if err := database.DB.Where("name = ? AND id != ?", *req.Name, uint(id)).First(&existingCategory).Error; err == nil {
			return c.JSON(http.StatusConflict, models.ErrorResponse{
				Error: "Category name already exists",
			})
		}
	}

	// Update fields if provided
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.DisplayName != nil {
		category.DisplayName = *req.DisplayName
	}
	if req.Emoji != nil {
		category.Emoji = *req.Emoji
	}
	if req.Color != nil {
		category.Color = *req.Color
	}
	if req.OrderIndex != nil {
		category.OrderIndex = *req.OrderIndex
	}
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to update category",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Category updated successfully",
		Data:    category,
	})
}

// DeleteCategory deletes a category (soft delete)
func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid category ID",
		})
	}

	// Check if category has menu items
	var count int64
	database.DB.Model(&models.MenuItem{}).Where("category_id = ?", uint(id)).Count(&count)
	if count > 0 {
		return c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: "Cannot delete category with existing menu items",
		})
	}

	if err := database.DB.Delete(&models.Category{}, uint(id)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete category",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Category deleted successfully",
	})
}
