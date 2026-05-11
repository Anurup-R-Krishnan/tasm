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
  const { initAuth, isInitializing, isAuthenticated, isSetupCompleted } = useAuth();

  if (isInitializing.value) {
    await initAuth();
  }

  const isAuthPath = to.name === 'Login' || to.path.startsWith('/login');
  const isRegisterPath = to.name === 'Register' || to.path.startsWith('/register');
  const isWelcomePath = to.name === 'Welcome' || to.path.startsWith('/welcome');
  const isFirstRunPath = to.name === 'FirstRun' || to.path.startsWith('/first-run');
  const startedSetupFromWelcome = to.query['setup'] === '1';

  // /first-run: only reachable via the "Start Setup" button (?setup=1).
  // Direct URL access without the param → Welcome.
  // No restriction based on isSetupCompleted — users can always run setup.
  if (isFirstRunPath) {
    if (!startedSetupFromWelcome) return { name: 'Welcome' };
    return true;
  }

  // Unauthenticated user hitting a protected route → Welcome
  const requiresAuth = to.matched.some((record) => record.meta['requiresAuth']);
  if (requiresAuth && !isAuthenticated.value) return { name: 'Welcome' };

  // /register before setup is done → redirect into setup wizard
  if (isRegisterPath && !isSetupCompleted.value) {
    return { name: 'FirstRun', query: { setup: '1' } };
  }

  // Authenticated + setup done: bounce away from auth/welcome pages to the app
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
