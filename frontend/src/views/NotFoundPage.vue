<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6 text-center">
        <BaseCard shadow>
          <div class="py-5">
            <h1 class="display-1 text-muted">404</h1>
            <h2 class="mb-4">Page Not Found</h2>
            <p class="text-muted mb-4">
              The page you're looking for doesn't exist or has been moved.
            </p>
            <BaseButton variant="primary" @click="$router.push('/')">
              Go to Home
            </BaseButton>
          </div>
        </BaseCard>
      </div>
    </div>
  </div>
</template>

<script setup>
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'
</script>
