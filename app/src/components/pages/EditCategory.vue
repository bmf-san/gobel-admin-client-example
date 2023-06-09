<template>
  <Main>
    <div class="container-readable">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Edit a category</h1>
          <Alert :error="error" />
          <FormCreateSingleInput
            v-model:text="textComputed"
            input-text-name="category"
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
import Loader from "@/components/atoms/Loader";
import Alert from "@/components/atoms/Alert";
import FormCreateSingleInput from "@/components/organisms/FormCreateSingleInput";
import apiCategories from "@/modules/api/categories";
import Main from "@/components/templates/Main";
import { useRoute, useRouter } from "vue-router";

export default defineComponent({
  name: "EditCategory",
  components: {
    Main,
    Loader,
    Alert,
    FormCreateSingleInput,
  },
  setup() {
    const route = useRoute();
    const router = useRouter();
    const error = ref("");
    const loading = ref("");
    const text = ref("");
    const textComputed = computed({
      get: () => text.value,
      set: (value) => (text.value = value),
    });
    const getCategoryById = async (id) => {
      loading.value = true;
      apiCategories
        .getCategoryById(id)
        .then((res) => {
          text.value = res.data.name;
          loading.value = false;
        })
        .catch((e) => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    const id = route.params.id;
    getCategoryById(id);
    const save = async () => {
      loading.value = true;
      apiCategories
        .patchCategory(route.params.id, text.value)
        .then(() => {
          loading.value = false;
          router.push({ name: "Categories" });
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
