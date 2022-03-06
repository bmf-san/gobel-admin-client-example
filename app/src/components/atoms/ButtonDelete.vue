<template>
  <a
    class="color-danger delete-link"
    @click.prevent.stop="$emit('deleteItems')"
    >{{ text }}</a
  >
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "ButtonDelete",
  props: {
    text: {
      type: String,
      required: true,
      default: ""
    }
  },
  emits: ["deleteItems"]
});
</script>

<style scoped>
.delete-link:hover {
  color: var(--danger-color);
}
</style>
