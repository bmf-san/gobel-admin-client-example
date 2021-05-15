<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>CreateTag</h1>
        <Error :error="error" />
        <form @submit.prevent="save">
          <input type="text" name="name" v-model="name" />
          <input class="submit-button" type="submit" value="Save" />
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";
import router from "../router";

export default {
  name: "CreateTag",
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
            `/private/tags`,
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
            router.push({ name: "EditTag", params: { id: res.data.id } });
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
