<script setup lang="ts">
import Header from "./components/AppHeader.vue";
import NotificationsRoot from "./components/NotificationsRoot.vue";
</script>

<template>
  <Header />
  <main class="bg-gray-100 px-8 pt-[76px] dark:bg-gray-900">
    <router-view />
  </main>
  <NotificationsRoot />
</template>
