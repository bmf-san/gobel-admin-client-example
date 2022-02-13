<template>
  <form @submit.prevent="$emit('submitForm')">
    <input-text v-model:text="textComputed" :name="inputTextName" />
    <button-submit :value="buttonSubmitValue" />
  </form>
</template>

<script>
import { defineComponent, toRefs, computed } from "vue";

import ButtonSubmit from "@/components/atoms/ButtonSubmit";
import InputText from "@/components/atoms/InputText";

export default defineComponent({
  name: "FormCreateSingleInput",
  components: {
    ButtonSubmit,
    InputText
  },
  props: {
    text: {
      type: String,
      required: true,
      default: ""
    },
    inputTextName: {
      type: String,
      required: true,
      default: ""
    },
    buttonSubmitValue: {
      type: String,
      required: true,
      default: ""
    }
  },
  emits: ["update:text", "submitForm"],
  setup(props, { emit }) {
    const { text } = toRefs(props);
    const textComputed = computed({
      get: () => text.value,
      set: value => {
        emit("update:text", value);
      }
    });
    return {
      textComputed
    };
  }
});
</script>
