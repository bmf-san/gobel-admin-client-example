<template>
  <div class="editcomment">
    <h1>EditComment</h1>
    <Loader v-show="loading" />
    <div>
      <Error :error="error" />
      <form @submit.prevent="save">
        <p>{{ title }}</p>
        <p>{{ body }}</p>
        <select v-model="status">
          <option disabled value="">Select a status</option>
          <option v-for="status of statuses" :key="status">
            {{ status }}
          </option>
        </select>
        <input type="submit" value="Save" />
      </form>
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import Error from "../components/Error";
import consts from "../consts/comments";
import apiClient from "../modules/apiClient";

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
      status: "",
    };
  },
  created() {
    const id = this.$route.params.id;
    this.getComment(id).then(() => {
        this.getPost(this.postId);
    });
  },
  methods: {
    async getPost(id) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/posts/${id}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
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
            `/private/comments/${this.$route.params.id}`,
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

<style scoped></style>
