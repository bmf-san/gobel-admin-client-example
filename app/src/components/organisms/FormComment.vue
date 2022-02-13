<template>
  <form @submit.prevent="$emit('submitForm')">
    <select-option
      v-model:selected="selectedComputed"
      :options="selectOptions"
      :name="selectName"
    />
    <button-submit :value="buttonSubmitValue" />
  </form>
</template>

<script>
import { defineComponent, toRefs, computed } from "vue";

import ButtonSubmit from "@/components/atoms/ButtonSubmit";
import SelectOption from "@/components/atoms/SelectOption";

export default defineComponent({
  name: "FormComment",
  components: {
    ButtonSubmit,
    SelectOption
  },
  props: {
    selected: {
      type: String,
      required: true,
      default: ""
    },
    selectOptions: {
      type: Object,
      required: true,
      default: () => {}
    },
    buttonSubmitValue: {
      type: String,
      required: true,
      default: ""
    }
  },
  emits: ["update:selected", "submitForm"],
  setup(props, { emit }) {
    const { selected } = toRefs(props);
    const selectedComputed = computed({
      get: () => selected.value,
      set: value => {
        emit("update:selected", value);
      }
    });
    return {
      selectedComputed
    };
  }
});
</script>
