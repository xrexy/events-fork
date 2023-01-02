<script setup lang="ts">
import { useEvent } from "@/composables/use-events";

const { fetchAll, createThumbnailURI } = useEvent();

const events = await fetchAll().then((records) =>
  records.map((record) => ({
    ...record,
    thumbnail: createThumbnailURI(record.id, record.thumbnail),
  }))
);
</script>

<template>
  <div
    v-for="event in events"
    :key="event.title"
    class="w-fit border bg-slate-900 p-2 dark:border-slate-800"
  >
    <div>
      <img
        :src="event.thumbnail"
        class="aspect-square h-auto w-32"
        :alt="`imgage for ${event.title}`"
      />
    </div>
    <div class="w-fit text-gray-500">{{ event.description }}</div>
  </div>
</template>
