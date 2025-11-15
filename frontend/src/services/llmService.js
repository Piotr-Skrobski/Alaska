/**
 * LLM Service (Ollama)
 */
import apiClient from './api'
import { API_ENDPOINTS } from '@/config/api'

export default {
  /**
   * Generate review from LLM
   * @param {string} movieTitle
   * @param {number} rating
   * @returns {Promise<string>}
   */
  async generateReview(movieTitle, rating) {
    const prompt = `Write a short, ${rating}/10, review for movie ${movieTitle}.`

    const requestData = {
      model: 'llama3',
      prompt: prompt,
      stream: false
    }

    const response = await apiClient.post(API_ENDPOINTS.LLM.GENERATE, requestData, {
      headers: {
        'Content-Type': 'application/json'
      }
    })

    return response.data.response
  },

  /**
   * Generate text from custom prompt
   * @param {string} prompt
   * @param {string} model
   * @returns {Promise<string>}
   */
  async generate(prompt, model = 'llama3') {
    const requestData = {
      model: model,
      prompt: prompt,
      stream: false
    }

    const response = await apiClient.post(API_ENDPOINTS.LLM.GENERATE, requestData, {
      headers: {
        'Content-Type': 'application/json'
      }
    })

    return response.data.response
  }
}
