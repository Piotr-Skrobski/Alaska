<template>
  <div class="input-group">
    <BaseInput
      v-model="query"
      :placeholder="placeholder"
      @keyup.enter="handleSearch"
    />
    <BaseButton variant="primary" @click="handleSearch" :loading="loading">
      {{ buttonText }}
    </BaseButton>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import BaseInput from '@/components/atoms/BaseInput.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'Search...'
  },
  buttonText: {
    type: String,
    default: 'Search'
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'search'])

const query = ref(props.modelValue)

watch(() => props.modelValue, (newVal) => {
  query.value = newVal
})

watch(query, (newVal) => {
  emit('update:modelValue', newVal)
})

const handleSearch = () => {
  if (query.value.trim()) {
    emit('search', query.value.trim())
  }
}
</script>
