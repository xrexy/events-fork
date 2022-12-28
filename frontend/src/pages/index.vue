<script setup lang="ts">
import Wrapper from "@/components/ContentWrapper.vue";

const events: { title: string; description: string }[] = [
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
  {
    description: "Some description",
    title: "Some title",
  },
];
</script>

<template>
  <div class="flex h-full flex-col">
    <Wrapper additional-classes="h-2/4">
      <template #title>Events</template>
      <div class="grid grid-flow-col grid-rows-3">
        <div
          v-for="event in events"
          :key="event.title"
          class="border bg-slate-900 p-2 dark:border-slate-800"
        >
          <div class="w-fit text-2xl font-bold">{{ event.title }}</div>
          <div class="w-fit text-gray-500">{{ event.description }}</div>
        </div>
      </div>
    </Wrapper>
    <Wrapper additional-classes="h-2/4">
      <template #title>Ideas</template>
      <h1 class="text-red-500">Test tailwind...</h1>
    </Wrapper>
  </div>
</template>
