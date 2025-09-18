// API Configuration
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

class ApiService {
  constructor() {
    this.baseURL = API_BASE_URL;
    this.token = localStorage.getItem('auth_token');
  }

  setToken(token) {
    this.token = token;
    if (token) {
      localStorage.setItem('auth_token', token);
    } else {
      localStorage.removeItem('auth_token');
    }
  }

  getHeaders() {
    const headers = {
      'Content-Type': 'application/json',
    };
    
    if (this.token) {
      headers['Authorization'] = `Bearer ${this.token}`;
    }
    
    return headers;
  }

  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`;
    const config = {
      headers: this.getHeaders(),
      ...options,
    };

    try {
      const response = await fetch(url, config);
      
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error('API request failed:', error);
      throw error;
    }
  }

  // Auth endpoints
  async loginMainAdmin(credentials) {
    return this.request('/api/auth/main-admin/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    });
  }

  async loginShopAdmin(credentials) {
    return this.request('/api/auth/shop-admin/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    });
  }

  // Public endpoints
  async getPublicCategories() {
    return this.request('/api/public/categories');
  }

  async getPublicMenu() {
    return this.request('/api/public/menu');
  }

  async getPublicShopSettings() {
    return this.request('/api/public/shop');
  }

  // Main admin endpoints
  async getTenants() {
    return this.request('/api/admin/tenants');
  }

  async createTenant(data) {
    return this.request('/api/admin/tenants', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getTenant(id) {
    return this.request(`/api/admin/tenants/${id}`);
  }

  async updateTenant(id, data) {
    return this.request(`/api/admin/tenants/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteTenant(id) {
    return this.request(`/api/admin/tenants/${id}`, {
      method: 'DELETE',
    });
  }

  async getCoffeeShops(tenantId) {
    return this.request(`/api/admin/tenants/${tenantId}/shops`);
  }

  async createCoffeeShop(tenantId, data) {
    return this.request(`/api/admin/tenants/${tenantId}/shops`, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getCoffeeShop(id) {
    return this.request(`/api/admin/shops/${id}`);
  }

  async updateCoffeeShop(id, data) {
    return this.request(`/api/admin/shops/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteCoffeeShop(id) {
    return this.request(`/api/admin/shops/${id}`, {
      method: 'DELETE',
    });
  }

  async createShopAdmin(shopId, data) {
    return this.request(`/api/admin/shops/${shopId}/admins`, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Category management (main admin only)
  async getAllCategories() {
    return this.request('/api/admin/categories');
  }

  async createCategory(data) {
    return this.request('/api/admin/categories', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getCategory(id) {
    return this.request(`/api/admin/categories/${id}`);
  }

  async updateCategory(id, data) {
    return this.request(`/api/admin/categories/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteCategory(id) {
    return this.request(`/api/admin/categories/${id}`, {
      method: 'DELETE',
    });
  }

  // Shop admin endpoints
  async getCategories() {
    return this.request('/api/admin/categories');
  }

  async getMenuItems() {
    return this.request('/api/admin/menu');
  }

  async createMenuItem(data) {
    return this.request('/api/admin/menu', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getMenuItem(id) {
    return this.request(`/api/admin/menu/${id}`);
  }

  async updateMenuItem(id, data) {
    return this.request(`/api/admin/menu/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteMenuItem(id) {
    return this.request(`/api/admin/menu/${id}`, {
      method: 'DELETE',
    });
  }

  async getShopSettings() {
    return this.request('/api/admin/settings');
  }

  async updateShopSettings(data) {
    return this.request('/api/admin/settings', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }
}

// Create and export a singleton instance
export const apiService = new ApiService();
export default apiService;
