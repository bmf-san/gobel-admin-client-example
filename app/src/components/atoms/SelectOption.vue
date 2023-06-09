<template>
  <select v-model="selectedComputed">
    <option disabled value="">Select a status</option>
    <option v-for="statusVal of options" :key="statusVal">
      {{ statusVal }}
    </option>
  </select>
</template>

<script>
import { computed, defineComponent, toRefs } from "vue";

export default defineComponent({
  name: "SelectOption",
  props: {
    selected: {
      type: String,
      required: true,
      default: "",
    },
    options: {
      type: Object,
      required: true,
      default: () => {},
    },
  },
  emits: ["update:selected"],
  setup(props, { emit }) {
    const { selected } = toRefs(props);
    const selectedComputed = computed({
      get: () => selected.value,
      set: (value) => emit("update:selected", value),
    });
    return {
      selectedComputed,
    };
  },
});
</script>
