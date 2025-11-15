<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="text-center mb-4">
          <h1>Welcome back!</h1>
          <p class="text-muted">Log in to continue reviewing movies</p>
        </div>
        <TheLoginForm />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import TheLoginForm from '@/components/organisms/TheLoginForm.vue'

const router = useRouter()
const authStore = useAuthStore()

onBeforeMount(() => {
  if (authStore.isLoggedIn) {
    router.push('/')
  }
})
</script>
