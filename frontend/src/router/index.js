import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import Register from '@/views/Register.vue'
import Login from '@/views/Login.vue'
import MoviePage from '@/views/MoviePage.vue'
import Profile from '@/views/Profile.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomePage,
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/movies',
      name: 'movies',
      component: MoviePage
    },
    {
      path: '/movies/:id',
      name: 'movie-details',
      component: MoviePage
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
      meta: { requiresAuth: true }
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/views/NotFoundPage.vue')
    }
  ],
})

// Navigation guard for protected routes
router.beforeEach((to, from, next) => {
  // Check if route requires authentication
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // Import the auth store dynamically to avoid circular dependency
    import('@/stores/auth').then(({ useAuthStore }) => {
      const authStore = useAuthStore()

      if (!authStore.isLoggedIn) {
        // Redirect to login page with the intended destination
        next({
          name: 'login',
          query: { redirect: to.fullPath }
        })
      } else {
        next()
      }
    })
  } else {
    next()
  }
})

export default router
