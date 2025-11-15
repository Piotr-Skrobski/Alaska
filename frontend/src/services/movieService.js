/**
 * Movie Service
 */
import apiClient from './api'
import { API_ENDPOINTS } from '@/config/api'

export default {
  /**
   * Search movie by title
   * @param {string} title
   * @returns {Promise}
   */
  async searchByTitle(title) {
    const response = await apiClient.get(API_ENDPOINTS.MOVIES.BY_TITLE(title))
    return response.data
  },

  /**
   * Get movie by IMDb ID
   * @param {string} imdbId
   * @returns {Promise}
   */
  async getByImdbId(imdbId) {
    const response = await apiClient.get(API_ENDPOINTS.MOVIES.BY_IMDB_ID(imdbId))
    return response.data
  },

  /**
   * Get movie by ID
   * @param {string} id
   * @returns {Promise}
   */
  async getById(id) {
    const response = await apiClient.get(API_ENDPOINTS.MOVIES.BY_ID(id))
    return response.data
  },

  /**
   * Create/update a movie
   * @param {Object} movieData
   * @returns {Promise}
   */
  async create(movieData) {
    const response = await apiClient.post(API_ENDPOINTS.MOVIES.CREATE, movieData)
    return response.data
  },

  /**
   * Delete a movie
   * @param {string} id
   * @returns {Promise}
   */
  async delete(id) {
    const response = await apiClient.delete(API_ENDPOINTS.MOVIES.DELETE(id))
    return response.data
  }
}
