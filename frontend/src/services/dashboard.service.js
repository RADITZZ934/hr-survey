import api from './api';
export const getStats = (params) => api.get('/dashboard/stats', { params });
export const getSurveyTrends = (params) => api.get('/dashboard/trends', { params });

