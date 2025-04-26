<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-9">
        <div class="card shadow-sm animate__animated animate__fadeIn">
          <div class="card-body">
            <h2 class="card-title text-center mb-4">Create an account</h2>
            <form @submit.prevent="register">
              <div class="mb-3">
                <label for="username" class="form-label">Username</label>
                <input type="text" id="username" v-model="form.username"
                  :class="['form-control', { 'is-invalid': errors.username }]" required />
                <div v-if="errors.username" class="invalid-feedback">
                  {{ errors.username }}
                </div>
              </div>

              <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" id="email" v-model="form.email"
                  :class="['form-control', { 'is-invalid': errors.email }]" required />
                <div v-if="errors.email" class="invalid-feedback">
                  {{ errors.email }}
                </div>
              </div>

              <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input type="password" id="password" v-model="form.password"
                  :class="['form-control', { 'is-invalid': errors.password }]" required />
                <div v-if="errors.password" class="invalid-feedback">
                  {{ errors.password }}
                </div>
              </div>

              <div v-if="errorMessage" class="alert alert-danger text-center">
                {{ errorMessage }}
              </div>

              <div class="d-grid mb-3">
                <button type="submit" class="btn btn-success" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2" role="status"
                    aria-hidden="true"></span>
                  {{ isSubmitting ? 'Creating account...' : 'Register' }}
                </button>
              </div>

              <p class="text-center">
                Already have an account?
                <router-link to="/login" class="text-decoration-none">Log in</router-link>
              </p>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import axios from 'axios'

export default {
  name: 'RegisterForm',
  data() {
    return {
      form: {
        username: '',
        email: '',
        password: ''
      },
      errors: {
        username: '',
        email: '',
        password: ''
      },
      errorMessage: '',
      isSubmitting: false
    }
  },
  methods: {
    validateForm() {
      let isValid = true
      this.errors = {
        username: '',
        email: '',
        password: ''
      }

      if (!this.form.username || this.form.username.length < 3) {
        this.errors.username = 'Username must be at least 3 characters'
        isValid = false
      }

      if (!this.form.email || !this.form.email.includes('@')) {
        this.errors.email = 'Please enter a valid email address'
        isValid = false
      }

      if (!this.form.password || this.form.password.length < 6) {
        this.errors.password = 'Password must be at least 6 characters'
        isValid = false
      }

      return isValid
    },
    async register() {
      if (!this.validateForm()) {
        return
      }

      this.isSubmitting = true
      this.errorMessage = ''

      try {
        await axios.post('http://localhost:8080/users/register', this.form, {
          withCredentials: true
        })

        this.$router.push('/login?registered=true')
      } catch (error) {
        console.error('Registration error:', error)
        if (error.response && error.response.data) {
          this.errorMessage = error.response.data
        } else {
          this.errorMessage = 'An error occurred during registration. Please try again.'
        }
      } finally {
        this.isSubmitting = false
      }
    }
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

.animate__animated.animate__fadeIn {
  animation: fadeIn 0.8s ease-out both;
}
</style>