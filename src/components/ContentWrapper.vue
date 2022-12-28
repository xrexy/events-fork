<script setup lang="ts">
const props = defineProps({
  additionalClasses: {
    type: String,
    default: "",
  },
});
</script>

<template>
  <div
    class="flex max-h-full flex-col justify-start gap-2 overflow-hidden text-gray-800 dark:text-gray-100"
    :class="props.additionalClasses"
  >
    <div class="text-2xl font-bold"><slot name="title"></slot></div>
    <div class="h-full"><slot></slot></div>
  </div>
</template>
