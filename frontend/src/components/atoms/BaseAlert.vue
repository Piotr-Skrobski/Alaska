<template>
  <div v-if="modelValue" :class="alertClasses" role="alert">
    <div class="d-flex justify-content-between align-items-start">
      <div class="flex-grow-1">
        <slot />
      </div>
      <button
        v-if="dismissible"
        type="button"
        class="btn-close"
        @click="$emit('update:modelValue', false)"
        aria-label="Close"
      ></button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ALERT_VARIANTS } from '@/utils/constants'

const props = defineProps({
  variant: {
    type: String,
    default: ALERT_VARIANTS.INFO,
    validator: (value) => Object.values(ALERT_VARIANTS).includes(value)
  },
  dismissible: {
    type: Boolean,
    default: false
  },
  modelValue: {
    type: Boolean,
    default: true
  }
})

defineEmits(['update:modelValue'])

const alertClasses = computed(() => {
  const classes = ['alert', `alert-${props.variant}`]

  if (props.dismissible) {
    classes.push('alert-dismissible')
  }

  return classes
})
</script>
