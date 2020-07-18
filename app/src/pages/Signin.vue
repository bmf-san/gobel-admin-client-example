<template>
  <div>
    <Loader v-show="loading" />
    <form @submit.prevent="signin" v-show="isActive">
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
      loading: true,
      email: "",
      password: "",
    };
  },
  methods: {
    async signin() {
      try {
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
            this.loading = false;
          });
      } catch (e) {
        this.loading = false;
        console.log(e);
        // TODO: show error
        return e;
      }
    },
  },
};
</script>
