import api from './api';

export const getActionPlans = (surveyId) => {
  const params = surveyId ? { survey_id: surveyId } : {};
  return api.get('/action-plans', { params });
};

export const createActionPlan = (data) => api.post('/action-plans', data);

export const updateActionPlan = (id, data) => api.put(`/action-plans/${id}`, data);
