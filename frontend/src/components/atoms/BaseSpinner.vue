<template>
  <div :class="spinnerClasses" role="status">
    <span class="visually-hidden">{{ text }}</span>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  size: {
    type: String,
    default: null,
    validator: (value) => !value || ['sm'].includes(value)
  },
  variant: {
    type: String,
    default: 'primary'
  },
  text: {
    type: String,
    default: 'Loading...'
  }
})

const spinnerClasses = computed(() => {
  const classes = ['spinner-border']

  if (props.size) {
    classes.push(`spinner-border-${props.size}`)
  }

  if (props.variant) {
    classes.push(`text-${props.variant}`)
  }

  return classes
})
</script>
