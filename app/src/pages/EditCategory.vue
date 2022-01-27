<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>EditCategory</h1>
        <Error :error="error" />
        <form @submit.prevent="save">
          <input v-model="name" type="text" name="name" />
          <input class="submit-button" type="submit" value="Save" />
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Loader from "@/components/Loader.vue";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";

export default defineComponent({
  name: "EditCategory",
  components: {
    Loader,
    Error
  },
  data() {
    return {
      loading: false,
      error: "",
      name: ""
    };
  },
  created() {
    const id = this.$route.params.id;
    this.getCategory(id);
  },
  methods: {
    // TODO: DELETE 対応
    async getCategory(id) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/categories/${id}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.name = res.data.name;
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
            `/private/categories/${this.$route.params.id}`,
            {
              name: this.name
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
});
</script>
