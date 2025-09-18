import apiService from '../services/api';

export class Category {
  constructor(data = {}) {
    this.id = data.id;
    this.name = data.name || '';
    this.display_name = data.display_name || '';
    this.emoji = data.emoji || '';
    this.color = data.color || '';
    this.order_index = data.order_index || 0;
    this.is_active = data.is_active !== undefined ? data.is_active : true;
    this.created_at = data.created_at;
    this.updated_at = data.updated_at;
  }

  // Static methods for API operations
  static async getPublicCategories() {
    try {
      const response = await apiService.getPublicCategories();
      return response.map(cat => new Category(cat));
    } catch (error) {
      console.error('Failed to fetch public categories:', error);
      throw error;
    }
  }

  static async getAllCategories() {
    try {
      const response = await apiService.getAllCategories();
      return response.map(cat => new Category(cat));
    } catch (error) {
      console.error('Failed to fetch all categories:', error);
      throw error;
    }
  }

  static async getCategories() {
    try {
      const response = await apiService.getCategories();
      return response.map(cat => new Category(cat));
    } catch (error) {
      console.error('Failed to fetch categories:', error);
      throw error;
    }
  }

  static async getById(id) {
    try {
      const response = await apiService.getCategory(id);
      return new Category(response);
    } catch (error) {
      console.error('Failed to fetch category:', error);
      throw error;
    }
  }

  static async create(data) {
    try {
      const response = await apiService.createCategory(data);
      return new Category(response.data || response);
    } catch (error) {
      console.error('Failed to create category:', error);
      throw error;
    }
  }

  async update(data) {
    try {
      const response = await apiService.updateCategory(this.id, data);
      const updatedData = response.data || response;
      Object.assign(this, updatedData);
      return this;
    } catch (error) {
      console.error('Failed to update category:', error);
      throw error;
    }
  }

  async delete() {
    try {
      await apiService.deleteCategory(this.id);
      return true;
    } catch (error) {
      console.error('Failed to delete category:', error);
      throw error;
    }
  }
}
