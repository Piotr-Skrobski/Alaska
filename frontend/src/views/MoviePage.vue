<template>
  <div class="container mt-4">
    <div class="card mb-4">
      <div class="card-header bg-primary text-white">
        <h3>Find a movie</h3>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-8">
            <div class="input-group mb-3">
              <input type="text" class="form-control" placeholder="Enter movie title" v-model="searchTitle"
                @keyup.enter="searchMovie">
              <button class="btn btn-primary" type="button" @click="searchMovie">
                Search
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="loading" class="text-center my-5">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>

    <div v-if="error" class="alert alert-danger" role="alert">
      {{ error }}
    </div>

    <div v-if="currentMovie" class="card mb-4">
      <div class="card-header bg-secondary text-white">
        <h3>{{ currentMovie.title }}</h3>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-4">
            <img :src="currentMovie.poster_url || 'https://via.placeholder.com/300x450'" class="img-fluid rounded"
              alt="Movie Poster">
          </div>
          <div class="col-md-8">
            <h4>{{ currentMovie.title }} ({{ currentMovie.year }})</h4>
            <p class="text-muted">IMDb ID: {{ currentMovie.imdb_id }}</p>
            <div class="mb-3">
              <span class="badge bg-warning text-dark me-2">Rating: {{ currentMovie.imdb_rating || 'N/A' }}</span>
              <span class="badge bg-info text-dark me-2">{{ currentMovie.genre }}</span>
            </div>
            <p>{{ currentMovie.plot }}</p>
            <p><strong>Director:</strong> {{ currentMovie.director }}</p>
            <p><strong>Actors:</strong> {{ currentMovie.actors }}</p>
          </div>
        </div>
      </div>
    </div>

    <div v-if="currentMovie" class="card mb-4">
      <div class="card-header bg-success text-white d-flex justify-content-between align-items-center">
        <h3>Reviews</h3>
        <span class="badge bg-light text-dark">{{ reviews ? reviews.length : 0 }} Reviews</span>
      </div>
      <div class="card-body">
        <div v-if="!reviews || reviews.length === 0" class="text-center py-3">
          <p class="text-muted">No reviews yet. Be the first to review!</p>
        </div>
        <div v-else class="mb-4">
          <div v-for="review in reviews" :key="review.id" class="card mb-3">
            <div class="card-header d-flex justify-content-between">
              <div>
                <strong>User {{ review.userID }}</strong>
              </div>
              <div>
                <span class="badge bg-warning text-dark">Rating: {{ review.rating }}/10</span>
              </div>
            </div>
            <div class="card-body">
              <p>{{ review.comment }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="currentMovie && isLoggedIn" class="card mb-4">
      <div class="card-header bg-primary text-white">
        <h3>Write a review</h3>
      </div>
      <div class="card-body">
        <form @submit.prevent="submitReview">
          <div class="mb-3">
            <label class="form-label">Rating (1-10)</label>
            <div class="star-rating">
              <i v-for="n in 10" :key="n" class="fa-star"
                :class="n <= newReview.rating ? 'fas text-warning' : 'far text-muted'"
                style="cursor: pointer; font-size: 1.5rem; margin-right: 4px;" @click="setRating(n)">
              </i>
            </div>
            <div class="mt-2">Selected rating: {{ newReview.rating }}/10</div>
          </div>
          <div class="mb-3">
            <label for="comment" class="form-label">Your review</label>
            <textarea class="form-control" id="comment" rows="4" v-model="newReview.comment" required
              placeholder="Share your thoughts about this movie..."></textarea>
          </div>

          <button type="submit" class="btn btn-primary" :disabled="submitting">
            <span v-if="submitting" class="spinner-border spinner-border-sm me-2" role="status"
              aria-hidden="true"></span>
            Submit review
          </button>
        </form>

        <div class="mt-4">
          <button class="btn btn-secondary" :disabled="llmSubmitting" @click="generateReviewFromLLM">
            <span v-if="llmSubmitting" class="spinner-border spinner-border-sm me-2" role="status"
              aria-hidden="true"></span>
            Generate Review from LLM
          </button>
        </div>
      </div>
    </div>

    <div v-if="currentMovie && !isLoggedIn" class="alert alert-info">
      <p>Please <router-link to="/login">login</router-link> to write a review.</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import authService from '@/services/auth';

const API_URL = 'http://localhost:8080';

export default {
  name: 'MoviePage',
  data() {
    return {
      searchTitle: '',
      currentMovie: null,
      reviews: [],
      loading: false,
      error: null,
      submitting: false,
      newReview: {
        rating: 5,
        comment: ''
      },
      isLoggedIn: false,
      currentUser: null,
      llmSubmitting: false,
    };
  },

  created() {
    this.checkAuthStatus();

    const movieId = this.$route.params.id;
    if (movieId) {
      this.fetchMovieByImdbId(movieId);
    }
  },

  methods: {
    setRating(rating) {
      this.newReview.rating = rating;
      console.log(`Rating set to: ${rating}`);
    },

    checkAuthStatus() {
      this.isLoggedIn = authService.isLoggedIn();
      if (this.isLoggedIn) {
        this.currentUser = authService.getUser();
      }
    },

    searchMovie() {
      if (!this.searchTitle.trim()) return;

      this.loading = true;
      this.error = null;
      this.reviews = [];

      axios.get(`${API_URL}/movies/title/${encodeURIComponent(this.searchTitle)}`)
        .then(response => {
          console.log(response);
          this.currentMovie = response.data;
          this.loadReviews();
        })
        .catch(error => {
          this.error = `Error finding movie: ${error.response?.data || error.message}`;
          this.currentMovie = null;
          this.reviews = [];
        })
        .finally(() => {
          this.loading = false;
        });
    },

    fetchMovieByImdbId(imdbId) {
      this.loading = true;
      this.error = null;
      this.reviews = [];

      axios.get(`${API_URL}/movies/imdb/${imdbId}`)
        .then(response => {
          this.currentMovie = response.data;
          this.loadReviews();
        })
        .catch(error => {
          this.error = `Error finding movie: ${error.response?.data || error.message}`;
          this.currentMovie = null;
          this.reviews = [];
        })
        .finally(() => {
          this.loading = false;
        });
    },

    loadReviews() {
      if (!this.currentMovie || !this.currentMovie.imdb_id) {
        this.reviews = [];
        return;
      }

      axios.get(`${API_URL}/reviews/movie/${this.currentMovie.imdb_id}`)
        .then(response => {
          console.log('Reviews loaded:', response.data);
          this.reviews = Array.isArray(response.data) ? response.data : [];
        })
        .catch(error => {
          console.error('Error loading reviews:', error);
          this.reviews = [];
        });
    },

    submitReview() {
      if (!this.isLoggedIn || !this.currentMovie) return;

      this.submitting = true;

      const reviewData = {
        user_id: this.currentUser.user.id,
        movie_id: this.currentMovie.imdb_id,
        rating: this.newReview.rating,
        comment: this.newReview.comment
      };

      axios.post(`${API_URL}/reviews`, reviewData, {
        withCredentials: true
      })
        .then(() => {
          this.loadReviews();

          this.newReview = {
            rating: 5,
            comment: ''
          };

          alert('Review submitted successfully!');
        })
        .catch(error => {
          this.error = `Error submitting review: ${error.response?.data || error.message}`;
        })
        .finally(() => {
          this.submitting = false;
        });
    },

    generateReviewFromLLM() {
      if (!this.isLoggedIn || !this.currentMovie) return;

      this.llmSubmitting = true;

      const prompt = `Write a short, ${this.newReview.rating}/10, review for movie ${this.currentMovie.title}.`;

      const requestData = {
        model: "llama3",
        prompt: prompt,
        stream: false
      };

      axios.post('/ollama/api/generate', requestData, {
        headers: {
          "Content-Type": "application/json"
        }
      })
        .then((response) => {
          const reviewText = response.data.response;
          this.newReview.comment = reviewText;
        })
        .catch((error) => {
          console.error('Error generating review:', error);
          alert('Error generating review from LLM');
        })
        .finally(() => {
          this.llmSubmitting = false;
        });
    },
  }
};
</script>