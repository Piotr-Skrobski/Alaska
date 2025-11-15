<template>
  <div class="card movie-card h-100" :class="{ 'shadow-sm': true }">
    <img
      :src="movie.poster_url || movie.Poster || 'https://via.placeholder.com/300x450'"
      class="card-img-top"
      :alt="movie.title || movie.Title"
      style="height: 400px; object-fit: cover;"
    >
    <div class="card-body d-flex flex-column">
      <h5 class="card-title">{{ movie.title || movie.Title }}</h5>
      <p class="text-muted mb-2">{{ movie.year || movie.Year }}</p>
      <div class="mb-3">
        <BaseBadge v-if="movie.imdb_rating || movie.imdbRating" variant="warning" class="me-2">
          {{ movie.imdb_rating || movie.imdbRating }}
        </BaseBadge>
        <BaseBadge v-if="movie.genre || movie.Genre" variant="info">
          {{ movie.genre || movie.Genre }}
        </BaseBadge>
      </div>
      <div class="mt-auto">
        <BaseButton variant="primary" size="sm" block @click="$emit('click', movie)">
          View Details
        </BaseButton>
      </div>
    </div>
  </div>
</template>

<script setup>
import BaseBadge from '@/components/atoms/BaseBadge.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'

defineProps({
  movie: {
    type: Object,
    required: true
  }
})

defineEmits(['click'])
</script>
