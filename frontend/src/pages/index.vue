<script setup lang="ts">
import EventsGrid from "@/components/EventsGrid.vue";
</script>

<template>
  <div class="flex h-full flex-col">
    <div class="flex w-full flex-col">
      <p class="text-left text-2xl font-bold">Events</p>
      <Suspense>
        <template #default>
          <div class="flex flex-row flex-wrap">
            <EventsGrid />
          </div>
        </template>
        <template #fallback>
          <h1>Loading....</h1>
        </template>
      </Suspense>
    </div>
  </div>
</template>
