<template>
  <div v-if="isOpen" class="modal d-block" tabindex="-1" style="background: rgba(0,0,0,0.5);">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">{{ title }}</h5>
          <button type="button" class="btn-close" @click="$emit('cancel')" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <p>{{ message }}</p>
        </div>
        <div class="modal-footer">
          <BaseButton variant="secondary" @click="$emit('cancel')">
            {{ cancelText }}
          </BaseButton>
          <BaseButton :variant="confirmVariant" @click="$emit('confirm')">
            {{ confirmText }}
          </BaseButton>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import BaseButton from '@/components/atoms/BaseButton.vue'

defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    required: true
  },
  message: {
    type: String,
    required: true
  },
  confirmText: {
    type: String,
    default: 'Confirm'
  },
  cancelText: {
    type: String,
    default: 'Cancel'
  },
  confirmVariant: {
    type: String,
    default: 'danger'
  }
})

defineEmits(['confirm', 'cancel'])
</script>
