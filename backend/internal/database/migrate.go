package database

import (
	"coffee-shop-platform/internal/models"
)

func Migrate() error {
	return DB.AutoMigrate(
		&models.MainAdmin{},
		&models.Tenant{},
		&models.CoffeeShop{},
		&models.ShopAdmin{},
		&models.Category{},
		&models.MenuItem{},
	)
}
