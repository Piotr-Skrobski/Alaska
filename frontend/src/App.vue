<script setup>
import { RouterLink, RouterView } from 'vue-router'
import AppHeader from './components/AppHeader.vue';
</script>

<template>
  <AppHeader />
  <RouterView />
</template>

<style>
body {
  background-color: bisque;
}
</style>