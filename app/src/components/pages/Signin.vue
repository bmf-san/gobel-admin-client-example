<template>
  <Main>
    <div class="container-readable">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Sign in</h1>
          <Alert :error="error" />
          <FormSignin
            v-model:email="emailComputed"
            v-model:password="passwordComputed"
            input-email-name="email"
            input-password-name="password"
            input-password-is-autocomplete="true"
            button-submit-value="Sign in"
            @submitForm="signin"
          />
        </div>
      </div>
    </div>
  </Main>
</template>

<script>
import { computed, defineComponent, ref } from "vue";
import Main from "@/components/templates/Main";
import Loader from "@/components/atoms/Loader";
import Alert from "@/components/atoms/Alert";
import FormSignin from "@/components/organisms/FormSignin";
import apiSignin from "@/modules/api/signin";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "Signin",
  components: {
    Main,
    Loader,
    Alert,
    FormSignin,
  },
  setup() {
    const router = useRouter();
    const error = ref("");
    const loading = ref("");
    const email = ref("");
    const emailComputed = computed({
      get: () => email.value,
      set: (value) => (email.value = value),
    });
    const password = ref("");
    const passwordComputed = computed({
      get: () => password.value,
      set: (value) => (password.value = value),
    });
    const signin = async () => {
      loading.value = true;
      apiSignin
        .postSignin(email.value, password.value)
        .then(() => {
          loading.value = false;
          router.push({ name: "Home" });
        })
        .catch((e) => {
          console.log(e);
          error.value = e.response.data.message;
        })
        .finally(() => {
          loading.value = false;
        });
    };
    return {
      error,
      loading,
      email,
      emailComputed,
      password,
      passwordComputed,
      signin,
    };
  },
});
</script>
