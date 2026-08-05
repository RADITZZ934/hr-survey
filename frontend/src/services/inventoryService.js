import api from './api';

export const inventoryService = {
  /**
   * Mengambil data daftar toko (id_store, nama, koordinat) dari sistem inventory.
   * Endpoint: GET /LIT/store
   * Pastikan VITE_API_URL di .env mengarah ke backend yang menyediakan endpoint ini.
   */
  async getLocations() {
    try {
      const response = await api.get('/LIT/store');
      // Menangani berbagai format response: array langsung atau objek { data: [...] }
      return Array.isArray(response.data) ? response.data : response.data.data || [];
    } catch (error) {
      console.error('Error fetching locations:', error);
      throw error;
    }
  }
};
