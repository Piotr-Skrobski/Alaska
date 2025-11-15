/**
 * Review Service
 */
import apiClient from './api'
import { API_ENDPOINTS } from '@/config/api'

export default {
  /**
   * Create a review
   * @param {Object} reviewData - { user_id, movie_id, rating, comment }
   * @returns {Promise}
   */
  async create(reviewData) {
    const response = await apiClient.post(API_ENDPOINTS.REVIEWS.CREATE, reviewData)
    return response.data
  },

  /**
   * Get reviews for a movie
   * @param {string} movieId
   * @returns {Promise}
   */
  async getByMovie(movieId) {
    const response = await apiClient.get(API_ENDPOINTS.REVIEWS.BY_MOVIE(movieId))
    return response.data
  },

  /**
   * Get reviews by a user
   * @param {number} userId
   * @returns {Promise}
   */
  async getByUser(userId) {
    const response = await apiClient.get(API_ENDPOINTS.REVIEWS.BY_USER(userId))
    return response.data
  }
}
