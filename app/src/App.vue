<template>
  <div id="app">
    <Loader v-show="loading" />
    <div id="nav">
      <router-link to="/">Home</router-link>
      <router-link to="/posts">Posts</router-link>
      <router-link to="/categories">Categories</router-link>
      <router-link to="/tags">Tags</router-link>
      <router-link to="/comments">Comments</router-link>
      <!-- TODO: It may be necessary to introduce a store pattern to adjust the display conditions -->
      <button @click="signout()">Signout</button>
    </div>
    <router-view />
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import apiClient from "./modules/apiClient";
import router from "./router";

export default {
  name: "App",
  components: {
    Loader
  },
  data() {
    return {
      loading: false
    };
  },
  methods: {
    async signout() {
      try {
        this.loading = true;
        await apiClient
          .post(
            "/private/signout",
            {},
            {
              headers: {
                Authorization: "Bearer " + localStorage.getItem("access_token")
              }
            }
          )
          .then(() => {
            this.loading = false;
            this.isActive = false;
            localStorage.removeItem("access_token");
            localStorage.removeItem("refresh_token");
            router.push({ name: "Signin" });
          });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
