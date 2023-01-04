<script setup lang="ts">
import EventsGrid from "@/components/EventsGrid.vue";
import { useAuth } from "@/stores/auth";

const { isAuthenticated } = useAuth();
</script>

<template>
  <div class="flex h-full flex-col">
    <div class="flex w-full flex-col">
      <p class="text-left text-2xl font-bold">Events</p>

      <Suspense v-if="isAuthenticated">
        <div class="flex flex-row flex-wrap">
          <EventsGrid />
        </div>

        <template #fallback>
          <h1>Loading....</h1>
        </template>
      </Suspense>
      <h1 v-else>Uhhh.. login first..?</h1>
    </div>
  </div>
</template>
