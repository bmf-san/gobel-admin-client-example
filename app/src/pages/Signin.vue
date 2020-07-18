<template>
  <div>
    <Loader v-show="loading" />
    <form @submit.prevent="signin">
      <input type="email" name="email" v-model="email" />
      <input type="password" name="password" v-model="password" />
      <button type="submit">Signin</button>
    </form>
  </div>
</template>

<script>
import Loader from "../components/Loader";
import apiClient from "../modules/apiClient";

export default {
  name: "Signin",
  components: {
    Loader,
  },
  data() {
    return {
      loading: false,
      email: "",
      password: "",
    };
  },
  methods: {
    async signin() {
      try {
        this.loading = true;
        await apiClient
          .post("/signin", {
            params: {
              email: this.email,
              password: this.password,
            },
          })
          .then((res) => {
            // TODO: set cookie
            console.log(res);
          });
      } catch (e) {
        console.log(e);
        // TODO: show error
      }
      this.loading = false;
    },
  },
};
</script>
