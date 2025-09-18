package routes

import (
	"coffee-shop-platform/internal/handlers"
	"coffee-shop-platform/internal/middleware"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo) {
	// Initialize handlers
	authHandler := handlers.NewAuthHandler()
	tenantHandler := handlers.NewTenantHandler()
	coffeeShopHandler := handlers.NewCoffeeShopHandler()
	menuHandler := handlers.NewMenuHandler()
	categoryHandler := handlers.NewCategoryHandler()

	// CORS middleware
	e.Use(echomiddleware.CORS())

	// Rate limiting
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(20)))

	// Request logging
	e.Use(echomiddleware.Logger())

	// Recovery middleware
	e.Use(echomiddleware.Recover())

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Public routes (no authentication required)
	public := e.Group("/api/public")
	public.GET("/menu", menuHandler.GetPublicMenuItems, middleware.TenantResolver())
	public.GET("/shop", menuHandler.GetShopSettings, middleware.TenantResolver())
	public.GET("/categories", categoryHandler.GetCategories)

	// Authentication routes
	auth := e.Group("/api/auth")
	auth.POST("/main-admin/login", authHandler.MainAdminLogin)
	auth.POST("/shop-admin/login", authHandler.ShopAdminLogin)

	// Main admin routes (require main admin authentication)
	mainAdmin := e.Group("/api/admin")
	mainAdmin.Use(middleware.AuthMiddleware())
	mainAdmin.Use(middleware.MainAdminOnly())

	// Tenant management
	mainAdmin.GET("/tenants", tenantHandler.GetTenants)
	mainAdmin.POST("/tenants", tenantHandler.CreateTenant)
	mainAdmin.GET("/tenants/:id", tenantHandler.GetTenant)
	mainAdmin.PUT("/tenants/:id", tenantHandler.UpdateTenant)
	mainAdmin.DELETE("/tenants/:id", tenantHandler.DeleteTenant)

	// Coffee shop management
	mainAdmin.GET("/tenants/:tenantId/shops", coffeeShopHandler.GetCoffeeShops)
	mainAdmin.POST("/tenants/:tenantId/shops", coffeeShopHandler.CreateCoffeeShop)
	mainAdmin.GET("/shops/:id", coffeeShopHandler.GetCoffeeShop)
	mainAdmin.PUT("/shops/:id", coffeeShopHandler.UpdateCoffeeShop)
	mainAdmin.DELETE("/shops/:id", coffeeShopHandler.DeleteCoffeeShop)

	// Shop admin management
	mainAdmin.POST("/shops/:shopId/admins", coffeeShopHandler.CreateShopAdmin)

	// Category management (main admin only)
	mainAdmin.GET("/categories", categoryHandler.GetAllCategories)
	mainAdmin.POST("/categories", categoryHandler.CreateCategory)
	mainAdmin.GET("/categories/:id", categoryHandler.GetCategory)
	mainAdmin.PUT("/categories/:id", categoryHandler.UpdateCategory)
	mainAdmin.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Shop admin routes (require shop admin authentication and tenant resolution)
	shopAdmin := e.Group("/api/admin")
	shopAdmin.Use(middleware.TenantResolver())
	shopAdmin.Use(middleware.AuthMiddleware())
	shopAdmin.Use(middleware.ShopAdminOnly())

	// Menu management
	shopAdmin.GET("/menu", menuHandler.GetMenuItems)
	shopAdmin.POST("/menu", menuHandler.CreateMenuItem)
	shopAdmin.GET("/menu/:id", menuHandler.GetMenuItem)
	shopAdmin.PUT("/menu/:id", menuHandler.UpdateMenuItem)
	shopAdmin.DELETE("/menu/:id", menuHandler.DeleteMenuItem)

	// Shop settings
	shopAdmin.GET("/settings", menuHandler.GetShopSettings)
	shopAdmin.PUT("/settings", menuHandler.UpdateShopSettings)

	// Categories (read-only for shop admins)
	shopAdmin.GET("/categories", categoryHandler.GetCategories)
}
