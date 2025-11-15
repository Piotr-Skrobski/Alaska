/**
 * Authentication Service
 */
import apiClient from './api'
import { API_ENDPOINTS } from '@/config/api'

export default {
  /**
   * Register a new user
   * @param {Object} userData - { username, email, password }
   * @returns {Promise}
   */
  async register(userData) {
    const response = await apiClient.post(API_ENDPOINTS.AUTH.REGISTER, userData)
    return response.data
  },

  /**
   * Login user
   * @param {Object} credentials - { username, password }
   * @returns {Promise}
   */
  async login(credentials) {
    const response = await apiClient.post(API_ENDPOINTS.AUTH.LOGIN, credentials)
    return response.data
  },

  /**
   * Logout user
   * @returns {Promise}
   */
  async logout() {
    const response = await apiClient.post(API_ENDPOINTS.AUTH.LOGOUT)
    return response.data
  },

  /**
   * Get current user info
   * @returns {Promise}
   */
  async getCurrentUser() {
    const response = await apiClient.get(API_ENDPOINTS.AUTH.ME)
    return response.data
  },

  /**
   * Delete user account
   * @returns {Promise}
   */
  async deleteAccount() {
    const response = await apiClient.post(API_ENDPOINTS.AUTH.DELETE)
    return response.data
  }
}
