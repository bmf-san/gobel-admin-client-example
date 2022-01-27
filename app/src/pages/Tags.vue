<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>Tags</h1>
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
            <tr v-for="tag in tags" :key="tag.id">
              <td>{{ tag.id }}</td>
              <td>
                {{ tag.name }}
              </td>
              <td>{{ tag.created_at }}</td>
              <td>{{ tag.updated_at }}</td>
              <td>
                <router-link :to="{ name: 'EditTag', params: { id: tag.id } }"
                  >Edit</router-link
                >
              </td>
              <td>
                <a
                  class="color-danger delete-link"
                  @click.prevent.stop="deleteTag(tag.id)"
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
          name="Tags"
          :page="page"
          :limit="limit"
          :pagecount="pagecount"
          @click="getTags(page, limit)"
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
  name: "Tags",
  components: {
    Loader,
    Pagination
  },
  beforeRouteUpdate(to, from, next) {
    this.page = to.query.page;
    this.limit = to.query.limit;
    this.getTags(this.page, this.limit);
    next();
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
  mounted() {
    const page = this.$route.query.page;
    const limit = this.$route.query.limit;
    if (page == null || limit == null) {
      this.getTags(this.page, this.limit);
    } else {
      this.getTags(page, limit);
    }
  },
  methods: {
    async getTags(page, limit) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/tags?page=${page}&limit=${limit}`, {
            headers: {
              Authorization: "Bearer " + storage.getAccessToken()
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
    },
    async deleteTag(id) {
      if (!confirm("Is it really okay to delete it?")) {
        return;
      }
      try {
        this.loading = true;
        await apiClient
          .delete(`/private/tags/${id}`, {
            headers: {
              Authorization: "Bearer " + storage.getAccessToken()
            }
          })
          .then(res => {
            console.log(res);
            this.getTags(this.page, this.limit);
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
