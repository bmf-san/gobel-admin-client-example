<template>
  <input v-model="emailComputed" type="email" :name="name" />
</template>

<script>
import { computed, defineComponent, toRefs } from "vue";

export default defineComponent({
  name: "InputEmail",
  props: {
    email: {
      type: String,
      required: true,
      default: ""
    },
    name: {
      type: String,
      required: true,
      default: ""
    }
  },
  emits: ["update:email"],
  setup(props, { emit }) {
    const { email } = toRefs(props);
    const emailComputed = computed({
      get: () => email.value,
      set: value => emit("update:email", value)
    });
    return {
      emailComputed
    };
  }
});
</script>
