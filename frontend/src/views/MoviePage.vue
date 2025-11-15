<template>
  <div class="container mt-4">
    <MovieSearchSection :loading="movieStore.loading" @search="handleSearch" />

    <div v-if="movieStore.loading" class="text-center my-5">
      <BaseSpinner />
    </div>

    <ErrorAlert :message="movieStore.error" @dismiss="movieStore.clearError()" />

    <template v-if="movieStore.currentMovie">
      <MovieDetail :movie="movieStore.currentMovie" />
      <ReviewsList :reviews="movieStore.reviews" />

      <ReviewForm
        v-if="authStore.isLoggedIn"
        :movie-id="movieStore.currentMovie.imdb_id"
        :movie-title="movieStore.currentMovie.title"
        @submitted="movieStore.loadReviews(movieStore.currentMovie.imdb_id)"
      />

      <BaseAlert v-else variant="info">
        <p class="mb-0">
          Please <router-link to="/login">login</router-link> to write a review.
        </p>
      </BaseAlert>
    </template>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useMovieStore } from '@/stores/movie'
import { useAuthStore } from '@/stores/auth'
import BaseSpinner from '@/components/atoms/BaseSpinner.vue'
import BaseAlert from '@/components/atoms/BaseAlert.vue'
import MovieSearchSection from '@/components/organisms/MovieSearchSection.vue'
import MovieDetail from '@/components/organisms/MovieDetail.vue'
import ReviewsList from '@/components/organisms/ReviewsList.vue'
import ReviewForm from '@/components/organisms/ReviewForm.vue'
import ErrorAlert from '@/components/molecules/ErrorAlert.vue'

const route = useRoute()
const movieStore = useMovieStore()
const authStore = useAuthStore()

onMounted(() => {
  const movieId = route.params.id
  if (movieId) {
    movieStore.getByImdbId(movieId)
  }
})

const handleSearch = async (title) => {
  await movieStore.searchByTitle(title)
}
</script>
