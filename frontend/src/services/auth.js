// services/auth.js
import axios from 'axios'

const API_URL = 'http://localhost:8080'

export default {
  setUser(userData) {
    localStorage.setItem('user', JSON.stringify(userData))
  },

  getUser() {
    const userStr = localStorage.getItem('user')
    return userStr ? JSON.parse(userStr) : null
  },

  isLoggedIn() {
    return !!localStorage.getItem('user')
  },

  clearUser() {
    localStorage.removeItem('user')
  },

  async login(credentials) {
    const response = await axios.post(`${API_URL}/users/login`, credentials, {
      withCredentials: true
    })
    this.setUser(response.data)
    return response.data
  },

  async logout() {
    await axios.post(`${API_URL}/users/logout`, {}, {
      withCredentials: true
    })
    this.clearUser()
  },

  async register(userData) {
    return axios.post(`${API_URL}/users/register`, userData, {
      withCredentials: true
    })
  },

  async deleteAccount() {
    await axios.post(`${API_URL}/users/delete`, {}, {
      withCredentials: true
    })
    this.clearUser()
  },

  async getUserProfile() {
    const response = await axios.get(`${API_URL}/users/me`, {
      withCredentials: true
    })
    return response.data
  }
}