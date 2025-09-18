import React, { useState, useEffect } from "react";
import { MenuItem, CafeSettings, Category } from "../entities";
import { Link } from "react-router-dom";
import { createPageUrl, formatPrice } from "../utils";
import { Coffee, MapPin, Instagram, Phone, Settings, Star } from "lucide-react";
import { Card, CardContent } from "../components/ui";
import { Badge } from "../components/ui";
import { Button } from "../components/ui";
import { motion, AnimatePresence } from "framer-motion";

import { MenuItemCard, CategorySection, CafeHeader } from "../components/menu";

export default function MenuPage() {
  const [menuItems, setMenuItems] = useState([]);
  const [categories, setCategories] = useState([]);
  const [cafeSettings, setCafeSettings] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [selectedCategory, setSelectedCategory] = useState("all");

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    setLoading(true);
    setError(null);
    try {
      console.log("=== Menu.loadData() START ===");
      
      // Load categories and menu items in parallel
      const [categoriesData, menuData, settingsData] = await Promise.all([
        Category.getPublicCategories(),
        MenuItem.getPublicMenu(),
        CafeSettings.getPublicSettings()
      ]);

      console.log("Categories loaded:", categoriesData.length);
      console.log("Menu items loaded:", menuData.length);
      console.log("Settings loaded:", settingsData);

      setCategories(categoriesData);
      setMenuItems(menuData);
      setCafeSettings(settingsData);
      
      console.log("=== Menu.loadData() SUCCESS ===");
    } catch (error) {
      console.error("Failed to load data:", error);
      setError(error.message || "Failed to load menu data");
    } finally {
      setLoading(false);
    }
  };

  const filteredItems = selectedCategory === "all" 
    ? menuItems 
    : menuItems.filter(item => item.category_id === selectedCategory);

  const groupedItems = categories.reduce((acc, category) => {
    const items = filteredItems.filter(item => item.category_id === category.id);
    if (items.length > 0) {
      acc.push({
        ...category,
        items: items
      });
    }
    return acc;
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-amber-50 to-orange-100 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-amber-600 mx-auto mb-4"></div>
          <p className="text-amber-700 text-lg">در حال بارگذاری منو...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-amber-50 to-orange-100 flex items-center justify-center">
        <div className="text-center">
          <div className="text-red-500 text-6xl mb-4">⚠️</div>
          <h2 className="text-2xl font-bold text-red-700 mb-2">خطا در بارگذاری</h2>
          <p className="text-red-600 mb-4">{error}</p>
          <Button 
            onClick={loadData}
            className="bg-red-600 hover:bg-red-700 text-white"
          >
            تلاش مجدد
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-amber-50 to-orange-100">
      {/* Header */}
      <CafeHeader 
        settings={cafeSettings}
        onSettingsClick={() => window.location.href = createPageUrl("AdminLogin")}
      />

      {/* Category Filter */}
      <div className="container mx-auto px-4 py-6">
        <div className="flex flex-wrap gap-2 justify-center mb-8">
          <button
            onClick={() => setSelectedCategory("all")}
            className={`px-4 py-2 rounded-full font-medium transition-all duration-200 ${
              selectedCategory === "all"
                ? "bg-amber-600 text-white shadow-lg"
                : "bg-white text-amber-700 hover:bg-amber-50"
            }`}
          >
            همه
          </button>
          {categories.map((category) => (
            <button
              key={category.id}
              onClick={() => setSelectedCategory(category.id)}
              className={`px-4 py-2 rounded-full font-medium transition-all duration-200 ${
                selectedCategory === category.id
                  ? "bg-amber-600 text-white shadow-lg"
                  : "bg-white text-amber-700 hover:bg-amber-50"
              }`}
            >
              {category.emoji} {category.display_name}
            </button>
          ))}
        </div>

        {/* Menu Items */}
        <AnimatePresence>
          {groupedItems.length > 0 ? (
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -20 }}
              className="space-y-8"
            >
              {groupedItems.map((category) => (
                <CategorySection
                  key={category.id}
                  category={category}
                  items={category.items}
                />
              ))}
            </motion.div>
          ) : (
            <div className="text-center py-12">
              <div className="text-6xl mb-4">🍽️</div>
              <h3 className="text-2xl font-bold text-amber-700 mb-2">
                آیتمی یافت نشد
              </h3>
              <p className="text-amber-600">
                {selectedCategory === "all" 
                  ? "در حال حاضر آیتمی در منو موجود نیست"
                  : "در این دسته‌بندی آیتمی موجود نیست"
                }
              </p>
            </div>
          )}
        </AnimatePresence>
      </div>

      {/* Footer */}
      <footer className="bg-amber-800 text-white py-8 mt-12">
        <div className="container mx-auto px-4 text-center">
          <div className="flex justify-center space-x-6 mb-4">
            {cafeSettings.instagram_url && (
              <a
                href={cafeSettings.instagram_url}
                target="_blank"
                rel="noopener noreferrer"
                className="hover:text-amber-300 transition-colors"
              >
                <Instagram className="w-6 h-6" />
              </a>
            )}
            {cafeSettings.phone && (
              <a
                href={`tel:${cafeSettings.phone}`}
                className="hover:text-amber-300 transition-colors"
              >
                <Phone className="w-6 h-6" />
              </a>
            )}
          </div>
          <p className="text-amber-200">
            © 2024 {cafeSettings.name || "کافه رست"}. تمامی حقوق محفوظ است.
          </p>
        </div>
      </footer>
    </div>
  );
}
