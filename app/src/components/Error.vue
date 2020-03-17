<template>
  <div class="error" v-if="error">
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

<style scoped>
.error p {
  color: red;
}
</style>
