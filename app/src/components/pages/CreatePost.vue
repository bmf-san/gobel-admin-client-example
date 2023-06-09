<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>CreatePost</h1>
        <div>
          <Alert :error="error" />
          <form @submit.prevent="save">
            <label for="title">Title</label>
            <input v-model="title" type="text" name="title" />
            <label for="tags">Tags</label>
            <Multiselect
              v-model="tags"
              mode="tags"
              :close-on-select="false"
              :searchable="true"
              :create-option="false"
              :options="tagOptions"
              :limit="10"
            >
              <template #tag="{ option }">
                <span class="tag margin-right-1rem">
                  {{ option.label }}
                </span>
              </template>
            </Multiselect>
            <label for="category">Category</label>
            <select v-model="categoryId">
              <option disabled></option>
              <option
                v-for="category in categories"
                :key="category.id"
                :value="category.id"
              >
                {{ category.name }}
              </option>
            </select>
            <label for="tab">Body</label>
            <div class="tab">
              <input
                id="write-tab"
                type="radio"
                name="tab"
                class="tab-toggle"
                checked="checked"
              /><label class="tab-label" for="write-tab">Write</label>
              <div class="col-tab padding-left-0rem padding-right-0rem">
                <textarea v-model="markdown"></textarea>
              </div>
              <input
                id="preview-tab"
                type="radio"
                name="tab"
                class="tab-toggle"
              />
              <label class="tab-label" for="preview-tab">Preview</label>
              <div class="col-tab padding-left-0rem padding-right-0rem">
                <!-- eslint-disable-next-line vue/no-v-html -->
                <div class="preview" v-html="compiledMarkdown"></div>
              </div>
            </div>
            <label for="status">Status</label>
            <select v-model="status">
              <option disabled></option>
              <option v-for="statusVal in statuses" :key="statusVal">
                {{ statusVal }}
              </option>
            </select>
            <input class="button-submit" type="submit" value="Save" />
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, defineComponent, ref } from "vue";
import Loader from "@/components/atoms/Loader";
import Alert from "@/components/atoms/Alert";
import apiPosts from "@/modules/api/posts";
import apiTags from "@/modules/api/tags";
import apiCategories from "@/modules/api/categories";
import consts from "@/consts/posts";
import {marked} from "marked";
import Multiselect from "@vueform/multiselect";
import hljs from "highlight.js";
import "highlight.js/styles/github.css";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "CreatePost",
  components: {
    Loader,
    Alert,
    Multiselect
  },
  setup() {
    const router = useRouter();
    const loading = ref("");
    const error = ref("");
    const title = ref("");
    const tags = ref([]);
    const categoryId = ref("");
    const tagOptions = ref([]);
    const categories = ref([]);
    const markdown = ref("");
    const compiledMarkdown = computed(() => {
      return marked(markdown.value);
    });
    const status = ref("");
    const statuses = consts.STATUSES;
    const getTags = async () => {
      loading.value = true;
      apiTags
        .getTags()
        .then(res => {
          tagOptions.value = res.data.map(tag => {
            return {
              value: tag.id,
              label: tag.name
            };
          });
          loading.value = false;
        })
        .catch(e => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    const getCategories = async () => {
      loading.value = true;
      apiCategories
        .getCategories()
        .then(res => {
          categories.value = res.data;
          loading.value = false;
        })
        .catch(e => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    const save = async () => {
      loading.value = true;
      const tagIds = tags.value.map(id => {
        return {
          id: id
        };
      });
      apiPosts
        .postPost(
          categoryId.value,
          tagIds,
          title.value,
          markdown.value,
          compiledMarkdown.value,
          status.value
        )
        .then(() => {
          loading.value = false;
          router.push({ name: "Posts" });
        })
        .catch(e => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    getTags();
    getCategories();
    marked.use({
      langPrefix: "",
      highlight: function(code, lang) {
        return hljs.highlightAuto(code, [lang]).value;
      }
    });
    return {
      loading,
      error,
      title,
      tags,
      categoryId,
      tagOptions,
      categories,
      markdown,
      statuses,
      status,
      compiledMarkdown,
      getTags,
      getCategories,
      save
    };
  }
});
</script>

<style src="@vueform/multiselect/themes/default.css"></style>
<style>
.multiselect {
  height: 5rem;
  border-radius: 0.3rem;
  background-color: #fff;
  border: 0.1rem solid var(--border-color);
  width: 100%;
  font-size: 1.6rem;
  padding: 1rem;
  margin-bottom: 2.2rem;
}
.multiselect:focus {
  background: black;
}
.multiselect-tags {
  padding: 0;
}
.multiselect-tags-search {
  position: relative;
  padding: 1rem;
  height: 0.5rem;
}
</style>
<style scoped>
textarea {
  resize: none;
  min-height: 600px;
}
</style>
