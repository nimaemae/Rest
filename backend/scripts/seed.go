package scripts

import (
	"log"

	"coffee-shop-platform/internal/config"
	"coffee-shop-platform/internal/database"
	"coffee-shop-platform/internal/models"
	"coffee-shop-platform/internal/utils"
)

func SeedDatabase() error {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		return err
	}
	defer database.Close()

	// Run database migrations
	if err := database.Migrate(); err != nil {
		return err
	}

	// Seed main admin
	if err := seedMainAdmin(); err != nil {
		return err
	}

	// Seed categories
	if err := seedCategories(); err != nil {
		return err
	}

	// Seed sample tenant
	if err := seedSampleTenant(); err != nil {
		return err
	}

	log.Println("Database seeded successfully!")
	return nil
}

func seedMainAdmin() error {
	var count int64
	database.DB.Model(&models.MainAdmin{}).Count(&count)
	
	if count > 0 {
		log.Println("Main admin already exists, skipping...")
		return nil
	}

	passwordHash, err := utils.HashPassword("admin123")
	if err != nil {
		return err
	}

	admin := models.MainAdmin{
		Username:     "admin",
		PasswordHash: passwordHash,
		IsActive:     true,
	}

	return database.DB.Create(&admin).Error
}

func seedCategories() error {
	var count int64
	database.DB.Model(&models.Category{}).Count(&count)
	
	if count > 0 {
		log.Println("Categories already exist, skipping...")
		return nil
	}

	categories := []models.Category{
		{
			Name:        "coffee",
			DisplayName: "قهوه",
			Emoji:       "☕",
			Color:       "from-amber-400 to-orange-500",
			OrderIndex:  1,
			IsActive:    true,
		},
		{
			Name:        "shake",
			DisplayName: "شیک",
			Emoji:       "🥤",
			Color:       "from-pink-400 to-rose-500",
			OrderIndex:  2,
			IsActive:    true,
		},
		{
			Name:        "cold_bar",
			DisplayName: "بار سرد",
			Emoji:       "🧊",
			Color:       "from-sky-400 to-blue-500",
			OrderIndex:  3,
			IsActive:    true,
		},
		{
			Name:        "hot_bar",
			DisplayName: "بار گرم",
			Emoji:       "🔥",
			Color:       "from-red-500 to-orange-500",
			OrderIndex:  4,
			IsActive:    true,
		},
		{
			Name:        "tea",
			DisplayName: "چای",
			Emoji:       "��",
			Color:       "from-lime-400 to-green-500",
			OrderIndex:  5,
			IsActive:    true,
		},
		{
			Name:        "cake",
			DisplayName: "کیک",
			Emoji:       "🍰",
			Color:       "from-fuchsia-500 to-pink-600",
			OrderIndex:  6,
			IsActive:    true,
		},
		{
			Name:        "food",
			DisplayName: "غذا",
			Emoji:       "🍽️",
			Color:       "from-indigo-400 to-purple-500",
			OrderIndex:  7,
			IsActive:    true,
		},
		{
			Name:        "breakfast",
			DisplayName: "صبحانه",
			Emoji:       "🌅",
			Color:       "from-yellow-400 to-amber-500",
			OrderIndex:  8,
			IsActive:    true,
		},
	}

	for _, category := range categories {
		if err := database.DB.Create(&category).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d categories", len(categories))
	return nil
}

func seedSampleTenant() error {
	var count int64
	database.DB.Model(&models.Tenant{}).Count(&count)
	
	if count > 0 {
		log.Println("Sample tenant already exists, skipping...")
		return nil
	}

	// Create sample tenant
	tenant := models.Tenant{
		Subdomain: "demo",
		Name:      "Demo Coffee Shop",
		IsActive:  true,
	}

	if err := database.DB.Create(&tenant).Error; err != nil {
		return err
	}

	// Create sample coffee shop
	coffeeShop := models.CoffeeShop{
		TenantID:     tenant.ID,
		Name:         "Demo Coffee Shop",
		Location:     "Tehran, Iran",
		Phone:        "+98-21-12345678",
		InstagramURL: "https://instagram.com/democoffee",
		LogoURL:      "https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?w=200",
		HeroImageURL: "https://images.unsplash.com/photo-1554118811-1e0d58224f24?w=800",
		Description:  "Best coffee in Tehran",
		IsActive:     true,
	}

	if err := database.DB.Create(&coffeeShop).Error; err != nil {
		return err
	}

	// Create shop admin
	passwordHash, err := utils.HashPassword("shop123")
	if err != nil {
		return err
	}

	admin := models.ShopAdmin{
		CoffeeShopID: coffeeShop.ID,
		Username:     "shopadmin",
		PasswordHash: passwordHash,
		IsActive:     true,
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		return err
	}

	// Get categories for menu items
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return err
	}

	// Create category map for easy lookup
	categoryMap := make(map[string]uint)
	for _, cat := range categories {
		categoryMap[cat.Name] = cat.ID
	}

	// Create sample menu items with proper category references
	menuItems := getSampleMenuItems(coffeeShop.ID, categoryMap)

	for _, item := range menuItems {
		if err := database.DB.Create(&item).Error; err != nil {
			return err
		}
	}

	log.Printf("Created sample tenant: %s (subdomain: %s)", tenant.Name, tenant.Subdomain)
	log.Printf("Created sample coffee shop: %s", coffeeShop.Name)
	log.Printf("Created shop admin: %s", admin.Username)
	log.Printf("Created %d menu items", len(menuItems))

	return nil
}

