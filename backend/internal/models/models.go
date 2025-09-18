package models

import (
	"time"

	"gorm.io/gorm"
)

// Tenant represents the main tenant (platform owner)
type Tenant struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Subdomain string         `json:"subdomain" gorm:"uniqueIndex;not null"`
	Name      string         `json:"name" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	CoffeeShops []CoffeeShop `json:"coffee_shops,omitempty" gorm:"foreignKey:TenantID"`
}

// CoffeeShop represents a coffee shop under a tenant
type CoffeeShop struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	TenantID     uint           `json:"tenant_id" gorm:"not null"`
	Name         string         `json:"name" gorm:"not null"`
	Location     string         `json:"location"`
	Phone        string         `json:"phone"`
	InstagramURL string         `json:"instagram_url"`
	LogoURL      string         `json:"logo_url"`
	HeroImageURL string         `json:"hero_image_url"`
	Description  string         `json:"description"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Tenant     Tenant      `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	Admins     []ShopAdmin `json:"admins,omitempty" gorm:"foreignKey:CoffeeShopID"`
	MenuItems  []MenuItem  `json:"menu_items,omitempty" gorm:"foreignKey:CoffeeShopID"`
}

// ShopAdmin represents admin users for coffee shops
type ShopAdmin struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CoffeeShopID uint           `json:"coffee_shop_id" gorm:"not null"`
	Username     string         `json:"username" gorm:"not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	CoffeeShop CoffeeShop `json:"coffee_shop,omitempty" gorm:"foreignKey:CoffeeShopID"`
}

// MenuItem represents menu items for coffee shops
type MenuItem struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CoffeeShopID   uint           `json:"coffee_shop_id" gorm:"not null"`
	CategoryID     uint           `json:"category_id" gorm:"not null"`
	Name           string         `json:"name" gorm:"not null"`
	Price          int           `json:"price" gorm:"not null"`
	PricePremium    *int          `json:"price_premium"`
	HasDualPricing bool           `json:"has_dual_pricing" gorm:"default:false"`
	ImageURL       string         `json:"image_url"`
	OrderIndex     int            `json:"order_index" gorm:"default:0"`
	IsAvailable    bool           `json:"is_available" gorm:"default:true"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	CoffeeShop CoffeeShop `json:"coffee_shop,omitempty" gorm:"foreignKey:CoffeeShopID"`
	Category   Category   `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}

// MainAdmin represents the platform's main admin
type MainAdmin struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Request/Response DTOs

// TenantCreateRequest represents the request to create a tenant
type TenantCreateRequest struct {
	Subdomain string `json:"subdomain" validate:"required,min=3,max=50"`
	Name       string `json:"name" validate:"required,min=2,max=100"`
}

// TenantUpdateRequest represents the request to update a tenant
type TenantUpdateRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// CoffeeShopCreateRequest represents the request to create a coffee shop
type CoffeeShopCreateRequest struct {
	Name         string  `json:"name" validate:"required,min=2,max=100"`
	Location     string  `json:"location" validate:"omitempty,max=200"`
	Phone        string  `json:"phone" validate:"omitempty,max=20"`
	InstagramURL string  `json:"instagram_url" validate:"omitempty,url"`
	LogoURL      string  `json:"logo_url" validate:"omitempty,url"`
	HeroImageURL string  `json:"hero_image_url" validate:"omitempty,url"`
	Description  string  `json:"description" validate:"omitempty,max=500"`
}

// CoffeeShopUpdateRequest represents the request to update a coffee shop
type CoffeeShopUpdateRequest struct {
	Name         *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Location     *string `json:"location,omitempty" validate:"omitempty,max=200"`
	Phone        *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	InstagramURL *string `json:"instagram_url,omitempty" validate:"omitempty,url"`
	LogoURL      *string `json:"logo_url,omitempty" validate:"omitempty,url"`
	HeroImageURL *string `json:"hero_image_url,omitempty" validate:"omitempty,url"`
	Description  *string `json:"description,omitempty" validate:"omitempty,max=500"`
	IsActive     *bool   `json:"is_active,omitempty"`
}

// ShopAdminCreateRequest represents the request to create a shop admin
type ShopAdminCreateRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

// ShopAdminUpdateRequest represents the request to update a shop admin
type ShopAdminUpdateRequest struct {
	Username *string `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=6,max=100"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// MenuItemCreateRequest represents the request to create a menu item
type MenuItemCreateRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	CategoryID     uint   `json:"category_id" validate:"required"`
	Price          int    `json:"price" validate:"required,min=0"`
	PricePremium   *int   `json:"price_premium" validate:"omitempty,min=0"`
	HasDualPricing bool   `json:"has_dual_pricing"`
	ImageURL       string `json:"image_url" validate:"omitempty,url"`
	OrderIndex     int    `json:"order_index" validate:"min=0"`
	IsAvailable    bool   `json:"is_available"`
}

// MenuItemUpdateRequest represents the request to update a menu item
type MenuItemUpdateRequest struct {
	Name           *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	CategoryID     *uint   `json:"category_id,omitempty"`
	Price          *int    `json:"price,omitempty" validate:"omitempty,min=0"`
	PricePremium   *int    `json:"price_premium,omitempty" validate:"omitempty,min=0"`
	HasDualPricing *bool   `json:"has_dual_pricing,omitempty"`
	ImageURL       *string `json:"image_url,omitempty" validate:"omitempty,url"`
	OrderIndex     *int    `json:"order_index,omitempty" validate:"omitempty,min=0"`
	IsAvailable    *bool   `json:"is_available,omitempty"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  any    `json:"user"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
