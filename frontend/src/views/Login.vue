<template>
  <div class="login-page">
    <h1>Welcome back!</h1>
    <p>Log in to continue reviewing movies</p>
    <login-form />
  </div>
</template>

<script>
import LoginForm from '@/components/LoginForm.vue'

export default {
  name: 'LoginPage',
  components: {
    LoginForm
  },
  beforeMount() {
    const user = localStorage.getItem('user')
    if (user) {
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>
.login-page {
  max-width: 600px;
  margin: 40px auto;
  text-align: center;
}

h1 {
  margin-bottom: 10px;
  color: #333;
}

p {
  color: #666;
  margin-bottom: 30px;
}
</style>