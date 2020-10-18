<template>
  <div class="categories">
    <h1>Categories</h1>
    <Loader v-show="loading" />
    <div>
      <article v-for="category in categories" :key="category.id">
        <router-link :to="{ name: 'EditCategory', params: { id: category.id } }"
          ><h1>{{ category.name }}</h1></router-link
        >
      </article>
      <Pagination
        name="Categories"
        :page="page"
        :limit="limit"
        :pagecount="pagecount"
        @click.native="getCategories(page, limit)"
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
  name: "Categories",
  components: {
    Loader,
    Pagination
  },
  data() {
    return {
      loading: false,
      categories: [],
      page: defaultPage,
      limit: defaultLimit,
      pagecount: defaultPageCount
    };
  },
  created() {
    this.getCategories(this.page, this.limit);
  },
  beforeRouteUpdate(to, from, next) {
    this.page = to.query.page;
    this.limit = to.query.limit;
    this.getCategories(this.page, this.limit);
    next();
  },
  methods: {
    async getCategories(page, limit) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/categories?page=${page}&limit=${limit}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.categories = res.data;
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
