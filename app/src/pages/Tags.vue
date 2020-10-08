<template>
  <div class="tags">
    <h1>Tags</h1>
    <Loader v-show="loading" />
    <div>
      <article v-for="tag in tags" :key="tag.id">
        <router-link :to="{ name: 'EditTag', params: { id: tag.id } }"
          ><h1>{{ tag.name }}</h1></router-link
        >
      </article>
      <Pagination
        name="Tags"
        :page="page"
        :limit="limit"
        :pagecount="pagecount"
        @click.native="getTags(page, limit)"
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
  name: "Tags",
  components: {
    Loader,
    Pagination
  },
  data() {
    return {
      loading: false,
      tags: [],
      page: defaultPage,
      limit: defaultLimit,
      pagecount: defaultPageCount
    };
  },
  created() {
    this.getTags(this.page, this.limit);
  },
  beforeRouteUpdate(to, from, next) {
    this.page = to.query.page;
    this.limit = to.query.limit;
    this.getTags(this.page, this.limit);
    next();
  },
  methods: {
    async getTags(page, limit) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/tags?page=${page}&limit=${limit}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.tags = res.data;
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
