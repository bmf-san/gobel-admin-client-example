<template>
  <input v-model="textComputed" type="text" :name="name" />
</template>

<script>
import { computed, defineComponent, toRefs } from "vue";

export default defineComponent({
  name: "InputText",
  props: {
    text: {
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
  emits: ["update:text"],
  setup(props, { emit }) {
    const { text } = toRefs(props);
    const textComputed = computed({
      get: () => text.value,
      set: value => emit("update:text", value)
    });
    return {
      textComputed
    };
  }
});
</script>
