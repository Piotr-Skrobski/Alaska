/**
 * Movie Store (Pinia)
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import movieService from '@/services/movieService'
import reviewService from '@/services/reviewService'

export const useMovieStore = defineStore('movie', () => {
  // State
  const currentMovie = ref(null)
  const reviews = ref([])
  const loading = ref(false)
  const error = ref(null)

  // Actions
  async function searchByTitle(title) {
    loading.value = true
    error.value = null
    reviews.value = []

    try {
      const movie = await movieService.searchByTitle(title)
      currentMovie.value = movie

      // Load reviews for the movie
      if (movie?.imdb_id) {
        await loadReviews(movie.imdb_id)
      }

      return movie
    } catch (err) {
      error.value = err.response?.data || 'Error finding movie'
      currentMovie.value = null
      reviews.value = []
      throw err
    } finally {
      loading.value = false
    }
  }

  async function getByImdbId(imdbId) {
    loading.value = true
    error.value = null
    reviews.value = []

    try {
      const movie = await movieService.getByImdbId(imdbId)
      currentMovie.value = movie

      // Load reviews for the movie
      if (movie?.imdb_id) {
        await loadReviews(movie.imdb_id)
      }

      return movie
    } catch (err) {
      error.value = err.response?.data || 'Error finding movie'
      currentMovie.value = null
      reviews.value = []
      throw err
    } finally {
      loading.value = false
    }
  }

  async function loadReviews(movieId) {
    if (!movieId) {
      reviews.value = []
      return
    }

    try {
      const data = await reviewService.getByMovie(movieId)
      reviews.value = Array.isArray(data) ? data : []
    } catch (err) {
      console.error('Error loading reviews:', err)
      reviews.value = []
    }
  }

  function clearError() {
    error.value = null
  }

  function clearCurrentMovie() {
    currentMovie.value = null
    reviews.value = []
    error.value = null
  }

  return {
    // State
    currentMovie,
    reviews,
    loading,
    error,

    // Actions
    searchByTitle,
    getByImdbId,
    loadReviews,
    clearError,
    clearCurrentMovie
  }
})
