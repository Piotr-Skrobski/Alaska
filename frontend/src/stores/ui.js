/**
 * UI Store (Pinia)
 * Manages UI state like modals, alerts, etc.
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUIStore = defineStore('ui', () => {
  // State
  const confirmDialog = ref({
    isOpen: false,
    title: '',
    message: '',
    confirmText: 'Confirm',
    cancelText: 'Cancel',
    confirmVariant: 'danger',
    onConfirm: null,
    onCancel: null
  })

  const successMessage = ref(null)
  const errorMessage = ref(null)

  // Actions
  function showConfirm(options) {
    return new Promise((resolve) => {
      confirmDialog.value = {
        isOpen: true,
        title: options.title || 'Confirm',
        message: options.message || 'Are you sure?',
        confirmText: options.confirmText || 'Confirm',
        cancelText: options.cancelText || 'Cancel',
        confirmVariant: options.confirmVariant || 'danger',
        onConfirm: () => {
          confirmDialog.value.isOpen = false
          resolve(true)
        },
        onCancel: () => {
          confirmDialog.value.isOpen = false
          resolve(false)
        }
      }
    })
  }

  function hideConfirm() {
    confirmDialog.value.isOpen = false
  }

  function showSuccess(message, duration = 3000) {
    successMessage.value = message
    if (duration) {
      setTimeout(() => {
        successMessage.value = null
      }, duration)
    }
  }

  function hideSuccess() {
    successMessage.value = null
  }

  function showError(message, duration = 5000) {
    errorMessage.value = message
    if (duration) {
      setTimeout(() => {
        errorMessage.value = null
      }, duration)
    }
  }

  function hideError() {
    errorMessage.value = null
  }

  return {
    // State
    confirmDialog,
    successMessage,
    errorMessage,

    // Actions
    showConfirm,
    hideConfirm,
    showSuccess,
    hideSuccess,
    showError,
    hideError
  }
})
