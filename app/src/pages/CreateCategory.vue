<template>
  <div class="createcategory">
    <h1>CreateCategory</h1>
    <Loader v-show="loading" />
    <div>
      <Error :error="error" />
      <form @submit.prevent="save">
        <input type="text" name="name" v-model="name" />
        <input type="submit" value="Save" />
      </form>
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";
import router from "../router";

export default {
  name: "CreateCategory",
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
  methods: {
    async save() {
      // TODO: want to support autosave using localstorage.
      try {
        this.loading = true;
        await apiClient
          .post(
            `/private/categories`,
            {
              name: this.name
            },
            {
              headers: {
                Authorization: "Bearer " + localStorage.getItem("access_token")
              }
            }
          )
          .then(res => {
            this.loading = false;
            router.push({ name: "EditCategory", params: { id: res.data.id } });
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
