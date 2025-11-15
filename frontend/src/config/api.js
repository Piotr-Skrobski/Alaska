/**
 * API Configuration
 * Centralized API URLs and endpoints
 */

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export const API_ENDPOINTS = {
  // Auth endpoints
  AUTH: {
    REGISTER: `${API_BASE_URL}/users/register`,
    LOGIN: `${API_BASE_URL}/users/login`,
    LOGOUT: `${API_BASE_URL}/users/logout`,
    ME: `${API_BASE_URL}/users/me`,
    DELETE: `${API_BASE_URL}/users/delete`,
  },

  // Movie endpoints
  MOVIES: {
    BY_TITLE: (title) => `${API_BASE_URL}/movies/title/${encodeURIComponent(title)}`,
    BY_IMDB_ID: (id) => `${API_BASE_URL}/movies/imdb/${id}`,
    BY_ID: (id) => `${API_BASE_URL}/movies/${id}`,
    CREATE: `${API_BASE_URL}/movies`,
    DELETE: (id) => `${API_BASE_URL}/movies/${id}`,
  },

  // Review endpoints
  REVIEWS: {
    CREATE: `${API_BASE_URL}/reviews`,
    BY_MOVIE: (movieId) => `${API_BASE_URL}/reviews/movie/${movieId}`,
    BY_USER: (userId) => `${API_BASE_URL}/reviews/user/${userId}`,
  },

  // LLM endpoints (Ollama)
  LLM: {
    GENERATE: '/ollama/api/generate',
  }
}

export default {
  BASE_URL: API_BASE_URL,
  ENDPOINTS: API_ENDPOINTS
}
