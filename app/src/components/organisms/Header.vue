<template>
  <header>
    <div>
      <nav class="nav sp-hidden">
        <div class="col-nav">
          <div class="nav-left">
            <a class="nav-link-logo"><b>gobel-admin-client-example</b></a>
          </div>
          <div v-if="isActive" class="nav-right">
            <router-link class="nav-link" :to="{ name: 'Home' }"
              >Home</router-link
            >
            <router-link class="nav-link" :to="{ name: 'Posts' }"
              >Posts</router-link
            >
            <router-link class="nav-link" :to="{ name: 'Categories' }"
              >Categories</router-link
            >
            <router-link class="nav-link" :to="{ name: 'Tags' }"
              >Tags</router-link
            >
            <router-link class="nav-link" :to="{ name: 'Comments' }"
              >Comments</router-link
            >
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
          <div v-if="isActive" class="nav-right">
            <div id="nav-sp-drawer">
              <input id="nav-sp-input" type="checkbox" class="sp-hidden" />
              <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
              <label id="nav-sp-close" class="sp-hidden" for="nav-sp-input"
                ><span></span
              ></label>
              <div id="nav-sp-content">
                <h1>Navigation</h1>
                <router-link class="nav-link" :to="{ name: 'Home' }"
                  >Home</router-link
                >
                <router-link class="nav-link" :to="{ name: 'Posts' }"
                  >Posts</router-link
                >
                <router-link class="nav-link" :to="{ name: 'Categories' }"
                  >Categories</router-link
                >
                <router-link class="nav-link" :to="{ name: 'Tags' }"
                  >Tags</router-link
                >
                <router-link class="nav-link" :to="{ name: 'Comments' }"
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
import { computed, defineComponent, ref } from "vue";
import apiSignout from "@/modules/api/signout";
import router from "@/router";
import { useRoute } from "vue-router";

export default defineComponent({
  name: "Header",
  setup() {
    const route = useRoute();
    const isActive = computed(() => {
      return route.name !== "Signin";
    });
    const loading = ref("");
    const signout = async () => {
      loading.value = true;
      apiSignout
        .postSignout()
        .then(() => {
          loading.value = false;
          router.push({ name: "Signin" });
        })
        .catch(e => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    return {
      isActive,
      loading,
      signout
    };
  }
});
</script>
