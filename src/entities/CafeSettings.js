import apiService from '../services/api';

export class CafeSettings {
  constructor(data = {}) {
    this.id = data.id;
    this.name = data.name || '';
    this.location = data.location || '';
    this.phone = data.phone || '';
    this.instagram_url = data.instagram_url || '';
    this.logo_url = data.logo_url || '';
    this.hero_image_url = data.hero_image_url || '';
    this.description = data.description || '';
    this.is_active = data.is_active !== undefined ? data.is_active : true;
    this.created_at = data.created_at;
    this.updated_at = data.updated_at;
  }

  // Static methods for API operations
  static async getPublicSettings() {
    try {
      const response = await apiService.getPublicShopSettings();
      return new CafeSettings(response);
    } catch (error) {
      console.error('Failed to fetch public shop settings:', error);
      throw error;
    }
  }

  static async getSettings() {
    try {
      const response = await apiService.getShopSettings();
      return new CafeSettings(response);
    } catch (error) {
      console.error('Failed to fetch shop settings:', error);
      throw error;
    }
  }

  async update(data) {
    try {
      const response = await apiService.updateShopSettings(data);
      const updatedData = response.data || response;
      Object.assign(this, updatedData);
      return this;
    } catch (error) {
      console.error('Failed to update shop settings:', error);
      throw error;
    }
  }

  // Legacy methods for backward compatibility
  static async list() {
    console.log('CafeSettings.list() is deprecated. Use getSettings() instead.');
    try {
      const settings = await this.getSettings();
      return [settings];
    } catch (error) {
      return [new CafeSettings()];
    }
  }

  static async getById() {
    console.log('CafeSettings.getById() is deprecated. Use getSettings() instead.');
    return this.getSettings();
  }

  static async getDefault() {
    console.log('CafeSettings.getDefault() is deprecated. Use getSettings() instead.');
    return this.getSettings();
  }

  static async create() {
    console.log('CafeSettings.create() is deprecated. Use the backend admin panel instead.');
    return new CafeSettings();
  }

  static async update() {
    console.log('CafeSettings.update() is deprecated. Use instance.update() instead.');
  }

  static async delete() {
    console.log('CafeSettings.delete() is deprecated. Use the backend admin panel instead.');
  }

  static async save() {
    console.log('CafeSettings.save() is deprecated. Use instance.update() instead.');
  }

  static async seed() {
    console.log('CafeSettings.seed() is deprecated. Use the backend seeding instead.');
    return new CafeSettings();
  }
}
