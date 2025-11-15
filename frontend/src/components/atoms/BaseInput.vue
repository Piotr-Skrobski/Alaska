<template>
  <input
    :type="type"
    :value="modelValue"
    :class="inputClasses"
    :placeholder="placeholder"
    :disabled="disabled"
    :required="required"
    @input="$emit('update:modelValue', $event.target.value)"
    v-bind="$attrs"
  />
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: ''
  },
  type: {
    type: String,
    default: 'text'
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
  error: {
    type: Boolean,
    default: false
  },
  size: {
    type: String,
    default: null,
    validator: (value) => !value || ['sm', 'lg'].includes(value)
  }
})

defineEmits(['update:modelValue'])

const inputClasses = computed(() => {
  const classes = ['form-control']

  if (props.error) {
    classes.push('is-invalid')
  }

  if (props.size) {
    classes.push(`form-control-${props.size}`)
  }

  return classes
})
</script>
