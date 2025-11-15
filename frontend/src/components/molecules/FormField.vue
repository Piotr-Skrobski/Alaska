<template>
  <div class="mb-3">
    <BaseLabel :forId="id" :required="required">{{ label }}</BaseLabel>
    <component
      :is="inputComponent"
      :id="id"
      :modelValue="modelValue"
      :type="type"
      :error="!!error"
      :placeholder="placeholder"
      :disabled="disabled"
      :required="required"
      :rows="rows"
      @update:modelValue="$emit('update:modelValue', $event)"
    />
    <div v-if="error" class="invalid-feedback d-block">
      {{ error }}
    </div>
    <small v-if="helpText" class="form-text text-muted">
      {{ helpText }}
    </small>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import BaseLabel from '@/components/atoms/BaseLabel.vue'
import BaseInput from '@/components/atoms/BaseInput.vue'
import BaseTextarea from '@/components/atoms/BaseTextarea.vue'

const props = defineProps({
  id: {
    type: String,
    required: true
  },
  label: {
    type: String,
    required: true
  },
  modelValue: {
    type: [String, Number],
    default: ''
  },
  type: {
    type: String,
    default: 'text'
  },
  error: {
    type: String,
    default: null
  },
  placeholder: {
    type: String,
    default: ''
  },
  disabled: {
    type: Boolean,
    default: false
  },
  required: {
    type: Boolean,
    default: false
  },
  helpText: {
    type: String,
    default: null
  },
  rows: {
    type: Number,
    default: 3
  },
  multiline: {
    type: Boolean,
    default: false
  }
})

defineEmits(['update:modelValue'])

const inputComponent = computed(() => {
  return props.multiline ? BaseTextarea : BaseInput
})
</script>
