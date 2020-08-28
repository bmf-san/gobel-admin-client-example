<template>
  <div class="posts">
    <h1>Posts</h1>
    <Loader v-show="loading" />
    <div>
      <article v-for="post in posts" :key="post.id">
        <router-link :to="{ name: 'EditPost', params: { id: post.id } }"
          ><h1>{{ post.title }}</h1></router-link
        >
      </article>
      <Pagination
        name="Posts"
        :page="page"
        :limit="limit"
        :pagecount="pagecount"
        @click.native="getPosts(page, limit)"
      />
    </div>
  </div>
</template>

<script>
const defaultPage = 1;
const defaultLimit = 10;
const defaultPageCount = 10;

import Loader from "@/components/Loader.vue";
import Pagination from "@/components/Pagination.vue";
import apiClient from "../modules/apiClient";
export default {
  name: "Posts",
  components: {
    Loader,
    Pagination
  },
  data() {
    return {
      loading: false,
      posts: [],
      page: defaultPage,
      limit: defaultLimit,
      pagecount: defaultPageCount
    };
  },
  created() {
    this.getPosts(this.page, this.limit);
  },
  beforeRouteUpdate(to, from, next) {
    this.page = to.query.page;
    this.limit = to.query.limit;
    this.getPosts(this.page, this.limit);
    next();
  },
  methods: {
    async getPosts(page, limit) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/posts?page=${page}&limit=${limit}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.posts = res.data;
            this.page = Number(res.headers["pagination-page"]);
            this.limit = Number(res.headers["pagination-limit"]);
            this.pagecount = Number(res.headers["pagination-pagecount"]);
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

<style scoped></style>
