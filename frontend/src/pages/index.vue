<script setup lang="ts">
import Wrapper from "@/components/ContentWrapper.vue";
import EventsGrid from "@/components/EventsGrid.vue";
</script>

<template>
  <div class="flex h-full flex-col">
    <Wrapper additional-classes="h-2/4">
      <template #title>Events</template>
      <Suspense>
        <template #default>
          <div class="grid grid-flow-col grid-rows-3">
            <EventsGrid />
          </div>
        </template>
        <template #fallback>
          <h1>Loading....</h1>
        </template>
      </Suspense>
    </Wrapper>
  </div>
</template>
