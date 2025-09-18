import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { createPageUrl } from "../utils";
import { 
  Coffee, 
  Plus, 
  Edit, 
  Trash2, 
  Eye, 
  EyeOff, 
  Settings,
  LogOut,
  Menu,
  Package,
  Users,
  BarChart3
} from "lucide-react";
import { Button } from "../components/ui";
import { Card, CardContent, CardHeader, CardTitle } from "../components/ui";
import { motion, AnimatePresence } from "framer-motion";
import { MenuItem, Category, CafeSettings } from "../entities";
import apiService from "../services/api";

export default function AdminDashboardPage() {
  const navigate = useNavigate();
  const [menuItems, setMenuItems] = useState([]);
  const [categories, setCategories] = useState([]);
  const [cafeSettings, setCafeSettings] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [activeTab, setActiveTab] = useState("menu");
  const [showAddForm, setShowAddForm] = useState(false);
  const [editingItem, setEditingItem] = useState(null);
  const [userType, setUserType] = useState("");

  // Form state
  const [formData, setFormData] = useState({
    name: "",
    category_id: "",
    price: "",
    price_premium: "",
    has_dual_pricing: false,
    image_url: "",
    order_index: 0,
    is_available: true
  });

  useEffect(() => {
    // Check if user is logged in
    const token = localStorage.getItem('auth_token');
    const userType = localStorage.getItem('user_type');
    
    if (!token || !userType) {
      navigate(createPageUrl("AdminLogin"));
      return;
    }

    setUserType(userType);
    loadData();
  }, [navigate]);

  const loadData = async () => {
    setLoading(true);
    setError("");
    
    try {
      console.log("=== AdminDashboard.loadData() START ===");
      
      const [menuData, categoriesData, settingsData] = await Promise.all([
        MenuItem.list(),
        Category.getCategories(),
        CafeSettings.getSettings()
      ]);

      console.log("Menu items loaded:", menuData.length);
      console.log("Categories loaded:", categoriesData.length);
      console.log("Settings loaded:", settingsData);

      setMenuItems(menuData);
      setCategories(categoriesData);
      setCafeSettings(settingsData);
      
      console.log("=== AdminDashboard.loadData() SUCCESS ===");
    } catch (error) {
      console.error("Failed to load data:", error);
      setError(error.message || "خطا در بارگذاری اطلاعات");
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");

    try {
      const data = {
        ...formData,
        price: parseInt(formData.price),
        price_premium: formData.has_dual_pricing ? parseInt(formData.price_premium) : null,
        category_id: parseInt(formData.category_id)
      };

      if (editingItem) {
        await editingItem.update(data);
        setMenuItems(prev => prev.map(item => 
          item.id === editingItem.id ? { ...item, ...data } : item
        ));
      } else {
        const newItem = await MenuItem.create(data);
        setMenuItems(prev => [...prev, newItem]);
      }

      resetForm();
    } catch (error) {
      console.error("Failed to save menu item:", error);
      setError(error.message || "خطا در ذخیره آیتم");
    } finally {
      setLoading(false);
    }
  };

  const handleEdit = (item) => {
    setEditingItem(item);
    setFormData({
      name: item.name,
      category_id: item.category_id.toString(),
      price: item.price.toString(),
      price_premium: item.price_premium ? item.price_premium.toString() : "",
      has_dual_pricing: item.has_dual_pricing,
      image_url: item.image_url,
      order_index: item.order_index,
      is_available: item.is_available
    });
    setShowAddForm(true);
  };

  const handleDelete = async (item) => {
    if (!window.confirm("آیا از حذف این آیتم اطمینان دارید؟")) return;

    try {
      await item.delete();
      setMenuItems(prev => prev.filter(i => i.id !== item.id));
    } catch (error) {
      console.error("Failed to delete menu item:", error);
      setError(error.message || "خطا در حذف آیتم");
    }
  };

  const resetForm = () => {
    setFormData({
      name: "",
      category_id: "",
      price: "",
      price_premium: "",
      has_dual_pricing: false,
      image_url: "",
      order_index: 0,
      is_available: true
    });
    setEditingItem(null);
    setShowAddForm(false);
  };

  const handleLogout = () => {
    apiService.setToken(null);
    localStorage.removeItem('user_type');
    localStorage.removeItem('user_info');
    navigate(createPageUrl("Menu"));
  };

  const groupedItems = categories.reduce((acc, category) => {
    const items = menuItems.filter(item => item.category_id === category.id);
    if (items.length > 0) {
      acc.push({
        ...category,
        items: items
      });
    }
    return acc;
  }, []);

  if (loading && menuItems.length === 0) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-amber-50 to-orange-100 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-amber-600 mx-auto mb-4"></div>
          <p className="text-amber-700 text-lg">در حال بارگذاری...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-amber-50 to-orange-100">
      {/* Header */}
      <header className="bg-white shadow-sm border-b border-amber-200">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <Coffee className="w-8 h-8 text-amber-600" />
              <div>
                <h1 className="text-2xl font-bold text-amber-800">پنل مدیریت</h1>
                <p className="text-amber-600 text-sm">
                  {cafeSettings.name || "کافه رست"}
                </p>
              </div>
            </div>
            <div className="flex items-center space-x-2">
              <Button
                onClick={() => navigate(createPageUrl("Menu"))}
                variant="outline"
                className="text-amber-600 border-amber-300 hover:bg-amber-50"
              >
                <Eye className="w-4 h-4 mr-2" />
                مشاهده منو
              </Button>
              <Button
                onClick={handleLogout}
                variant="outline"
                className="text-red-600 border-red-300 hover:bg-red-50"
              >
                <LogOut className="w-4 h-4 mr-2" />
                خروج
              </Button>
            </div>
          </div>
        </div>
      </header>

      <div className="container mx-auto px-4 py-6">
        {/* Tabs */}
        <div className="flex space-x-1 mb-6 bg-white rounded-lg p-1 shadow-sm">
          <button
            onClick={() => setActiveTab("menu")}
            className={`flex-1 py-2 px-4 rounded-md font-medium transition-colors ${
              activeTab === "menu"
                ? "bg-amber-600 text-white"
                : "text-amber-600 hover:bg-amber-50"
            }`}
          >
            <Menu className="w-4 h-4 inline mr-2" />
            مدیریت منو
          </button>
          {userType === "main" && (
            <button
              onClick={() => setActiveTab("categories")}
              className={`flex-1 py-2 px-4 rounded-md font-medium transition-colors ${
                activeTab === "categories"
                  ? "bg-amber-600 text-white"
                  : "text-amber-600 hover:bg-amber-50"
              }`}
            >
              <Package className="w-4 h-4 inline mr-2" />
              مدیریت دسته‌بندی
            </button>
          )}
          <button
            onClick={() => setActiveTab("settings")}
            className={`flex-1 py-2 px-4 rounded-md font-medium transition-colors ${
              activeTab === "settings"
                ? "bg-amber-600 text-white"
                : "text-amber-600 hover:bg-amber-50"
            }`}
          >
            <Settings className="w-4 h-4 inline mr-2" />
            تنظیمات
          </button>
        </div>

        {/* Error Message */}
        {error && (
          <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-6">
            {error}
          </div>
        )}

        {/* Menu Management Tab */}
        {activeTab === "menu" && (
          <div className="space-y-6">
            {/* Add/Edit Form */}
            <Card>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle>
                    {editingItem ? "ویرایش آیتم" : "افزودن آیتم جدید"}
                  </CardTitle>
                  <Button
                    onClick={() => setShowAddForm(!showAddForm)}
                    variant="outline"
                    size="sm"
                  >
                    {showAddForm ? "بستن" : "افزودن آیتم"}
                  </Button>
                </div>
              </CardHeader>
              {showAddForm && (
                <CardContent>
                  <form onSubmit={handleSubmit} className="space-y-4">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                      <div>
                        <label className="block text-sm font-medium text-amber-700 mb-1">
                          نام آیتم
                        </label>
                        <input
                          type="text"
                          name="name"
                          value={formData.name}
                          onChange={handleInputChange}
                          required
                          className="w-full px-3 py-2 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent"
                          placeholder="نام آیتم را وارد کنید"
                        />
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-amber-700 mb-1">
                          دسته‌بندی
                        </label>
                        <select
                          name="category_id"
                          value={formData.category_id}
                          onChange={handleInputChange}
                          required
                          className="w-full px-3 py-2 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent"
                        >
                          <option value="">انتخاب دسته‌بندی</option>
                          {categories.map(category => (
                            <option key={category.id} value={category.id}>
                              {category.emoji} {category.display_name}
                            </option>
                          ))}
                        </select>
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-amber-700 mb-1">
                          قیمت (تومان)
                        </label>
                        <input
                          type="number"
                          name="price"
                          value={formData.price}
                          onChange={handleInputChange}
                          required
                          min="0"
                          className="w-full px-3 py-2 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent"
                          placeholder="قیمت را وارد کنید"
                        />
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-amber-700 mb-1">
                          قیمت پریمیوم (تومان)
                        </label>
                        <input
                          type="number"
                          name="price_premium"
                          value={formData.price_premium}
                          onChange={handleInputChange}
                          min="0"
                          disabled={!formData.has_dual_pricing}
                          className="w-full px-3 py-2 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent disabled:bg-gray-100"
                          placeholder="قیمت پریمیوم را وارد کنید"
                        />
                      </div>
                      <div className="flex items-center">
                        <input
                          type="checkbox"
                          name="has_dual_pricing"
                          checked={formData.has_dual_pricing}
                          onChange={handleInputChange}
                          className="mr-2 text-amber-600"
                        />
                        <label className="text-sm font-medium text-amber-700">
                          قیمت دوگانه
                        </label>
                      </div>
                      <div className="flex items-center">
                        <input
                          type="checkbox"
                          name="is_available"
                          checked={formData.is_available}
                          onChange={handleInputChange}
                          className="mr-2 text-amber-600"
                        />
                        <label className="text-sm font-medium text-amber-700">
                          موجود است
                        </label>
                      </div>
                    </div>
                    <div>
                      <label className="block text-sm font-medium text-amber-700 mb-1">
                        آدرس تصویر
                      </label>
                      <input
                        type="url"
                        name="image_url"
                        value={formData.image_url}
                        onChange={handleInputChange}
                        className="w-full px-3 py-2 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent"
                        placeholder="آدرس تصویر را وارد کنید"
                      />
                    </div>
                    <div className="flex space-x-2">
                      <Button
                        type="submit"
                        disabled={loading}
                        className="bg-amber-600 hover:bg-amber-700 text-white"
                      >
                        {loading ? "در حال ذخیره..." : (editingItem ? "ویرایش" : "افزودن")}
                      </Button>
                      <Button
                        type="button"
                        onClick={resetForm}
                        variant="outline"
                      >
                        لغو
                      </Button>
                    </div>
                  </form>
                </CardContent>
              )}
            </Card>

            {/* Menu Items List */}
            <div className="space-y-6">
              {groupedItems.map((category) => (
                <Card key={category.id}>
                  <CardHeader>
                    <CardTitle className="flex items-center">
                      {category.emoji} {category.display_name}
                      <span className="mr-2 text-sm font-normal text-amber-600">
                        ({category.items.length} آیتم)
                      </span>
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                      {category.items.map((item) => (
                        <div
                          key={item.id}
                          className="border border-amber-200 rounded-lg p-4 hover:shadow-md transition-shadow"
                        >
                          <div className="flex items-start justify-between mb-2">
                            <h3 className="font-medium text-amber-800">{item.name}</h3>
                            <div className="flex space-x-1">
                              <Button
                                onClick={() => handleEdit(item)}
                                size="sm"
                                variant="outline"
                                className="text-amber-600 border-amber-300 hover:bg-amber-50"
                              >
                                <Edit className="w-3 h-3" />
                              </Button>
                              <Button
                                onClick={() => handleDelete(item)}
                                size="sm"
                                variant="outline"
                                className="text-red-600 border-red-300 hover:bg-red-50"
                              >
                                <Trash2 className="w-3 h-3" />
                              </Button>
                            </div>
                          </div>
                          <div className="text-sm text-amber-600">
                            {item.getFormattedPrice()}
                          </div>
                          <div className="flex items-center mt-2">
                            <span className={`text-xs px-2 py-1 rounded-full ${
                              item.is_available 
                                ? "bg-green-100 text-green-700" 
                                : "bg-red-100 text-red-700"
                            }`}>
                              {item.is_available ? "موجود" : "ناموجود"}
                            </span>
                          </div>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          </div>
        )}

        {/* Categories Management Tab (Main Admin Only) */}
        {activeTab === "categories" && userType === "main" && (
          <div className="space-y-6">
            <Card>
              <CardHeader>
                <CardTitle>مدیریت دسته‌بندی‌ها</CardTitle>
                <p className="text-amber-600 text-sm">
                  دسته‌بندی‌ها توسط ادمین اصلی مدیریت می‌شوند
                </p>
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  {categories.map((category) => (
                    <div
                      key={category.id}
                      className="border border-amber-200 rounded-lg p-4 hover:shadow-md transition-shadow"
                    >
                      <div className="flex items-center mb-2">
                        <span className="text-2xl mr-2">{category.emoji}</span>
                        <h3 className="font-medium text-amber-800">{category.display_name}</h3>
                      </div>
                      <p className="text-sm text-amber-600">{category.name}</p>
                      <div className="flex items-center mt-2">
                        <span className={`text-xs px-2 py-1 rounded-full ${
                          category.is_active 
                            ? "bg-green-100 text-green-700" 
                            : "bg-red-100 text-red-700"
                        }`}>
                          {category.is_active ? "فعال" : "غیرفعال"}
                        </span>
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </div>
        )}

        {/* Settings Tab */}
        {activeTab === "settings" && (
          <div className="space-y-6">
            <Card>
              <CardHeader>
                <CardTitle>تنظیمات فروشگاه</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-amber-700 mb-1">
                      نام فروشگاه
                    </label>
                    <input
                      type="text"
                      value={cafeSettings.name || ""}
                      readOnly
                      className="w-full px-3 py-2 border border-amber-200 rounded-lg bg-gray-50"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-amber-700 mb-1">
                      موقعیت
                    </label>
                    <input
                      type="text"
                      value={cafeSettings.location || ""}
                      readOnly
                      className="w-full px-3 py-2 border border-amber-200 rounded-lg bg-gray-50"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-amber-700 mb-1">
                      تلفن
                    </label>
                    <input
                      type="text"
                      value={cafeSettings.phone || ""}
                      readOnly
                      className="w-full px-3 py-2 border border-amber-200 rounded-lg bg-gray-50"
                    />
                  </div>
                  <p className="text-sm text-amber-600">
                    برای تغییر تنظیمات، با ادمین اصلی تماس بگیرید.
                  </p>
                </div>
              </CardContent>
            </Card>
          </div>
        )}
      </div>
    </div>
  );
}
