/**
 * Axios API Client with interceptors
 */
import axios from 'axios'
import { API_ENDPOINTS } from '@/config/api'

// Create axios instance
const apiClient = axios.create({
  withCredentials: true, // Important for cookies
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    // You can add auth tokens here if needed
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
apiClient.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    // Handle common errors
    if (error.response) {
      switch (error.response.status) {
        case 401:
          // Unauthorized - could trigger logout
          console.error('Unauthorized access')
          break
        case 403:
          console.error('Forbidden')
          break
        case 404:
          console.error('Resource not found')
          break
        case 500:
          console.error('Server error')
          break
      }
    }

    return Promise.reject(error)
  }
)

export default apiClient
