import { createRouter, createWebHistory } from 'vue-router';
import AppLayout from '../components/AppLayout.vue';
import { appRoutes } from './routes';
import { useAuth } from '../composables/useAuth';

const childRoutes = appRoutes.map((route) => ({
  path: route.path,
  name: route.name,
  meta: { title: route.title },
  component: route.component,
}));

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/welcome',
      name: 'Welcome',
      component: () => import('../views/WelcomeView.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/first-run',
      name: 'FirstRun',
      component: () => import('../views/FirstRunSetup.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/LoginView.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('../views/RegisterView.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/',
      component: AppLayout,
      meta: { requiresAuth: true },
      children: childRoutes,
    },
  ],
});

router.beforeEach(async (to) => {
  const { initAuth, isInitializing, isAuthenticated, isSetupCompleted, isFirstRun, currentUser } =
    useAuth();

  if (isInitializing.value) {
    await initAuth();
  }

  const isAuthPath = to.name === 'Login' || to.path.startsWith('/login');
  const isRegisterPath = to.name === 'Register' || to.path.startsWith('/register');
  const isWelcomePath = to.name === 'Welcome' || to.path.startsWith('/welcome');
  const isFirstRunPath = to.name === 'FirstRun' || to.path.startsWith('/first-run');

  // If setup completed, disallow /first-run
  if (isSetupCompleted.value && isFirstRunPath) return { name: 'Dashboard' };

  // If user tries to open FirstRun when not allowed, send to Welcome
  if (isFirstRunPath && !isSetupCompleted.value) {
    if (
      !isFirstRun.value &&
      !(isAuthenticated.value && currentUser.value?.role === 'System Admin')
    ) {
      return { name: 'Welcome' };
    }
  }

  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);

  // Protected route attempt by unauthenticated user → Welcome
  if (requiresAuth && !isAuthenticated.value) return { name: 'Welcome' };

  // Prevent auth pages when already logged in and setup done
  if (
    isAuthenticated.value &&
    isSetupCompleted.value &&
    (isAuthPath || isRegisterPath || isWelcomePath)
  ) {
    return { name: 'Dashboard' };
  }

  return true;
});

export default router;
