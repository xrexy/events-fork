<template>
  <svg
    :width="props.scale.width"
    :height="props.scale.height"
    :class="props.classes.svg"
    viewBox="0 0 24 24"
  >
    <path
      :class="props.classes.path"
      d="M11.38 2.019a7.5 7.5 0 1 0 10.6 10.6C21.662 17.854 17.316 22 12.001 22C6.477 22 2 17.523 2 12c0-5.315 4.146-9.661 9.38-9.981z"
    ></path>
  </svg>
</template>

<script>
export default {
  name: "RiMoonFill",
};
</script>

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
