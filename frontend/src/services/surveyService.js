import api from './api';

export const surveyService = {
  /**
   * Mengirimkan jawaban/respon survey dari customer ke backend.
   * Endpoint: POST /surveys/submit
   *
   * @param {Object} payload
   * @param {number|string} payload.survey_id     - ID survey yang sedang diisi
   * @param {number|string} payload.id_store      - ID toko tempat QR code discan
   * @param {string}        payload.nama_responden - Nama customer yang mengisi survey
   * @param {string}        payload.penilaian     - Nilai dropdown ('sangat_baik', 'baik', 'kurang')
   * @returns {Promise<{ status: string, message: string }>}
   */
  async submitResponse(payload) {
    try {
      const response = await api.post('/surveys/submit', payload);
      return response.data;
    } catch (error) {
      console.error('Error submitting survey:', error);
      throw error;
    }
  }
};
