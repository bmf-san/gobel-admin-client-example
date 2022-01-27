<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>Categories</h1>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Created at</th>
              <th>Updated at</th>
              <th>Edit</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="category in categories" :key="category.id">
              <td>{{ category.id }}</td>
              <td>
                {{ category.name }}
              </td>
              <td>{{ category.created_at }}</td>
              <td>{{ category.updated_at }}</td>
              <td>
                <router-link
                  :to="{ name: 'EditCategory', params: { id: category.id } }"
                  >Edit</router-link
                >
              </td>
              <td>
                <a
                  class="color-danger delete-link"
                  @click.prevent.stop="deleteCategories(category.id)"
                  >Delete</a
                >
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <Pagination
          name="Categories"
          :page="page"
          :limit="limit"
          :pagecount="pagecount"
          @click="getCategories(page, limit)"
        />
      </div>
    </div>
  </div>
</template>

<script>
const defaultPage = 1;
const defaultLimit = 10;
const defaultPageCount = 10;

import { defineComponent } from "vue";
import Loader from "@/components/Loader.vue";
import Pagination from "@/components/Pagination.vue";
import apiClient from "../modules/apiClient";
import storage from "../storage";

export default defineComponent({
  name: "Categories",
  components: {
    Loader,
    Pagination
  },
  beforeRouteUpdate(to, from, next) {
    this.page = to.query.page;
    this.limit = to.query.limit;
    this.getCategories(this.page, this.limit);
    next();
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
  mounted() {
    const page = this.$route.query.page;
    const limit = this.$route.query.limit;
    if (page == null || limit == null) {
      this.getCategories(this.page, this.limit);
    } else {
      this.getCategories(page, limit);
    }
  },
  methods: {
    async getCategories(page, limit) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/categories?page=${page}&limit=${limit}`, {
            headers: {
              Authorization: "Bearer " + storage.getAccessToken()
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
    },
    async deleteCategories(id) {
      if (!confirm("Is it really okay to delete it?")) {
        return;
      }
      try {
        this.loading = true;
        await apiClient
          .delete(`/private/categories/${id}`, {
            headers: {
              Authorization: "Bearer " + storage.getAccessToken()
            }
          })
          .then(res => {
            console.log(res);
            this.getCategories(this.page, this.limit);
            this.loading = false;
          });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    }
  }
});
</script>

<style scoped>
table {
  margin: 0 auto;
}
.delete-link:hover {
  color: var(--danger-color);
}
</style>
