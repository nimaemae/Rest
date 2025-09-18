import apiService from '../services/api';

export class MenuItem {
  constructor(data = {}) {
    this.id = data.id;
    this.name = data.name || '';
    this.category_id = data.category_id;
    this.category = data.category || {};
    this.price = data.price || 0;
    this.price_premium = data.price_premium || null;
    this.has_dual_pricing = data.has_dual_pricing || false;
    this.image_url = data.image_url || '';
    this.order_index = data.order_index || 0;
    this.is_available = data.is_available !== undefined ? data.is_available : true;
    this.created_at = data.created_at;
    this.updated_at = data.updated_at;
  }

  // Static methods for API operations
  static async list() {
    try {
      const response = await apiService.getMenuItems();
      return response.map(item => new MenuItem(item));
    } catch (error) {
      console.error('Failed to fetch menu items:', error);
      throw error;
    }
  }

  static async getPublicMenu() {
    try {
      const response = await apiService.getPublicMenu();
      return response.map(item => new MenuItem(item));
    } catch (error) {
      console.error('Failed to fetch public menu:', error);
      throw error;
    }
  }

  static async getById(id) {
    try {
      const response = await apiService.getMenuItem(id);
      return new MenuItem(response);
    } catch (error) {
      console.error('Failed to fetch menu item:', error);
      throw error;
    }
  }

  static async create(data) {
    try {
      const response = await apiService.createMenuItem(data);
      return new MenuItem(response.data || response);
    } catch (error) {
      console.error('Failed to create menu item:', error);
      throw error;
    }
  }

  async update(data) {
    try {
      const response = await apiService.updateMenuItem(this.id, data);
      const updatedData = response.data || response;
      Object.assign(this, updatedData);
      return this;
    } catch (error) {
      console.error('Failed to update menu item:', error);
      throw error;
    }
  }

  async delete() {
    try {
      await apiService.deleteMenuItem(this.id);
      return true;
    } catch (error) {
      console.error('Failed to delete menu item:', error);
      throw error;
    }
  }

  // Helper methods
  getDisplayPrice() {
    if (this.has_dual_pricing && this.price_premium) {
      return {
        regular: this.price,
        premium: this.price_premium
      };
    }
    return {
      regular: this.price,
      premium: null
    };
  }

  getFormattedPrice() {
    const prices = this.getDisplayPrice();
    if (prices.premium) {
      return `${this.formatPrice(prices.regular)} - ${this.formatPrice(prices.premium)}`;
    }
    return this.formatPrice(prices.regular);
  }

  formatPrice(price) {
    return new Intl.NumberFormat('fa-IR', {
      style: 'currency',
      currency: 'IRR',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(price);
  }

  // Legacy methods for backward compatibility
  static async seed() {
    console.log('MenuItem.seed() is deprecated. Use the backend seeding instead.');
    return [];
  }

  static async save() {
    console.log('MenuItem.save() is deprecated. Use create() or update() instead.');
  }
}
