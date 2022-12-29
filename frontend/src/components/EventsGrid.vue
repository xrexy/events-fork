<script setup lang="ts">
import { useEvent } from "@/composables/use-events";
import { useAuthStore } from "@/stores/auth";

const { fetchAll, createThumbnailURI } = useEvent();
const { isAuthenticated, login } = useAuthStore();

if (!isAuthenticated) {
  await login({ identity: "xrexy_test", password: "testpass" });
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
  <div
    v-for="event in events"
    :key="event.title"
    class="w-fit border bg-slate-900 p-2 dark:border-slate-800"
  >
    <div>
      <img
        :src="event.thumbnail"
        class="aspect-square h-auto w-32"
        :alt="`img_${event.id}`"
      />
    </div>
    <div class="w-fit text-gray-500">{{ event.description }}</div>
  </div>
</template>
