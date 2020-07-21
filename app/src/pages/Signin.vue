<template>
  <div>
    <Loader v-show="loading" />
    <!-- TODO: レスポンスエラー表示 -->
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
      // TODO: レスポンスエラー表示
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
            console.log(res);
          });
      } catch (e) {
        console.log(e);
        // TODO: show error
        // TODO: サーバーサイドのレスポンスがjsonじゃないので直す
      }
      this.loading = false;
    },
  },
};
</script>
