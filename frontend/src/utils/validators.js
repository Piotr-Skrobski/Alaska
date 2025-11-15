/**
 * Form Validation Utilities
 */

/**
 * Validate email format
 * @param {string} email
 * @returns {boolean}
 */
export function validateEmail(email) {
  if (!email) return false
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

/**
 * Validate password strength
 * @param {string} password
 * @returns {boolean}
 */
export function validatePassword(password) {
  if (!password) return false
  return password.length >= 6
}

/**
 * Validate username
 * @param {string} username
 * @returns {boolean}
 */
export function validateUsername(username) {
  if (!username) return false
  return username.length >= 3 && username.length <= 50
}

/**
 * Validate rating (1-10)
 * @param {number} rating
 * @returns {boolean}
 */
export function validateRating(rating) {
  return rating >= 1 && rating <= 10
}

/**
 * Validate required field
 * @param {*} value
 * @returns {boolean}
 */
export function validateRequired(value) {
  if (typeof value === 'string') {
    return value.trim().length > 0
  }
  return value !== null && value !== undefined
}

/**
 * Get password strength
 * @param {string} password
 * @returns {object} { strength: 'weak'|'medium'|'strong', score: 0-100 }
 */
export function getPasswordStrength(password) {
  if (!password) return { strength: 'weak', score: 0 }

  let score = 0

  // Length
  if (password.length >= 8) score += 25
  if (password.length >= 12) score += 25

  // Complexity
  if (/[a-z]/.test(password)) score += 10
  if (/[A-Z]/.test(password)) score += 10
  if (/[0-9]/.test(password)) score += 10
  if (/[^a-zA-Z0-9]/.test(password)) score += 20

  let strength = 'weak'
  if (score >= 70) strength = 'strong'
  else if (score >= 40) strength = 'medium'

  return { strength, score }
}
