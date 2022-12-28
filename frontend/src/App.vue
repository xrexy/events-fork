<script setup lang="ts">
import Header from "./components/AppHeader.vue";
</script>

<template>
  <Header />
  <main class="bg-gray-100 px-64 pt-[76px] dark:bg-gray-900">
    <router-view />
  </main>
</template>
