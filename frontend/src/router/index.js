import { createRouter, createWebHistory } from 'vue-router'
import Register from '@/views/Register.vue'
import Login from '@/views/Login.vue'
import MoviePage from '@/views/MoviePage.vue'
import Profile from '@/views/Profile.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/register',
      name: 'home',
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
      name: 'moviedetails',
      component: MoviePage
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile
    }
  ],
})

export default router
