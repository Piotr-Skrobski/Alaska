<template>
  <span :class="badgeClasses">
    <slot />
  </span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary'
  },
  pill: {
    type: Boolean,
    default: false
  }
})

const badgeClasses = computed(() => {
  const classes = ['badge', `bg-${props.variant}`]

  if (props.pill) {
    classes.push('rounded-pill')
  }

  return classes
})
</script>
