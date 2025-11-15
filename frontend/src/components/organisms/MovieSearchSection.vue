<template>
  <BaseCard headerText="Find a movie">
    <div class="row">
      <div class="col-md-8">
        <SearchBar
          v-model="searchQuery"
          placeholder="Enter movie title"
          button-text="Search"
          :loading="loading"
          @search="handleSearch"
        />
      </div>
    </div>
  </BaseCard>
</template>

<script setup>
import { ref } from 'vue'
import BaseCard from '@/components/atoms/BaseCard.vue'
import SearchBar from '@/components/molecules/SearchBar.vue'

const props = defineProps({
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['search'])

const searchQuery = ref('')

const handleSearch = (query) => {
  emit('search', query)
}
</script>
