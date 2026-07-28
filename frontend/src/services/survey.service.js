import api from './api';
export const getSurveys = () => api.get('/surveys');
export const getSurveyReport = (id, params) => api.get(`/surveys/${id}/report`, { params });
export const getSurveyResponses = (id, params) => api.get(`/surveys/${id}/responses`, { params });
export const createSurvey = (data) => api.post('/surveys', data);
export const getSurveyQuestions = (id) => api.get(`/surveys/${id}/questions`);
export const submitSurveyResponse = (id, data) => api.post(`/surveys/${id}/responses`, data);
export const deleteSurvey = (id) => api.delete(`/surveys/${id}`);
