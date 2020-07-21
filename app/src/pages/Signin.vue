<template>
  <div>
    <Loader v-show="loading" />
    <p v-if="errors.length">
    <b>Please correct the following error(s):</b>
    <ul>
      <li v-for="(error, key) in errors" :key="key">{{ error }}</li>
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
            // TODO: set error if has itt
            // TODO: set cookie
            console.log(res);
          });
      } catch (e) {
        console.log(e.message);
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>
