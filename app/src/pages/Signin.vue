<template>
  <div>
    <Loader v-show="loading" />
    <!-- TODO: 1.0.1 will support validation errors -->
    <Error :error="error" />
    <form @submit.prevent="signin">
      <input type="email" name="email" v-model="email" />
      <input type="password" name="password" autocomplete="on" v-model="password" />
      <button type="submit">Signin</button>
    </form>
  </div>
</template>

<script>
import Loader from "../components/Loader";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";
import router from '../router'

export default {
  name: "Signin",
  components: {
    Loader,
    Error,
  },
  data() {
    return {
      error: "",
      loading: false,
      email: "",
      password: "",
    };
  },
  methods: {
    // TODO: アクセストークンのリフレッシュ対応が必要
    // cf, https://qiita.com/nqyutq/items/c025ad9dbe732c85e124
    async signin() {
      try {
        this.loading = true;
        await apiClient
          .post("/signin", {
            email: this.email,
            password: this.password,
          })
          .then(() => {
            router.push('home');
          });
      } catch (e) {
        this.error = e.response.data.message
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>
