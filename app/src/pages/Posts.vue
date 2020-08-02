<template>
  <div class="posts">
    <h1>Posts</h1>
    <Loader v-show="loading" />
    <div>
      <!-- TODO: 表示&ページャー -->
      posts list
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import apiClient from "../modules/apiClient";
export default {
  name: "Posts",
  components: {
    Loader
  },
  data() {
    return {
      loading: true,
      posts: []
    };
  },
  created() {
    this.getPosts();
  },
  methods: {
    async getPosts() {
      try {
        this.loading = true;
        await apiClient
          .get("/private/posts", {
            headers: {
              'Authorization': "Bearer " + localStorage.getItem("access_token"),
            }
          })
          .then(res => {
            console.log(res.headers);
            this.posts = res.data;
            this.loading = false;
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

<style scoped>
</style>
