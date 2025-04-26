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

          <template v-if="!isLoggedIn">
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
                {{ username }}
              </a>
              <ul class="dropdown-menu dropdown-menu-end" :class="{ 'show': isDropdownOpen }">
                <li><router-link class="dropdown-item" to="/profile">My Profile</router-link></li>
                <li><router-link class="dropdown-item" to="/my-reviews">My Reviews</router-link></li>
                <li>
                  <hr class="dropdown-divider">
                </li>
                <li><a class="dropdown-item" href="#" @click.prevent="logout">Log Out</a></li>
              </ul>
            </li>
          </template>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
import auth from '@/services/auth'

export default {
  name: 'AppHeader',
  data() {
    return {
      isMobileMenuOpen: false,
      isDropdownOpen: false
    }
  },
  computed: {
    isLoggedIn() {
      return auth.isLoggedIn()
    },
    username() {
      const user = auth.getUser()
      return user ? user.username || 'User' : 'User'
    }
  },
  methods: {
    async logout() {
      try {
        await auth.logout()
        this.isMobileMenuOpen = false
        this.isDropdownOpen = false

        if (this.$route.path !== '/') {
          this.$router.push('/')
        } else {
          this.$router.go()
        }
      } catch (error) {
        console.error('Logout error:', error)
        auth.clearUser()
        this.$router.go()
      }
    },
    handleClickOutside(event) {
      if (this.isDropdownOpen && !event.target.closest('.dropdown')) {
        this.isDropdownOpen = false
      }
    }
  },
  mounted() {
    document.addEventListener('click', this.handleClickOutside)
  },
  beforeDestroy() {
    document.removeEventListener('click', this.handleClickOutside)
  }
}
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