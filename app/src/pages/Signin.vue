<template>
  <div>
    <Loader v-show="loading" />
    <p v-if="errors.length">
    <b>Please correct the following error(s):</b>
    <ul>
      <!-- TODO: error ouputs -->
      <li v-for="error in errors" :key="error.message">{{ error.message }}</li>
    </ul>
    </p>
    <form @submit.prevent="signin">
      <input type="email" name="email" v-model="email" />
      <input type="password" name="password" autocomplete="on" v-model="password" />
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
      errors: [],
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
            email: this.email,
            password: this.password,
          })
          .then((res) => {
            // TODO: set cookie
            // TODO: move to home by router.push
            console.log(res);
          });
      } catch (e) {
        // TODO: error mmessageを出力
        this.errors = e.response.data;
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>
