<template>
  <div class="post">
    <h1>Posts</h1>
    <input type="text" placeholder="name" :value="name" @input="updateName" />
    <Loader v-show="loading" />
    <ul v-show="isActive">
      <li>login:{{ user.login }}</li>
      <li>name:{{ user.name }}</li>
      <li>location:{{ user.location }}</li>
      <li>blog:{{ user.blog }}</li>
    </ul>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import apiClient from "../modules/apiClient";
export default {
  name: "Post",
  components: {
    Loader
  },
  data() {
    return {
      user: "",
      name: "",
      loading: true,
      isActive: false
    };
  },
  created() {
    this.getUser();
  },
  methods: {
    updateName(e) {
      this.name = e.target.value;
      this.getUser();
    },
    async getUser() {
      try {
        await apiClient.get("/users/" + this.name).then(response => {
          this.user = response.data;
          this.loading = false;
          this.isActive = true;
        });
      } catch (e) {
        this.loading = false;
        this.isActive = false;
        return e;
      }
    }
  }
};
</script>

<style scoped>
ul {
  text-align: left;
  width: fit-content;
  margin: 0 auto;
}
</style>