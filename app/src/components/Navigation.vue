<template>
  <header>
    <div>
      <nav class="nav sp-hidden">
        <div class="col-nav">
          <div class="nav-left">
            <a class="nav-link-logo"><b>gobel-admin-client-example</b></a>
          </div>
          <div class="nav-right" v-if="isActive">
            <router-link class="nav-link" to="/">Home</router-link>
            <router-link class="nav-link" to="/posts">Posts</router-link>
            <router-link class="nav-link" to="/categories"
              >Categories</router-link
            >
            <router-link class="nav-link" to="/tags">Tags</router-link>
            <router-link class="nav-link" to="/comments">Comments</router-link>
            <button class="nav-link color-text-reverse" @click="signout()">
              Sign out
            </button>
          </div>
        </div>
      </nav>
      <nav class="nav pc-hidden">
        <div class="col-nav">
          <div class="nav-left">
            <a class="nav-link text-decoration-none"
              >gobel-admin-client-example</a
            >
          </div>
          <div class="nav-right" v-if="isActive">
            <div id="nav-sp-drawer">
              <input id="nav-sp-input" type="checkbox" class="sp-hidden" />
              <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
              <label class="sp-hidden" id="nav-sp-close" for="nav-sp-input"
                ><span></span
              ></label>
              <div id="nav-sp-content">
                <h1>Navigation</h1>
                <router-link class="nav-link" to="/">Home</router-link>
                <router-link class="nav-link" to="/posts">Posts</router-link>
                <router-link class="nav-link" to="/categories"
                  >Categories</router-link
                >
                <router-link class="nav-link" to="/tags">Tags</router-link>
                <router-link class="nav-link" to="/comments"
                  >Comments</router-link
                >
                <button class="nav-link color-text-reverse" @click="signout()">
                  Sign out
                </button>
              </div>
            </div>
          </div>
        </div>
      </nav>
    </div>
  </header>
</template>

<script>
import apiClient from "../modules/apiClient";
import router from "../router";
import storage from "../storage";
export default {
  name: "Navigation",
  computed: {
    isActive() {
      return this.$route.name !== "Signin";
    }
  },
  methods: {
    async signout() {
      try {
        await apiClient
          .post(
            "/private/signout",
            {},
            {
              headers: {
                Authorization: "Bearer " + storage.getAccessToken()
              }
            }
          )
          .then(() => {
            storage.removeAccessToken();
            storage.removeRefreshToken();
            storage.setIsSignin(false);
            router.push({ name: "Signin" });
          });
      } catch (e) {
        console.log(e);
      }
    }
  }
};
</script>
