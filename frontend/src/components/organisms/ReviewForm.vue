<template>
  <BaseCard headerText="Write a review">
    <form @submit.prevent="handleSubmit">
      <div class="mb-3">
        <BaseLabel>Rating (1-10)</BaseLabel>
        <StarRating
          :rating="form.rating"
          :max-rating="10"
          :interactive="true"
          :show-value="true"
          @update:rating="form.rating = $event"
        />
      </div>

      <div class="mb-3">
        <BaseLabel forId="comment">Your review</BaseLabel>
        <BaseTextarea
          id="comment"
          v-model="form.comment"
          :rows="4"
          placeholder="Share your thoughts about this movie..."
          required
        />
      </div>

      <div class="d-flex gap-2 flex-wrap">
        <BaseButton variant="primary" type="submit" :loading="submitting">
          Submit review
        </BaseButton>
        <BaseButton variant="secondary" @click="handleGenerateFromLLM" :loading="llmGenerating">
          Generate Review from LLM
        </BaseButton>
      </div>
    </form>
  </BaseCard>
</template>

<script setup>
import { ref, computed } from 'vue'
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseLabel from '@/components/atoms/BaseLabel.vue'
import BaseTextarea from '@/components/atoms/BaseTextarea.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'
import StarRating from '@/components/molecules/StarRating.vue'
import { RATING } from '@/utils/constants'
import reviewService from '@/services/reviewService'
import llmService from '@/services/llmService'
import { useAuthStore } from '@/stores/auth'
import { useUIStore } from '@/stores/ui'

const props = defineProps({
  movieId: {
    type: String,
    required: true
  },
  movieTitle: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['submitted'])

const authStore = useAuthStore()
const uiStore = useUIStore()

const form = ref({
  rating: RATING.DEFAULT,
  comment: ''
})

const submitting = ref(false)
const llmGenerating = ref(false)

const handleSubmit = async () => {
  if (!form.value.comment.trim()) {
    uiStore.showError('Please write a review comment')
    return
  }

  submitting.value = true

  try {
    const reviewData = {
      user_id: authStore.user.id,
      movie_id: props.movieId,
      rating: form.value.rating,
      comment: form.value.comment
    }

    await reviewService.create(reviewData)

    // Reset form
    form.value = {
      rating: RATING.DEFAULT,
      comment: ''
    }

    uiStore.showSuccess('Review submitted successfully!')
    emit('submitted')
  } catch (error) {
    uiStore.showError(error.response?.data || 'Error submitting review')
  } finally {
    submitting.value = false
  }
}

const handleGenerateFromLLM = async () => {
  llmGenerating.value = true

  try {
    const generatedText = await llmService.generateReview(props.movieTitle, form.value.rating)
    form.value.comment = generatedText
  } catch (error) {
    uiStore.showError('Error generating review from LLM')
  } finally {
    llmGenerating.value = false
  }
}
</script>
