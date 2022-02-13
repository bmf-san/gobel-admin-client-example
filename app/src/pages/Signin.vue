<template>
  <div class="container-readable">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>Sign in</h1>
        <Error :error="error" />
        <form @submit.prevent="signin">
          <input v-model="email" type="email" name="email" />
          <input
            v-model="password"
            type="password"
            name="password"
            autocomplete="on"
          />
          <input class="submit-button" type="submit" value="Sign in" />
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Loader from "../components/Loader";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";
import router from "../router";
import storage from "../storage";

export default defineComponent({
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
            storage.setAccessToken(res.data.access_token);
            storage.setRefreshToken(res.data.refresh_token);
            storage.setIsSignin(true);
            router.push({ name: "Home" });
          });
      } catch (e) {
        console.log(e.message);
        this.error = e.response.data.message;
      } finally {
        this.loading = false;
      }
    }
  }
});
</script>
