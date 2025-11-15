<template>
  <div :class="cardClasses">
    <div v-if="$slots.header || headerText" class="card-header" :class="headerClass">
      <slot name="header">{{ headerText }}</slot>
    </div>
    <div class="card-body">
      <slot />
    </div>
    <div v-if="$slots.footer" class="card-footer" :class="footerClass">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  headerText: {
    type: String,
    default: null
  },
  headerClass: {
    type: String,
    default: 'bg-primary text-white'
  },
  footerClass: {
    type: String,
    default: ''
  },
  shadow: {
    type: Boolean,
    default: false
  }
})

const cardClasses = computed(() => {
  const classes = ['card', 'mb-4']

  if (props.shadow) {
    classes.push('shadow-sm')
  }

  return classes
})
</script>
