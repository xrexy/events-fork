<template>
  <svg
    :width="props.scale.width"
    :height="props.scale.height"
    :class="props.classes.svg"
    viewBox="0 0 32 32"
  >
    <path
      :class="props.classes.path"
      d="M16 22L6 12l1.4-1.4l8.6 8.6l8.6-8.6L26 12z"
    ></path>
  </svg>
</template>

<script setup lang="ts">
const props = defineProps({
  classes: {
    type: Object,
    default: () => ({
      path: "",
      svg: "",
    }),
  },
  scale: {
    type: Object,
    default: () => ({
      width: "1em",
      height: "1em",
    }),
  },
});
</script>
