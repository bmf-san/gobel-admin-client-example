<template>
  <div
    class="alert alert-danger margin-bottom-2rem color-text-reverse"
    v-if="error"
  >
    <p>{{ error }}</p>
  </div>
</template>

<script>
export default {
  name: "Error",
  props: {
    error: String
  }
};
</script>
