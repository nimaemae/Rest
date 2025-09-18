package main

import (
	"flag"
	"fmt"
	"log"

	"coffee-shop-platform/internal/config"
	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/routes"
	"coffee-shop-platform/scripts"

	"github.com/labstack/echo/v4"
)

func main() {
	migratePtr := flag.Bool("migrate", false, "Run database migrations and exit")
	seedPtr := flag.Bool("seed", false, "Seed database with sample data and exit")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	if *migratePtr {
		fmt.Println("Running database migrations...")
		if err := database.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migrations completed successfully!")
		return
	}

	if *seedPtr {
		fmt.Println("Seeding database with sample data...")
		if err := database.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		if err := scripts.SeedDatabase(); err != nil {
			log.Fatalf("Seeding failed: %v", err)
		}
		fmt.Println("Seeding completed successfully!")
		return
	}

	e := echo.New()

	// Store config in context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", cfg)
			c.Set("db", database.DB)
			return next(c)
		}
	})

	routes.SetupRoutes(e)

	log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
