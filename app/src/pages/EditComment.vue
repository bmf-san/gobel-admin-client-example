<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>EditComment</h1>
        <Error :error="error" />
        <form @submit.prevent="save">
          <label>Title</label>
          <p>{{ title }}</p>
          <label>Body</label>
          <p>{{ body }}</p>
          <select v-model="status">
            <option disabled value="">Select a status</option>
            <option v-for="status of statuses" :key="status">
              {{ status }}
            </option>
          </select>
          <input class="submit-button" type="submit" value="Save" />
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import Error from "../components/Error";
import consts from "../consts/comments";
import apiClient from "../modules/apiClient";
import storage from "../storage";

export default {
  name: "EditComment",
  components: {
    Loader,
    Error
  },
  data() {
    return {
      loading: false,
      error: "",
      postId: "",
      title: "",
      body: "",
      statuses: consts.STATUSES,
      status: ""
    };
  },
  created() {
    const id = this.$route.params.id;
    this.getComment(id).then(() => {
      this.getPost(this.postId);
    });
  },
  methods: {
    // TODO: DELETE 対応
    async getPost(id) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/posts/${id}`, {
            headers: {
              Authorization: "Bearer " + storage.getAccessToken()
            }
          })
          .then(res => {
            this.title = res.data.title;

            this.loading = false;
          });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    },
    async getComment(id) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/comments/${id}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.postId = res.data.id;
            this.body = res.data.body;
            this.status = res.data.status;
            this.loading = false;
          });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    },
    async save() {
      try {
        this.loading = true;
        await apiClient
          .patch(
            `/private/comments/${this.$route.params.id}/status`,
            {
              status: this.status
            },
            {
              headers: {
                Authorization: "Bearer " + localStorage.getItem("access_token")
              }
            }
          )
          .then(() => {
            this.loading = false;
          });
      } catch (e) {
        this.error = e.response.data.message;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
