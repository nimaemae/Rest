import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createPageUrl } from "../utils";
import { Coffee, Lock, User, Eye, EyeOff } from "lucide-react";
import { Button } from "../components/ui";
import { Card, CardContent, CardHeader, CardTitle } from "../components/ui";
import { motion } from "framer-motion";
import apiService from "../services/api";

export default function AdminLoginPage() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    username: "",
    password: "",
    adminType: "shop" // "main" or "shop"
  });
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
    // Clear error when user starts typing
    if (error) setError("");
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");

    try {
      console.log("=== AdminLogin.handleSubmit() START ===");
      console.log("Admin type:", formData.adminType);
      console.log("Username:", formData.username);

      let response;
      if (formData.adminType === "main") {
        response = await apiService.loginMainAdmin({
          username: formData.username,
          password: formData.password
        });
      } else {
        response = await apiService.loginShopAdmin({
          username: formData.username,
          password: formData.password
        });
      }

      console.log("Login response:", response);

      // Store token and user info
      apiService.setToken(response.token);
      localStorage.setItem('user_type', formData.adminType);
      localStorage.setItem('user_info', JSON.stringify(response.user));

      console.log("=== AdminLogin.handleSubmit() SUCCESS ===");

      // Navigate to dashboard
      navigate(createPageUrl("AdminDashboard"));
    } catch (error) {
      console.error("Login failed:", error);
      setError(error.message || "خطا در ورود. لطفاً اطلاعات را بررسی کنید.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-amber-50 to-orange-100 flex items-center justify-center p-4">
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        className="w-full max-w-md"
      >
        <Card className="shadow-2xl border-0">
          <CardHeader className="text-center pb-8">
            <div className="mx-auto w-16 h-16 bg-amber-100 rounded-full flex items-center justify-center mb-4">
              <Coffee className="w-8 h-8 text-amber-600" />
            </div>
            <CardTitle className="text-2xl font-bold text-amber-800">
              ورود به پنل مدیریت
            </CardTitle>
            <p className="text-amber-600 mt-2">
              لطفاً اطلاعات ورود خود را وارد کنید
            </p>
          </CardHeader>

          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              {/* Admin Type Selection */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-amber-700">
                  نوع ادمین
                </label>
                <div className="flex space-x-4">
                  <label className="flex items-center">
                    <input
                      type="radio"
                      name="adminType"
                      value="shop"
                      checked={formData.adminType === "shop"}
                      onChange={handleInputChange}
                      className="mr-2 text-amber-600"
                    />
                    <span className="text-sm text-amber-700">ادمین فروشگاه</span>
                  </label>
                  <label className="flex items-center">
                    <input
                      type="radio"
                      name="adminType"
                      value="main"
                      checked={formData.adminType === "main"}
                      onChange={handleInputChange}
                      className="mr-2 text-amber-600"
                    />
                    <span className="text-sm text-amber-700">ادمین اصلی</span>
                  </label>
                </div>
              </div>

              {/* Username */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-amber-700">
                  نام کاربری
                </label>
                <div className="relative">
                  <User className="absolute right-3 top-1/2 transform -translate-y-1/2 text-amber-400 w-5 h-5" />
                  <input
                    type="text"
                    name="username"
                    value={formData.username}
                    onChange={handleInputChange}
                    required
                    className="w-full pr-10 pl-4 py-3 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent text-right"
                    placeholder="نام کاربری خود را وارد کنید"
                  />
                </div>
              </div>

              {/* Password */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-amber-700">
                  رمز عبور
                </label>
                <div className="relative">
                  <Lock className="absolute right-3 top-1/2 transform -translate-y-1/2 text-amber-400 w-5 h-5" />
                  <input
                    type={showPassword ? "text" : "password"}
                    name="password"
                    value={formData.password}
                    onChange={handleInputChange}
                    required
                    className="w-full pr-10 pl-12 py-3 border border-amber-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-transparent text-right"
                    placeholder="رمز عبور خود را وارد کنید"
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute left-3 top-1/2 transform -translate-y-1/2 text-amber-400 hover:text-amber-600"
                  >
                    {showPassword ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
                  </button>
                </div>
              </div>

              {/* Error Message */}
              {error && (
                <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm">
                  {error}
                </div>
              )}

              {/* Submit Button */}
              <Button
                type="submit"
                disabled={loading}
                className="w-full bg-amber-600 hover:bg-amber-700 text-white py-3 font-medium disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {loading ? (
                  <div className="flex items-center justify-center">
                    <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white mr-2"></div>
                    در حال ورود...
                  </div>
                ) : (
                  "ورود"
                )}
              </Button>
            </form>

            {/* Back to Menu */}
            <div className="mt-6 text-center">
              <button
                onClick={() => navigate(createPageUrl("Menu"))}
                className="text-amber-600 hover:text-amber-700 text-sm font-medium"
              >
                ← بازگشت به منو
              </button>
            </div>
          </CardContent>
        </Card>
      </motion.div>
    </div>
  );
}
