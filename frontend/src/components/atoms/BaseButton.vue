<template>
  <button
    :class="buttonClasses"
    :disabled="disabled || loading"
    :type="type"
    v-bind="$attrs"
  >
    <span v-if="loading" class="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
    <slot />
  </button>
</template>

<script setup>
import { computed } from 'vue'
import { BUTTON_VARIANTS } from '@/utils/constants'

const props = defineProps({
  variant: {
    type: String,
    default: BUTTON_VARIANTS.PRIMARY,
    validator: (value) => Object.values(BUTTON_VARIANTS).includes(value)
  },
  size: {
    type: String,
    default: null,
    validator: (value) => !value || ['sm', 'lg'].includes(value)
  },
  loading: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  },
  type: {
    type: String,
    default: 'button'
  },
  outline: {
    type: Boolean,
    default: false
  },
  block: {
    type: Boolean,
    default: false
  }
})

const buttonClasses = computed(() => {
  const classes = ['btn']

  if (props.outline) {
    classes.push(`btn-outline-${props.variant}`)
  } else {
    classes.push(`btn-${props.variant}`)
  }

  if (props.size) {
    classes.push(`btn-${props.size}`)
  }

  if (props.block) {
    classes.push('w-100')
  }

  return classes
})
</script>
