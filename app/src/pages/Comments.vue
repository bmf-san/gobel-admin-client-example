<template>
  <div class="comments">
    <h1>Comments</h1>
    <Loader v-show="loading" />
    <div>
      <article v-for="comment in comments" :key="comment.id">
        <router-link :to="{ name: 'EditComment', params: { id: comment.id } }"
          ><h1>{{ comment.body }}</h1></router-link
        ><span>{{ comment.status }}</span>
      </article>
      <Pagination
        name="Comments"
        :page="page"
        :limit="limit"
        :pagecount="pagecount"
        @click.native="getComments(page, limit)"
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
  name: "Comments",
  components: {
    Loader,
    Pagination
  },
  data() {
    return {
      loading: false,
      comments: [],
      page: defaultPage,
      limit: defaultLimit,
      pagecount: defaultPageCount
    };
  },
  created() {
    this.getComments(this.page, this.limit);
  },
  beforeRouteUpdate(to, from, next) {
    this.page = to.query.page;
    this.limit = to.query.limit;
    this.getComments(this.page, this.limit);
    next();
  },
  methods: {
    async getComments(page, limit) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/comments?page=${page}&limit=${limit}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.comments = res.data;
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
