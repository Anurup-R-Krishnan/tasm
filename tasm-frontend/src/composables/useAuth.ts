import { ref, computed } from 'vue';
import { login as apiLogin, getMe, type LoginCredentials, type AuthResponse } from '../api/auth';
import type { SystemUser } from '../types/models';

const TOKEN_KEY = 'tasm_auth_token';
const isE2E = import.meta.env['VITE_E2E'] === 'true';

// Global shared state
const currentUser = ref<SystemUser | null>(null);
const token = ref<string | null>(localStorage.getItem(TOKEN_KEY));
const isInitializing = ref(true);

export function useAuth() {
  const isAuthenticated = computed(() => !!token.value);

  const setToken = (newToken: string | null) => {
    token.value = newToken;
    if (newToken) {
      localStorage.setItem(TOKEN_KEY, newToken);
    } else {
      localStorage.removeItem(TOKEN_KEY);
    }
  };

  const login = async (credentials: LoginCredentials) => {
    try {
      const response: AuthResponse = await apiLogin(credentials);
      setToken(response.token);
      currentUser.value = response.user;
      return true;
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  };

  const logout = () => {
    setToken(null);
    currentUser.value = null;
    if (typeof window !== 'undefined') {
      window.location.href = '/login';
    }
  };

  const initAuth = async () => {
    isInitializing.value = true;
    if (isE2E) {
      if (!token.value) {
        setToken('e2e-dummy-token');
      }
      isInitializing.value = false;
      return;
    }

    if (token.value) {
      try {
        const user = await getMe();
        currentUser.value = user;
      } catch (error) {
        console.error('Session restoration failed:', error);
        logout(); // Token invalid or expired
      }
    }
    isInitializing.value = false;
  };

  return {
    currentUser,
    token,
    isAuthenticated,
    isInitializing,
    login,
    logout,
    initAuth,
  };
}
