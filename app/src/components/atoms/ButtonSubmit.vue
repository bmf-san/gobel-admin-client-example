<template>
  <input class="submit-button" type="submit" :value="value" />
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "ButtonSubmit",
  props: {
    value: {
      type: String,
      required: true,
      default: "",
    },
  },
});
</script>