func getSampleMenuItems(coffeeShopID uint, categoryMap map[string]uint) []models.MenuItem {
	return []models.MenuItem{
		// Coffee Category - قهوه
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "اسپرسو لاین (قهوه 80/20 عربیکا)",
			Price:          45000,
			PricePremium:   &[]int{55000}[0],
			HasDualPricing: true,
			ImageURL:       "https://restcafe.storage.c2.liara.space/cafe/Screenshot%20from%202025-09-16%2015-20-33.png",
			OrderIndex:     1,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "آیس آمریکانو (قهوه 50/50 عربیکا)",
			Price:          35000,
			PricePremium:   &[]int{45000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1559056199-641a0ac8b55e?w=400",
			OrderIndex:     2,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "آمریکانو (قهوه 80/20 عربیکا)",
			Price:          30000,
			PricePremium:   &[]int{40000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1559056199-641a0ac8b55e?w=400",
			OrderIndex:     3,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "آفاگاتو (قهوه 50/50 عربیکا)",
			Price:          55000,
			PricePremium:   &[]int{65000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1572442388796-11668a67e53d?w=400",
			OrderIndex:     4,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "کاپوچینو (قهوه 80/20 عربیکا)",
			Price:          40000,
			PricePremium:   &[]int{50000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1572442388796-11668a67e53d?w=400",
			OrderIndex:     5,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "لته (قهوه 50/50 عربیکا)",
			Price:          42000,
			PricePremium:   &[]int{52000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:     6,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "موکا (قهوه 80/20 عربیکا)",
			Price:          48000,
			PricePremium:   &[]int{58000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:     7,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "ماکیاتو (قهوه 50/50 عربیکا)",
			Price:          38000,
			PricePremium:   &[]int{48000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1514432320407-a09c9e4aef1d?w=400",
			OrderIndex:     8,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "آیس لته (قهوه 80/20 عربیکا)",
			Price:          45000,
			PricePremium:   &[]int{55000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:     9,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "زومار (قهوه 50/50 عربیکا)",
			Price:          50000,
			PricePremium:   &[]int{60000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1514432320407-a09c9e4aef1d?w=400",
			OrderIndex:     10,
			IsAvailable:    true,
		},
		{
			CoffeeShopID:   coffeeShopID,
			CategoryID:     categoryMap["coffee"],
			Name:           "خیارپلو (قهوه 80/20 عربیکا)",
			Price:          52000,
			PricePremium:   &[]int{62000}[0],
			HasDualPricing: true,
			ImageURL:       "https://images.unsplash.com/photo-1514432320407-a09c9e4aef1d?w=400",
			OrderIndex:     11,
			IsAvailable:    true,
		},

		// Shake Category - شیک
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "نوتلا",
			Price:        65000,
			ImageURL:     "https://images.unsplash.com/photo-1572490122747-3968b75cc699?w=400",
			OrderIndex:   12,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "بادام",
			Price:        60000,
			ImageURL:     "https://images.unsplash.com/photo-1553530666-ba11a7da3888?w=400",
			OrderIndex:   13,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "لوتوس",
			Price:        70000,
			ImageURL:     "https://images.unsplash.com/photo-1572490122747-3968b75cc699?w=400",
			OrderIndex:   14,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "OREO",
			Price:        68000,
			ImageURL:     "https://images.unsplash.com/photo-1553530666-ba11a7da3888?w=400",
			OrderIndex:   15,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "نوستالژِ",
			Price:        55000,
			ImageURL:     "https://images.unsplash.com/photo-1572490122747-3968b75cc699?w=400",
			OrderIndex:   16,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "بری",
			Price:        58000,
			ImageURL:     "https://images.unsplash.com/photo-1553530666-ba11a7da3888?w=400",
			OrderIndex:   17,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "شکلات",
			Price:        62000,
			ImageURL:     "https://images.unsplash.com/photo-1572490122747-3968b75cc699?w=400",
			OrderIndex:   18,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["shake"],
			Name:         "قهوه",
			Price:        50000,
			ImageURL:     "https://images.unsplash.com/photo-1553530666-ba11a7da3888?w=400",
			OrderIndex:   19,
			IsAvailable:  true,
		},

		// Cold Bar Category - بار سرد
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "ردگاردن",
			Price:        45000,
			ImageURL:     "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:   20,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "لیموناد نعناع",
			Price:        40000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   21,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "فروزن لایت",
			Price:        35000,
			ImageURL:     "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:   22,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "مانگوپشن",
			Price:        48000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   23,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "آب نبات",
			Price:        42000,
			ImageURL:     "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:   24,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "موهیتو",
			Price:        50000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   25,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cold_bar"],
			Name:         "ترش",
			Price:        38000,
			ImageURL:     "https://images.unsplash.com/photo-1578314675249-a6910f80cc4e?w=400",
			OrderIndex:   26,
			IsAvailable:  true,
		},

		// Hot Bar Category - بار گرم
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["hot_bar"],
			Name:         "هات چاکلت",
			Price:        55000,
			ImageURL:     "https://images.unsplash.com/photo-1542990253-0d0f5be5f0ed?w=400",
			OrderIndex:   27,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["hot_bar"],
			Name:         "یونانی",
			Price:        45000,
			ImageURL:     "https://images.unsplash.com/photo-1542990253-0d0f5be5f0ed?w=400",
			OrderIndex:   28,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["hot_bar"],
			Name:         "شیرشکلات",
			Price:        50000,
			ImageURL:     "https://images.unsplash.com/photo-1542990253-0d0f5be5f0ed?w=400",
			OrderIndex:   29,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["hot_bar"],
			Name:         "شیرنسکافه",
			Price:        48000,
			ImageURL:     "https://images.unsplash.com/photo-1542990253-0d0f5be5f0ed?w=400",
			OrderIndex:   30,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["hot_bar"],
			Name:         "شیرکاکائو",
			Price:        52000,
			ImageURL:     "https://images.unsplash.com/photo-1542990253-0d0f5be5f0ed?w=400",
			OrderIndex:   31,
			IsAvailable:  true,
		},

		// Tea Category - چای
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["tea"],
			Name:         "دمنوش",
			Price:        35000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   32,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["tea"],
			Name:         "ساده",
			Price:        25000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   33,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["tea"],
			Name:         "ماسالا",
			Price:        40000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   34,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["tea"],
			Name:         "ماچا",
			Price:        45000,
			ImageURL:     "https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=400",
			OrderIndex:   35,
			IsAvailable:  true,
		},

		// Cake Category - کیک
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cake"],
			Name:         "چیز کیک",
			Price:        85000,
			ImageURL:     "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400",
			OrderIndex:   36,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cake"],
			Name:         "دبل چاکلت",
			Price:        95000,
			ImageURL:     "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400",
			OrderIndex:   37,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cake"],
			Name:         "فرانسوی",
			Price:        90000,
			ImageURL:     "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400",
			OrderIndex:   38,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cake"],
			Name:         "هویج",
			Price:        75000,
			ImageURL:     "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400",
			OrderIndex:   39,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["cake"],
			Name:         "پای سیب",
			Price:        80000,
			ImageURL:     "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400",
			OrderIndex:   40,
			IsAvailable:  true,
		},

		// Food Category - غذا
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["food"],
			Name:         "پاستا",
			Price:        120000,
			ImageURL:     "https://images.unsplash.com/photo-1621996346565-e3dbc353d2e5?w=400",
			OrderIndex:   41,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["food"],
			Name:         "سیب زمینی با سس مخصوص",
			Price:        65000,
			ImageURL:     "https://images.unsplash.com/photo-1528735602786-469f3817357d?w=400",
			OrderIndex:   42,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["food"],
			Name:         "سالاد سزار",
			Price:        85000,
			ImageURL:     "https://images.unsplash.com/photo-1528735602786-469f3817357d?w=400",
			OrderIndex:   43,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["food"],
			Name:         "سالاد ویژه رست",
			Price:        95000,
			ImageURL:     "https://images.unsplash.com/photo-1528735602786-469f3817357d?w=400",
			OrderIndex:   44,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["food"],
			Name:         "سالاد ماکارونی",
			Price:        70000,
			ImageURL:     "https://images.unsplash.com/photo-1621996346565-e3dbc353d2e5?w=400",
			OrderIndex:   45,
			IsAvailable:  true,
		},

		// Breakfast Category - صبحانه
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["breakfast"],
			Name:         "صبحانه ایرانی",
			Price:        150000,
			ImageURL:     "https://images.unsplash.com/photo-1482049016688-2d3e1b311543?w=400",
			OrderIndex:   46,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["breakfast"],
			Name:         "املت",
			Price:        75000,
			ImageURL:     "https://images.unsplash.com/photo-1482049016688-2d3e1b311543?w=400",
			OrderIndex:   47,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["breakfast"],
			Name:         "املت سوجوک",
			Price:        95000,
			ImageURL:     "https://images.unsplash.com/photo-1482049016688-2d3e1b311543?w=400",
			OrderIndex:   48,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["breakfast"],
			Name:         "صبحانه انگلیسی",
			Price:        180000,
			ImageURL:     "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=400",
			OrderIndex:   49,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["breakfast"],
			Name:         "خوراک عدسی",
			Price:        65000,
			ImageURL:     "https://images.unsplash.com/photo-1482049016688-2d3e1b311543?w=400",
			OrderIndex:   50,
			IsAvailable:  true,
		},
		{
			CoffeeShopID: coffeeShopID,
			CategoryID:   categoryMap["breakfast"],
			Name:         "نیمرو",
			Price:        60000,
			ImageURL:     "https://images.unsplash.com/photo-1482049016688-2d3e1b311543?w=400",
			OrderIndex:   51,
			IsAvailable:  true,
		},
	}
}
