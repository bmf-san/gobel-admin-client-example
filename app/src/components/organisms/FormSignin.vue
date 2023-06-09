<template>
  <form @submit.prevent="$emit('submitForm')">
    <input-email v-model:email="emailComputed" :name="inputEmailName" />
    <input-password
      v-model:password="passwordComputed"
      :name="inputPasswordName"
      :autocomplete="inputPasswordIsAutocomplete"
    />
    <button-submit :value="buttonSubmitValue" />
  </form>
</template>

<script>
import { defineComponent, toRefs, computed } from "vue";

import ButtonSubmit from "@/components/atoms/ButtonSubmit";
import InputEmail from "@/components/atoms/InputEmail.vue";
import InputPassword from "@/components/atoms/InputPassword.vue";

export default defineComponent({
  name: "FormSignin",
  components: {
    ButtonSubmit,
    InputEmail,
    InputPassword,
  },
  props: {
    email: {
      type: String,
      required: true,
      default: "",
    },
    inputEmailName: {
      type: String,
      required: true,
      default: "",
    },
    password: {
      type: String,
      required: true,
      default: "",
    },
    inputPasswordName: {
      type: String,
      required: true,
      default: "",
    },
    inputPasswordIsAutocomplete: {
      type: Boolean,
      default: false,
    },
    buttonSubmitValue: {
      type: String,
      required: true,
      default: "",
    },
  },
  emits: ["update:email", "update:password", "submitForm"],
  setup(props, { emit }) {
    const { email } = toRefs(props);
    const emailComputed = computed({
      get: () => email.value,
      set: (value) => {
        emit("update:email", value);
      },
    });
    const { password } = toRefs(props);
    const passwordComputed = computed({
      get: () => password.value,
      set: (value) => {
        emit("update:password", value);
      },
    });
    return {
      emailComputed,
      passwordComputed,
    };
  },
});
</script>
