<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>Posts</h1>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>Category</th>
              <th>Status</th>
              <th>Created at</th>
              <th>Updated at</th>
              <th>Edit</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="post in posts" :key="post.id">
              <td>{{ post.id }}</td>
              <td>
                {{ post.title }}
              </td>
              <td>{{ post.category.name }}</td>
              <td>{{ post.status }}</td>
              <td>{{ post.created_at }}</td>
              <td>{{ post.updated_at }}</td>
              <td>
                <router-link :to="{ name: 'EditPost', params: { id: post.id } }"
                  >Edit</router-link
                >
              </td>
              <td>
                <a
                  @click.prevent.stop="deletePost(post.id)"
                  class="color-danger delete-link"
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
          name="Posts"
          :page="page"
          :limit="limit"
          :pagecount="pagecount"
          @click.native="getPosts(page, limit)"
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
  mounted() {
    const page = this.$route.query.page;
    const limit = this.$route.query.limit;
    if (page == null || limit == null) {
      this.getPosts(this.page, this.limit);
    } else {
      this.getPosts(page, limit);
    }
  },
  beforeRouteUpdate(to, from, next) {
    this.getPosts(to.query.page, to.query.limit);

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
    },
    async deletePost(id) {
      if (!confirm("Is it really okay to delete it?")) {
        return;
      }
      try {
        this.loading = true;
        await apiClient
          .delete(`/private/posts/${id}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            console.log(res);
            this.getPosts(this.page, this.limit);
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
