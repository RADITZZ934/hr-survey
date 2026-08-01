import api from './api';

export const getEmployees = (params) => api.get('/admin/employees', { params });
export const createEmployee = (data) => api.post('/admin/employees', data);
export const updateEmployee = (id, data) => api.put(`/admin/employees/${id}`, data);
export const deleteEmployee = (id) => api.delete(`/admin/employees/${id}`);

export const getCriticalAlerts = () => api.get('/admin/alerts');
export const markAlertAsRead = (id) => api.put(`/admin/alerts/${id}/read`);
export const getDepartmentSatisfaction = (params) => api.get('/admin/departments/satisfaction', { params });
