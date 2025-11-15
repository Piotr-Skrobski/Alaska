<template>
  <BaseCard headerClass="bg-success text-white">
    <template #header>
      <div class="d-flex justify-content-between align-items-center">
        <h3 class="mb-0">Reviews</h3>
        <BaseBadge variant="light" class="text-dark">
          {{ reviews ? reviews.length : 0 }} Reviews
        </BaseBadge>
      </div>
    </template>

    <div v-if="!reviews || reviews.length === 0" class="text-center py-3">
      <p class="text-muted mb-0">No reviews yet. Be the first to review!</p>
    </div>

    <div v-else>
      <ReviewCard
        v-for="review in reviews"
        :key="review.id"
        :review="review"
        class="mb-3"
      />
    </div>
  </BaseCard>
</template>

<script setup>
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseBadge from '@/components/atoms/BaseBadge.vue'
import ReviewCard from '@/components/molecules/ReviewCard.vue'

defineProps({
  reviews: {
    type: Array,
    default: () => []
  }
})
</script>
