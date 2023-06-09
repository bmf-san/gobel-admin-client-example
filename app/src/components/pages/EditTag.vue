<template>
  <Main>
    <div class="container-readable">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Edit a tag</h1>
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
import Loader from "@/components/atoms/Loader";
import Alert from "@/components/atoms/Alert";
import FormCreateSingleInput from "@/components/organisms/FormCreateSingleInput";
import apiTags from "@/modules/api/tags";
import Main from "@/components/templates/Main";
import { useRoute, useRouter } from "vue-router";

export default defineComponent({
  name: "EditTag",
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
    const getTagById = async (id) => {
      loading.value = true;
      apiTags
        .getTagById(id)
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
    getTagById(id);
    const save = async () => {
      loading.value = true;
      apiTags
        .patchTag(route.params.id, text.value)
        .then(() => {
          loading.value = false;
          router.push({ name: "Tags" });
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
