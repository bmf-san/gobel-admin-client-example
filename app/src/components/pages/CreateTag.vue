<template>
  <Main>
    <div class="container-readable">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Create a tag</h1>
          <Alert :error="error" />
          <FormCreateSingleInput
            v-model:text="textComputed"
            input-text-name="tag"
            button-submit-value="Save"
            @submitForm="save"
          />
        </div>
      </div>
    </div>
  </Main>
</template>

<script>
import { computed, defineComponent, ref } from "vue";
import Main from "@/components/templates/Main.vue";
import Loader from "@/components/atoms/Loader.vue";
import Alert from "@/components/atoms/Alert";
import FormCreateSingleInput from "@/components/organisms/FormCreateSingleInput";
import apiTags from "@/modules/api/tags";
import router from "@/router";

export default defineComponent({
  name: "CreateTag",
  components: {
    Main,
    Loader,
    Alert,
    FormCreateSingleInput,
  },
  setup() {
    const error = ref("");
    const loading = ref("");
    const text = ref("");
    const textComputed = computed({
      get: () => text.value,
      set: (value) => (text.value = value),
    });
    const save = async () => {
      loading.value = true;
      apiTags
        .postTag(text.value)
        .then((res) => {
          loading.value = false;
          router.push({ name: "EditTag", params: { id: res.data.id } });
        })
        .catch((e) => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    return {
      error,
      loading,
      text,
      textComputed,
      save,
    };
  },
});
</script>

<style scoped></style>
