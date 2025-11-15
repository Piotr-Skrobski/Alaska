/**
 * Data Formatting Utilities
 */

/**
 * Format date to readable string
 * @param {string|Date} date
 * @returns {string}
 */
export function formatDate(date) {
  if (!date) return ''

  const d = new Date(date)
  if (isNaN(d.getTime())) return ''

  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  }).format(d)
}

/**
 * Format date with time
 * @param {string|Date} date
 * @returns {string}
 */
export function formatDateTime(date) {
  if (!date) return ''

  const d = new Date(date)
  if (isNaN(d.getTime())) return ''

  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(d)
}

/**
 * Format relative time (e.g., "2 hours ago")
 * @param {string|Date} date
 * @returns {string}
 */
export function formatRelativeTime(date) {
  if (!date) return ''

  const d = new Date(date)
  if (isNaN(d.getTime())) return ''

  const now = new Date()
  const diffMs = now - d
  const diffSecs = Math.floor(diffMs / 1000)
  const diffMins = Math.floor(diffSecs / 60)
  const diffHours = Math.floor(diffMins / 60)
  const diffDays = Math.floor(diffHours / 24)

  if (diffSecs < 60) return 'just now'
  if (diffMins < 60) return `${diffMins} minute${diffMins > 1 ? 's' : ''} ago`
  if (diffHours < 24) return `${diffHours} hour${diffHours > 1 ? 's' : ''} ago`
  if (diffDays < 7) return `${diffDays} day${diffDays > 1 ? 's' : ''} ago`

  return formatDate(d)
}

/**
 * Truncate text to max length
 * @param {string} text
 * @param {number} maxLength
 * @returns {string}
 */
export function truncateText(text, maxLength = 100) {
  if (!text || text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

/**
 * Format rating display
 * @param {number} rating
 * @param {number} maxRating
 * @returns {string}
 */
export function formatRating(rating, maxRating = 10) {
  if (rating === null || rating === undefined) return 'N/A'
  return `${rating}/${maxRating}`
}
