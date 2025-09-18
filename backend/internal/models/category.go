package models

import (
	"time"

	"gorm.io/gorm"
)

// Category represents menu categories managed by main admin
type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	DisplayName string         `json:"display_name" gorm:"not null"`
	Emoji       string         `json:"emoji"`
	Color       string         `json:"color"`
	OrderIndex  int            `json:"order_index" gorm:"default:0"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	MenuItems []MenuItem `json:"menu_items,omitempty" gorm:"foreignKey:CategoryID"`
}

// CategoryCreateRequest represents the request to create a category
type CategoryCreateRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	DisplayName string `json:"display_name" validate:"required,min=2,max=100"`
	Emoji       string `json:"emoji" validate:"omitempty,max=10"`
	Color       string `json:"color" validate:"omitempty,max=50"`
	OrderIndex  int    `json:"order_index" validate:"min=0"`
}

// CategoryUpdateRequest represents the request to update a category
type CategoryUpdateRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=50"`
	DisplayName *string `json:"display_name,omitempty" validate:"omitempty,min=2,max=100"`
	Emoji       *string `json:"emoji,omitempty" validate:"omitempty,max=10"`
	Color       *string `json:"color,omitempty" validate:"omitempty,max=50"`
	OrderIndex  *int    `json:"order_index,omitempty" validate:"omitempty,min=0"`
	IsActive    *bool   `json:"is_active,omitempty"`
}
