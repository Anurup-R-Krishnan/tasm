const defaultApiBaseUrl = 'http://localhost:8080/api'

export const API_BASE_URL =
  import.meta.env['VITE_API_BASE_URL']?.replace(/\/$/, '') ?? defaultApiBaseUrl

export class ApiError extends Error {
  status: number

  constructor(message: string, status: number) {
    super(message)
    this.name = 'ApiError'
    this.status = status
  }
}

export async function apiRequest<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers: {
      Accept: 'application/json',
      ...(init?.headers ?? {}),
    },
  })

  if (!response.ok) {
    let message = `Request failed with status ${response.status}`

    try {
      const payload = (await response.json()) as { error?: string }
      if (payload.error) {
        message = payload.error
      }
    } catch {
      // Fall back to the default message when the response is not JSON.
    }

    throw new ApiError(message, response.status)
  }

  return (await response.json()) as T
}
