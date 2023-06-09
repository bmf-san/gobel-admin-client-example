<template>
  <Main>
    <div class="container-readable">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>EditComment</h1>
          <Alert :error="error" />
          <label>Title</label>
          <p>{{ title }}</p>
          <label>Body</label>
          <p>{{ body }}</p>
          <FormComment
            v-model:selected="selectedComputed"
            :select-options="selectOptions"
            select-name="comment"
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
import FormComment from "@/components/organisms/FormComment";
import consts from "@/consts/comments";
import apiComments from "@/modules/api/comments";
import apiPosts from "@/modules/api/posts";
import Main from "@/components/templates/Main";
import { useRoute, useRouter } from "vue-router";

export default defineComponent({
  name: "EditComment",
  components: {
    Main,
    Loader,
    Alert,
    FormComment,
  },
  setup() {
    const route = useRoute();
    const router = useRouter();
    const loading = ref("");
    const error = ref("");
    const title = ref("");
    const body = ref("");
    const selected = ref("");
    const selectedComputed = computed({
      get: () => selected.value,
      set: (value) => (selected.value = value),
    });
    const selectOptions = consts.STATUSES;
    const postId = ref("");
    const getPostById = async (id) => {
      loading.value = true;
      apiPosts
        .getPostById(id)
        .then((res) => {
          title.value = res.data.title;
          loading.value = false;
        })
        .catch((e) => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    const getCommentById = async (id) => {
      loading.value = true;
      apiComments
        .getCommentById(id)
        .then((res) => {
          postId.value = res.data.id;
          body.value = res.data.body;
          selected.value = res.data.status;
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
    getCommentById(id).then(() => {
      getPostById(postId.value);
    });
    const save = async () => {
      loading.value = true;
      apiComments
        .patchComment(route.params.id, selected.value)
        .then(() => {
          loading.value = false;
          router.push({ name: "Comments" });
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
      title,
      body,
      selected,
      selectedComputed,
      selectOptions,
      getPostById,
      getCommentById,
      save,
    };
  },
});
</script>
