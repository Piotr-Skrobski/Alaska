<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-7">
        <div class="card shadow-sm animate__animated animate__fadeIn">
          <div class="card-body">
            <h2 class="card-title text-center mb-4">My Profile</h2>

            <div v-if="loading" class="text-center my-4">
              <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
              </div>
            </div>

            <div v-if="errorMessage" class="alert alert-danger text-center">
              {{ errorMessage }}
            </div>

            <div v-if="user && !loading">
              <ul class="list-group mb-4">
                <li class="list-group-item">
                  <strong>ID:</strong> {{ user.user_id }}
                </li>
              </ul>

              <div class="d-grid">
                <button class="btn btn-danger" @click="confirmDelete" :disabled="deleting">
                  <span v-if="deleting" class="spinner-border spinner-border-sm me-2" role="status"
                    aria-hidden="true"></span>
                  {{ deleting ? 'Deleting...' : 'Delete my account' }}
                </button>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

const API_URL = 'http://localhost:8080';

export default {
  name: 'ProfilePage',
  data() {
    return {
      user: null,
      loading: false,
      deleting: false,
      errorMessage: ''
    };
  },
  created() {
    this.fetchProfile();
  },
  methods: {
    fetchProfile() {
      this.loading = true;
      axios.get(`${API_URL}/users/me`, {
        withCredentials: true
      })
        .then(response => {
          console.log(response);
          this.user = response.data;
        })
        .catch(error => {
          this.errorMessage = error.response?.data || 'Failed to load profile.';
        })
        .finally(() => {
          this.loading = false;
        });
    },
    confirmDelete() {
      if (!confirm('Are you sure you want to delete your account? This action cannot be undone.')) {
        return;
      }

      this.deleting = true;
      axios.post(`${API_URL}/users/delete`, {}, {
        withCredentials: true
      })
        .then(() => {
          alert('Your account has been deleted.');
          this.$router.push('/register');
        })
        .catch(error => {
          this.errorMessage = error.response?.data || 'Failed to delete account.';
        })
        .finally(() => {
          this.deleting = false;
        });
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
