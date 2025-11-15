<template>
  <BaseCard shadow>
    <h2 class="card-title text-center mb-4">Log In</h2>

    <SuccessAlert v-if="registrationSuccess" message="Registration successful! You can now log in." />

    <form @submit.prevent="handleLogin">
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
        placeholder="Enter your password"
      />

      <ErrorAlert :message="errorMessage" @dismiss="errorMessage = null" />

      <BaseButton variant="primary" type="submit" block :loading="isSubmitting">
        {{ isSubmitting ? 'Logging in...' : 'Log In' }}
      </BaseButton>

      <p class="text-center mt-3">
        Don't have an account? <router-link to="/register" class="text-success">Register</router-link>
      </p>
    </form>
  </BaseCard>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { validateEmail } from '@/utils/validators'
import BaseCard from '@/components/atoms/BaseCard.vue'
import BaseButton from '@/components/atoms/BaseButton.vue'
import FormField from '@/components/molecules/FormField.vue'
import ErrorAlert from '@/components/molecules/ErrorAlert.vue'
import SuccessAlert from '@/components/molecules/SuccessAlert.vue'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const form = ref({
  email: '',
  password: ''
})

const errors = ref({
  email: '',
  password: ''
})

const errorMessage = ref(null)
const isSubmitting = ref(false)
const registrationSuccess = ref(false)

onMounted(() => {
  registrationSuccess.value = route.query.registered === 'true'
})

const validateForm = () => {
  errors.value = {
    email: '',
    password: ''
  }

  let isValid = true

  if (!validateEmail(form.value.email)) {
    errors.value.email = 'Please enter a valid email address'
    isValid = false
  }

  if (!form.value.password) {
    errors.value.password = 'Password is required'
    isValid = false
  }

  return isValid
}

const handleLogin = async () => {
  if (!validateForm()) {
    return
  }

  isSubmitting.value = true
  errorMessage.value = null

  try {
    await authStore.login(form.value)

    const redirectTo = route.query.redirect || '/'
    router.push(redirectTo)
  } catch (error) {
    if (error.response?.status === 401) {
      errorMessage.value = 'Invalid email or password'
    } else {
      errorMessage.value = error.response?.data || 'An error occurred. Please try again.'
    }
  } finally {
    isSubmitting.value = false
  }
}
</script>
