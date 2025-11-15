<template>
  <BaseCard :shadow="true">
    <div class="d-flex justify-content-between align-items-start mb-2">
      <div>
        <strong>{{ review.username || `User ${review.userID}` }}</strong>
        <small v-if="review.created_at" class="text-muted d-block">
          {{ formatDate(review.created_at) }}
        </small>
      </div>
      <BaseBadge variant="warning">{{ review.rating }}/10</BaseBadge>
    </div>
    <p class="mb-0">{{ review.comment }}</p>
  </BaseCard>
</template>

<script setup>
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseBadge from '@/components/atoms/BaseBadge.vue'
import { formatDate } from '@/utils/formatters'

defineProps({
  review: {
    type: Object,
    required: true
  }
})
</script>
