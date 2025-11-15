<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="text-center mb-4">
          <h1>Join Alaska!</h1>
          <p class="text-muted">Create an account to review your favorite movies</p>
        </div>
        <TheRegisterForm />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import TheRegisterForm from '@/components/organisms/TheRegisterForm.vue'

const router = useRouter()
const authStore = useAuthStore()

onBeforeMount(() => {
  if (authStore.isLoggedIn) {
    router.push('/')
  }
})
</script>
