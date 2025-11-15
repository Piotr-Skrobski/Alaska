<template>
  <BaseAlert v-if="message" variant="success" :dismissible="dismissible" v-model="isVisible">
    {{ message }}
  </BaseAlert>
</template>

<script setup>
import { ref, watch } from 'vue'
import BaseAlert from '@/components/atoms/BaseAlert.vue'

const props = defineProps({
  message: {
    type: String,
    default: null
  },
  dismissible: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['dismiss'])

const isVisible = ref(true)

watch(isVisible, (newVal) => {
  if (!newVal) {
    emit('dismiss')
  }
})

watch(() => props.message, () => {
  isVisible.value = true
})
</script>
