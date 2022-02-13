<template>
  <div
    v-if="error"
    class="alert alert-danger margin-bottom-2rem color-text-reverse"
  >
    <p>{{ error }}</p>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "Error",
  props: {
    error: {
      type: String,
      required: false,
      default: ""
    }
  }
});
</script>
