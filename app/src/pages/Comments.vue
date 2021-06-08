<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>Comments</h1>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Body</th>
              <th>Status</th>
              <th>Created at</th>
              <th>Updated at</th>
              <th>Edit</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="comment in comments" :key="comment.id">
              <td>{{ comment.id }}</td>
              <td>
                {{ comment.body }}
              </td>
              <td>
                {{ comment.status }}
              </td>
              <td>{{ comment.created_at }}</td>
              <td>{{ comment.updated_at }}</td>
              <td>
                <router-link
                  :to="{ name: 'EditComment', params: { id: comment.id } }"
                  >Edit</router-link
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
          name="Comments"
          :page="page"
          :limit="limit"
          :pagecount="pagecount"
          @click.native="getComments(page, limit)"
        />
      </div>
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
  mounted() {
    const page = this.$route.query.page;
    const limit = this.$route.query.limit;
    if (page == null || limit == null) {
      this.getComments(this.page, this.limit);
    } else {
      this.getComments(page, limit);
    }
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

<style scoped>
table {
  margin: 0 auto;
}
.delete-link:hover {
  color: var(--danger-color);
}
</style>
