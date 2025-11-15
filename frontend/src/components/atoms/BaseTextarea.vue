<template>
  <textarea
    :value="modelValue"
    :class="textareaClasses"
    :placeholder="placeholder"
    :disabled="disabled"
    :required="required"
    :rows="rows"
    @input="$emit('update:modelValue', $event.target.value)"
    v-bind="$attrs"
  ></textarea>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
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
  rows: {
    type: Number,
    default: 3
  }
})

defineEmits(['update:modelValue'])

const textareaClasses = computed(() => {
  const classes = ['form-control']

  if (props.error) {
    classes.push('is-invalid')
  }

  return classes
})
</script>
