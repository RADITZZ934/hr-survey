import { createRouter, createWebHistory } from 'vue-router';
import AdminLayout from '../layouts/AdminLayout.vue';
import AuthLayout from '../layouts/AuthLayout.vue';
import SurveyLayout from '../layouts/SurveyLayout.vue';

const routes = [
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/auth',
    component: AuthLayout,
    children: [
      { path: 'login', component: () => import('../pages/auth/Login.vue') }
    ]
  },
  {
    path: '/dashboard',
    component: AdminLayout,
    meta: { requiresAuth: true },
    children: [
      { path: '', component: () => import('../pages/dashboard/Dashboard.vue') },
      { path: 'surveys', component: () => import('../pages/surveys/SurveyList.vue') },
      { path: 'reports', component: () => import('../pages/reports/Reports.vue') },
      { path: 'action-plans', component: () => import('../pages/action-plans/ActionPlans.vue') },
      { path: 'employees', component: () => import('../pages/employees/EmployeeManagement.vue') }
    ]
  },
  {
    path: '/survey',
    component: SurveyLayout,
    children: [
      { path: 'identity', component: () => import('../pages/public-survey/Identity.vue') },
      { path: 'form', component: () => import('../pages/public-survey/SurveyForm.vue') },
      { path: 'thanks', component: () => import('../pages/public-survey/ThankYou.vue') }
    ]
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!token) {
      next({ path: '/auth/login' });
    } else {
      next();
    }
  } else if (to.path === '/auth/login' && token) {
    next({ path: '/dashboard' });
  } else {
    next();
  }
});

export default router;
