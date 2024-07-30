<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// Khai báo state để lưu trữ dữ liệu
const apiData = ref(null)
const error = ref(null)

// Gọi API khi component được mount
onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:3000/api/v1/applicant/1')
    apiData.value = response.data
  } catch (err) {
    error.value = 'Failed to fetch data'
    console.error(err)
  }
})
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
  </header>

  <main>
    <!-- Hiển thị dữ liệu từ API -->
    <div v-if="error">{{ error }}</div>
    <div v-else-if="apiData">
      <pre>{{ apiData }}</pre>
    </div>
    <div v-else>Loading...</div>
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }
}
</style>
