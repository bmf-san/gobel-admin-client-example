<template>
  <input
    v-model="passwordComputed"
    type="password"
    :name="name"
    :autocomplete="autocomplete ? 'on' : 'off'"
  />
</template>

<script>
import { computed, defineComponent, toRefs } from "vue";

export default defineComponent({
  name: "InputPassword",
  props: {
    password: {
      type: String,
      required: true,
      default: ""
    },
    name: {
      type: String,
      required: true,
      default: ""
    },
    autocomplete: {
      type: Boolean,
      default: false
    }
  },
  emits: ["update:password"],
  setup(props, { emit }) {
    const { password } = toRefs(props);
    const passwordComputed = computed({
      get: () => password.value,
      set: value => emit("update:password", value)
    });
    return {
      passwordComputed
    };
  }
});
</script>
