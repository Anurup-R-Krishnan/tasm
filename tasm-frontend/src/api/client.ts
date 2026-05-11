const defaultApiBaseUrl = '/api';

const API_BASE_URL = import.meta.env['VITE_API_BASE_URL']?.replace(/\/$/, '') ?? defaultApiBaseUrl;

class ApiError extends Error {
  status: number;

  constructor(message: string, status: number) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
  }
}

export async function apiRequest<T>(
  path: string,
  init?: Omit<RequestInit, 'body'> & { body?: any },
): Promise<T> {
  const token = localStorage.getItem('tasm_auth_token');
  const headers: Record<string, string> = {
    Accept: 'application/json',
    ...((init?.headers as Record<string, string>) ?? {}),
  };

  let body = init?.body;
  if (body && typeof body === 'object' && !(body instanceof FormData) && !(body instanceof Blob)) {
    body = JSON.stringify(body);
    headers['Content-Type'] = 'application/json';
  }

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  let response: Response;
  let attempt = 0;
  const maxRetries = init?.method && init.method !== 'GET' ? 0 : 2;

  while (true) {
    try {
      response = await fetch(`${API_BASE_URL}${path}`, {
        ...init,
        body,
        headers,
      });
      break;
    } catch (error) {
      if (attempt >= maxRetries) throw error;
      attempt++;
      await new Promise((resolve) => setTimeout(resolve, 500 * Math.pow(2, attempt - 1)));
    }
  }

  if (!response.ok) {
    if (response.status === 401 && !path.includes('/auth/login') && !path.includes('/auth/me')) {
      localStorage.removeItem('tasm_auth_token');
      window.location.href = '/login';
    }

    let message = `Request failed with status ${response.status}`;

    try {
      const payload = (await response.json()) as { error?: string };
      if (payload.error) {
        message = payload.error;
      }
    } catch {
      // Fall back to the default message when the response is not JSON.
    }

    throw new ApiError(message, response.status);
  }

  return (await response.json()) as T;
}
