<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-9">
        <div class="card shadow">
          <div class="card-body p-4">
            <h2 class="card-title text-center mb-4">Log In</h2>

            <div v-if="registrationSuccess" class="alert alert-success" role="alert">
              Registration successful! You can now log in.
            </div>

            <form @submit.prevent="login">
              <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" class="form-control" :class="{ 'is-invalid': errors.email }" id="email"
                  v-model="form.email" required>
                <div v-if="errors.email" class="invalid-feedback">{{ errors.email }}</div>
              </div>

              <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input type="password" class="form-control" :class="{ 'is-invalid': errors.password }" id="password"
                  v-model="form.password" required>
                <div v-if="errors.password" class="invalid-feedback">{{ errors.password }}</div>
              </div>

              <div v-if="errorMessage" class="alert alert-danger text-center py-2 mb-3" role="alert">
                {{ errorMessage }}
              </div>

              <div class="d-grid">
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2" role="status"
                    aria-hidden="true"></span>
                  {{ isSubmitting ? 'Logging in...' : 'Log In' }}
                </button>
              </div>

              <p class="text-center mt-3">
                Don't have an account? <router-link to="/register" class="text-success">Register</router-link>
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
  name: 'LoginForm',
  data() {
    return {
      form: {
        email: '',
        password: ''
      },
      errors: {
        email: '',
        password: ''
      },
      errorMessage: '',
      isSubmitting: false,
      registrationSuccess: false
    }
  },
  created() {
    this.registrationSuccess = this.$route.query.registered === 'true'
  },
  methods: {
    validateForm() {
      let isValid = true
      this.errors = {
        email: '',
        password: ''
      }

      if (!this.form.email || !this.form.email.includes('@')) {
        this.errors.email = 'Please enter a valid email address'
        isValid = false
      }

      if (!this.form.password) {
        this.errors.password = 'Password is required'
        isValid = false
      }

      return isValid
    },
    async login() {
      if (!this.validateForm()) {
        return
      }

      this.isSubmitting = true
      this.errorMessage = ''

      try {
        const response = await axios.post('http://localhost:8080/users/login', this.form, {
          withCredentials: true
        })

        localStorage.setItem('user', JSON.stringify(response.data))

        const redirectTo = this.$route.query.redirect || '/'
        this.$router.push(redirectTo)
      } catch (error) {
        console.error('Login error:', error)
        if (error.response && error.response.status === 401) {
          this.errorMessage = 'Invalid email or password'
        } else if (error.response && error.response.data) {
          this.errorMessage = error.response.data
        } else {
          this.errorMessage = 'An error occurred. Please try again.'
        }
      } finally {
        this.isSubmitting = false
      }
    }
  }
}
</script>