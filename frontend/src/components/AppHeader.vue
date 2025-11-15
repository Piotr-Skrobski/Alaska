<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-white shadow-sm sticky-top">
    <div class="container">
      <router-link class="navbar-brand" to="/">
        <h1 class="h4 mb-0 fw-bold">Alaska: The Movie Reviews</h1>
      </router-link>

      <button class="navbar-toggler" type="button" @click="isMobileMenuOpen = !isMobileMenuOpen"
        :class="{ 'collapsed': !isMobileMenuOpen }">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" :class="{ 'show': isMobileMenuOpen }">
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link class="nav-link" to="/">Home</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/movies">Movies</router-link>
          </li>

          <template v-if="!authStore.isLoggedIn">
            <li class="nav-item">
              <router-link class="nav-link" to="/login">Log In</router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link btn btn-primary text-white px-3 mx-lg-2"
                to="/register">Register</router-link>
            </li>
          </template>

          <template v-else>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" href="#" role="button"
                @click.prevent="isDropdownOpen = !isDropdownOpen" aria-expanded="false">
                {{ authStore.username }}
              </a>
              <ul class="dropdown-menu dropdown-menu-end" :class="{ 'show': isDropdownOpen }">
                <li><router-link class="dropdown-item" to="/profile">My Profile</router-link></li>
                <li>
                  <hr class="dropdown-divider">
                </li>
                <li><a class="dropdown-item" href="#" @click.prevent="handleLogout">Log Out</a></li>
              </ul>
            </li>
          </template>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const isMobileMenuOpen = ref(false)
const isDropdownOpen = ref(false)

const handleLogout = async () => {
  try {
    await authStore.logout()
    isMobileMenuOpen.value = false
    isDropdownOpen.value = false

    if (route.path !== '/') {
      router.push('/')
    } else {
      router.go()
    }
  } catch (error) {
    console.error('Logout error:', error)
    router.go()
  }
}

const handleClickOutside = (event) => {
  if (isDropdownOpen.value && !event.target.closest('.dropdown')) {
    isDropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.navbar {
  min-height: 70px;
}

@media (max-width: 992px) {
  .dropdown-menu {
    position: static !important;
    border: none;
    box-shadow: none;
    padding-left: 1.5rem;
  }

  .dropdown-divider {
    margin-left: -1.5rem;
    margin-right: -1.5rem;
  }
}
</style>