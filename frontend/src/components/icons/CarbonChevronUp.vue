<template>
  <svg
    :width="props.scale.width"
    :height="props.scale.height"
    :class="props.classes.svg"
    viewBox="0 0 32 32"
  >
    <path
      :class="props.classes.path"
      d="m16 10l10 10l-1.4 1.4l-8.6-8.6l-8.6 8.6L6 20z"
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
