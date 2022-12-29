<script setup lang="ts">
import { useEvent } from "@/composables/use-events";
import { useAuthStore } from "@/stores/auth";

const { fetchAll, createThumbnailURI } = useEvent();
const { isAuthenticated, login } = useAuthStore();

if (!isAuthenticated) {
  await login({ identity: "xrexy_test_2", password: "testpass" });
}

const events = await fetchAll().then((records) =>
  records.map((record) => ({
    ...record,
    thumbnail: createThumbnailURI(record.id, record.thumbnail),
  }))
);

console.log(events);
</script>

<template>
  <div class="h-fit w-fit">
    <div
      v-for="event in events"
      :key="event.title"
      class="border bg-slate-900 p-2 dark:border-slate-800"
    >
      <div class="w-fit text-2xl font-bold">{{ event.title }}</div>
      <div class="w-fit text-gray-500">{{ event.description }}</div>
      <img class="h-40 w-40" :src="event.thumbnail" />
    </div>
  </div>
</template>
