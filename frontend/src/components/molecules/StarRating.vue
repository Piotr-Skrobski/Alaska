<template>
  <div class="star-rating">
    <i
      v-for="n in maxRating"
      :key="n"
      class="fa-star"
      :class="n <= rating ? 'fas text-warning' : 'far text-muted'"
      :style="{ cursor: interactive ? 'pointer' : 'default', fontSize: size, marginRight: '4px' }"
      @click="interactive && setRating(n)"
      role="button"
      :aria-label="`Rate ${n} out of ${maxRating}`"
    ></i>
    <span v-if="showValue" class="ms-2">{{ rating }}/{{ maxRating }}</span>
  </div>
</template>

<script setup>
import { RATING } from '@/utils/constants'

const props = defineProps({
  rating: {
    type: Number,
    default: RATING.DEFAULT,
    validator: (value) => value >= RATING.MIN && value <= RATING.MAX
  },
  maxRating: {
    type: Number,
    default: RATING.MAX
  },
  interactive: {
    type: Boolean,
    default: true
  },
  size: {
    type: String,
    default: '1.5rem'
  },
  showValue: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:rating'])

const setRating = (value) => {
  if (props.interactive) {
    emit('update:rating', value)
  }
}
</script>
