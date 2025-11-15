<template>
  <BaseCard v-if="movie" headerClass="bg-secondary text-white">
    <template #header>
      <h3 class="mb-0">{{ movie.title }}</h3>
    </template>

    <div class="row">
      <div class="col-md-4">
        <img
          :src="movie.poster_url || 'https://via.placeholder.com/300x450'"
          class="img-fluid rounded"
          :alt="`${movie.title} Poster`"
        >
      </div>
      <div class="col-md-8">
        <h4>{{ movie.title }} ({{ movie.year }})</h4>
        <p class="text-muted">IMDb ID: {{ movie.imdb_id }}</p>

        <div class="mb-3">
          <BaseBadge variant="warning" class="me-2">
            Rating: {{ movie.imdb_rating || 'N/A' }}
          </BaseBadge>
          <BaseBadge variant="info">
            {{ movie.genre }}
          </BaseBadge>
        </div>

        <p>{{ movie.plot }}</p>
        <p><strong>Director:</strong> {{ movie.director }}</p>
        <p><strong>Actors:</strong> {{ movie.actors }}</p>
      </div>
    </div>
  </BaseCard>
</template>

<script setup>
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseBadge from '@/components/atoms/BaseBadge.vue'

defineProps({
  movie: {
    type: Object,
    required: true
  }
})
</script>
