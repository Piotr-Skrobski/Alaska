<template>
  <div class="register-page">
    <h1>Join Alaska!</h1>
    <p>Create an account to review your favorite movies</p>
    <register-form />
  </div>
</template>

<script>
import RegisterForm from '@/components/RegisterForm.vue'

export default {
  name: 'RegisterPage',
  components: {
    RegisterForm
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
.register-page {
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