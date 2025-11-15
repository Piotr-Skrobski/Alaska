<template>
  <div class="container mt-4">
    <div class="jumbotron text-center py-5 mb-5">
      <h1 class="display-4">Welcome to Alaska</h1>
      <p class="lead">Your ultimate movie review platform</p>
      <hr class="my-4">
      <p>Discover, review, and share your thoughts on your favorite movies.</p>
      <BaseButton variant="primary" size="lg" @click="$router.push('/movies')">
        Start Exploring
      </BaseButton>
    </div>

    <div class="row">
      <div class="col-md-4 mb-4">
        <BaseCard shadow>
          <div class="text-center">
            <i class="fas fa-search fa-3x text-primary mb-3"></i>
            <h5>Search Movies</h5>
            <p>Find detailed information about any movie using our powerful search feature.</p>
          </div>
        </BaseCard>
      </div>

      <div class="col-md-4 mb-4">
        <BaseCard shadow>
          <div class="text-center">
            <i class="fas fa-star fa-3x text-warning mb-3"></i>
            <h5>Rate & Review</h5>
            <p>Share your opinions and rate movies from 1 to 10 stars.</p>
          </div>
        </BaseCard>
      </div>

      <div class="col-md-4 mb-4">
        <BaseCard shadow>
          <div class="text-center">
            <i class="fas fa-robot fa-3x text-success mb-3"></i>
            <h5>AI-Powered Reviews</h5>
            <p>Generate movie reviews using our integrated AI assistant.</p>
          </div>
        </BaseCard>
      </div>
    </div>

    <div class="text-center mt-5">
      <h3 class="mb-4">Get Started Today</h3>
      <template v-if="!authStore.isLoggedIn">
        <BaseButton variant="success" class="me-2" @click="$router.push('/register')">
          Create Account
        </BaseButton>
        <BaseButton variant="outline" :outline="true" @click="$router.push('/login')">
          Sign In
        </BaseButton>
      </template>
      <template v-else>
        <p class="text-muted">Welcome back, {{ authStore.username }}!</p>
        <BaseButton variant="primary" @click="$router.push('/profile')">
          Go to Profile
        </BaseButton>
      </template>
    </div>
  </div>
</template>

<script setup>
import { useAuthStore } from '@/stores/auth'
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'

const authStore = useAuthStore()
</script>

<style scoped>
.jumbotron {
  background-color: #f8f9fa;
  border-radius: 0.5rem;
}
</style>
