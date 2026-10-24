import api from './api';

export const storeService = {
  /**
   * Mengambil daftar toko dari endpoint /api/LIT/store
   */
  async getStores() {
    try {
      const response = await api.get('/LIT/store');
      const raw = Array.isArray(response.data) ? response.data : (response.data?.data || []);
      
      return raw.map(item => {
        const id = item.id ?? item.Id ?? item.id_store;
        const kode = item.kode ?? item.Kode ?? `LB-${id}`;
        const name = item.name ?? item.Name ?? item.nama_store ?? 'Toko Tanpa Nama';
        const koordinat = (item.koordinat ?? item.Koordinat ?? '').trim();
        const aplikasiId = item.aplikasi_id ?? item.AplikasiId ?? 1;

        let latitude = null;
        let longitude = null;
        if (koordinat && koordinat.includes(',')) {
          const parts = koordinat.split(',').map(s => parseFloat(s.trim()));
          if (!isNaN(parts[0]) && !isNaN(parts[1])) {
            latitude = parts[0];
            longitude = parts[1];
          }
        }

        return {
          id,
          kode,
          name,
          koordinat,
          latitude,
          longitude,
          aplikasi_id: aplikasiId,
          raw: item
        };
      });
    } catch (error) {
      console.error('Error in storeService.getStores:', error);
      throw error;
    }
  },

  /**
   * Menyinkronkan daftar toko dari API pusat Laskar Buah dan stores.json
   */
  async syncStores() {
    try {
      const response = await api.post('/LIT/store/sync');
      return response.data;
    } catch (error) {
      console.error('Error in storeService.syncStores:', error);
      throw error;
    }
  }
};

export default storeService;
