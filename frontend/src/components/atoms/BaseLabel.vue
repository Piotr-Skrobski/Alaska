<template>
  <label :for="forId" class="form-label">
    <slot />
    <span v-if="required" class="text-danger ms-1">*</span>
  </label>
</template>

<script setup>
defineProps({
  forId: {
    type: String,
    default: null
  },
  required: {
    type: Boolean,
    default: false
  }
})
</script>
