/**
 * Auth Store (Pinia)
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import authService from '@/services/authService'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Getters
  const isLoggedIn = computed(() => user.value !== null)
  const username = computed(() => user.value?.username || '')
  const userId = computed(() => user.value?.id || null)

  // Actions
  async function login(credentials) {
    loading.value = true
    error.value = null

    try {
      const response = await authService.login(credentials)
      user.value = response.user || response
      return response
    } catch (err) {
      error.value = err.response?.data?.message || 'Login failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function register(userData) {
    loading.value = true
    error.value = null

    try {
      const response = await authService.register(userData)
      // Auto-login after registration
      if (response.user) {
        user.value = response.user
      }
      return response
    } catch (err) {
      error.value = err.response?.data?.message || 'Registration failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function logout() {
    loading.value = true
    error.value = null

    try {
      await authService.logout()
      user.value = null
    } catch (err) {
      error.value = err.response?.data?.message || 'Logout failed'
      // Clear user anyway on logout error
      user.value = null
    } finally {
      loading.value = false
    }
  }

  async function fetchCurrentUser() {
    loading.value = true
    error.value = null

    try {
      const response = await authService.getCurrentUser()
      user.value = response.user || response
      return response
    } catch (err) {
      error.value = err.response?.data?.message || 'Failed to fetch user'
      user.value = null
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteAccount() {
    loading.value = true
    error.value = null

    try {
      await authService.deleteAccount()
      user.value = null
    } catch (err) {
      error.value = err.response?.data?.message || 'Failed to delete account'
      throw err
    } finally {
      loading.value = false
    }
  }

  function clearError() {
    error.value = null
  }

  return {
    // State
    user,
    loading,
    error,

    // Getters
    isLoggedIn,
    username,
    userId,

    // Actions
    login,
    register,
    logout,
    fetchCurrentUser,
    deleteAccount,
    clearError
  }
})
