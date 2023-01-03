<script setup lang="ts">
import { useEvent } from "@/composables/use-events";
import { useRouter } from "vue-router";

const router = useRouter();
const { fetchAll, createThumbnailURI } = useEvent();

const events = await fetchAll().then((records) =>
  records.map((record) => ({
    ...record,
    thumbnail: createThumbnailURI(record.id, record.thumbnail),
  }))
);
</script>

<template>
  <div v-if="events.length > 0">
    <div
      v-for="event in events"
      :key="event.title"
      @click="() => router.push(`/events/${event.id}`)"
      class="w-fit border bg-slate-900 p-2 dark:border-slate-800"
    >
      <img
        :src="event.thumbnail"
        class="aspect-square h-auto w-32"
        :alt="`imgage for ${event.title}`"
      />
      <div class="w-fit text-gray-500">{{ event.description }}</div>
    </div>
  </div>
  <div v-else>
    <h1>No events found :(</h1>
  </div>
</template>
