<template>
  <div>
    <h1>Login</h1>
    <Loader v-show="loading" />
    <Error :error="error" />
    <form @submit.prevent="signin">
      <input type="email" name="email" v-model="email" />
      <input
        type="password"
        name="password"
        autocomplete="on"
        v-model="password"
      />
      <button type="submit">Signin</button>
    </form>
  </div>
</template>

<script>
import Loader from "../components/Loader";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";
import router from "../router";

export default {
  name: "Signin",
  components: {
    Loader,
    Error
  },
  data() {
    return {
      error: "",
      loading: false,
      email: "",
      password: ""
    };
  },
  methods: {
    async signin() {
      try {
        this.loading = true;
        await apiClient
          .post("/signin", {
            email: this.email,
            password: this.password
          })
          .then(res => {
            localStorage.setItem("access_token", res.data.access_token);
            localStorage.setItem("refresh_token", res.data.refresh_token);
            router.push("home");
          });
      } catch (e) {
        this.error = e.response.data.message;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
