<script setup lang="ts">
import { useEvent } from "@/composables/use-events";
import { useAuth } from "@/stores/auth";
import { computed, onBeforeMount } from "vue";
import { useRouter } from "vue-router";

const { fetch, record: event } = useEvent();
const { isAuthenticated } = useAuth();
const router = useRouter();

const fail = (path: string, reason: string) => {
  console.warn("Mount failed: " + reason);
  router.push(path);
};

const eventId = computed(() => {
  const id = router.currentRoute.value.params.id ?? "";
  return typeof id === "string" ? id : id[0];
});

onBeforeMount(async () => {
  console.log("Mounting event view page");

  if (!isAuthenticated) return fail("/auth/login", "Not authenticated");

  if (!eventId.value) return fail("/", "No event id");

  await fetch(eventId.value)
    .catch(() => {})
    .then((success) => {
      if (!success) return fail("/", "Failed to fetch event");
    });

  if (!event.value) return fail("/", "No event found");
});
</script>

<template>
  <div v-if="!event"></div>
  <!-- event wont reach here if undefined, it's just to make typescript shut up -->
  <div v-else>
    <h1>Event View page {{ event.title }}</h1>
  </div>
</template>
