<template>
  <BaseCard shadow>
    <h2 class="card-title text-center mb-4">Create an account</h2>

    <form @submit.prevent="handleRegister">
      <FormField
        id="username"
        label="Username"
        v-model="form.username"
        type="text"
        :error="errors.username"
        required
        placeholder="Choose a username"
        help-text="At least 3 characters"
      />

      <FormField
        id="email"
        label="Email"
        v-model="form.email"
        type="email"
        :error="errors.email"
        required
        placeholder="Enter your email"
      />

      <FormField
        id="password"
        label="Password"
        v-model="form.password"
        type="password"
        :error="errors.password"
        required
        placeholder="Choose a password"
        help-text="At least 6 characters"
      />

      <ErrorAlert :message="errorMessage" @dismiss="errorMessage = null" />

      <BaseButton variant="success" type="submit" block :loading="isSubmitting">
        {{ isSubmitting ? 'Creating account...' : 'Register' }}
      </BaseButton>

      <p class="text-center mt-3">
        Already have an account?
        <router-link to="/login" class="text-decoration-none">Log in</router-link>
      </p>
    </form>
  </BaseCard>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { validateEmail, validateUsername, validatePassword } from '@/utils/validators'
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'
import FormField from '@/components/molecules/FormField.vue'
import ErrorAlert from '@/components/molecules/ErrorAlert.vue'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({
  username: '',
  email: '',
  password: ''
})

const errors = ref({
  username: '',
  email: '',
  password: ''
})

const errorMessage = ref(null)
const isSubmitting = ref(false)

const validateForm = () => {
  errors.value = {
    username: '',
    email: '',
    password: ''
  }

  let isValid = true

  if (!validateUsername(form.value.username)) {
    errors.value.username = 'Username must be at least 3 characters'
    isValid = false
  }

  if (!validateEmail(form.value.email)) {
    errors.value.email = 'Please enter a valid email address'
    isValid = false
  }

  if (!validatePassword(form.value.password)) {
    errors.value.password = 'Password must be at least 6 characters'
    isValid = false
  }

  return isValid
}

const handleRegister = async () => {
  if (!validateForm()) {
    return
  }

  isSubmitting.value = true
  errorMessage.value = null

  try {
    await authStore.register(form.value)

    router.push('/login?registered=true')
  } catch (error) {
    errorMessage.value = error.response?.data || 'An error occurred during registration. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.card {
  animation: fadeIn 0.8s ease-out both;
}
</style>
