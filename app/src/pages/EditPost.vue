<template>
  <div class="container">
    <Loader v-show="loading" />
    <div class="row">
      <div class="col">
        <h1>EditPost</h1>
        <div>
          <Error :error="error" />
          <form @submit.prevent="save">
            <label for="title">Title</label>
            <input v-model="title" type="text" name="title" />
            <multiselect
              v-model="tags"
              placeholder=""
              label="name"
              track-by="id"
              :options="tagOptions"
              :multiple="true"
              :taggable="true"
              @tag="addTag"
            ></multiselect>
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
                <div class="preview" v-html="compileMarkdown"></div>
              </div>
            </div>
            <label for="status">Status</label>
            <select v-model="status">
              <option disabled></option>
              <option v-for="statusVal of statuses" :key="statusVal">
                {{ statusVal }}
              </option>
            </select>
            <input class="submit-button" type="submit" value="Save" />
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Loader from "@/components/Loader.vue";
import Error from "../components/Error";
import apiClient from "../modules/apiClient";
import consts from "../consts/posts";
import marked from "marked";
import Multiselect from "vue-multiselect";
import hljs from "highlight.js";
import "highlight.js/styles/github.css";
import storage from "../storage";

export default defineComponent({
  name: "EditPost",
  components: {
    Loader,
    Error,
    Multiselect
  },
  data() {
    return {
      loading: false,
      error: "",
      title: "",
      tags: [],
      categoryId: "",
      tagOptions: [],
      categories: [],
      markdown: "",
      statuses: consts.STATUSES,
      status: ""
    };
  },
  computed: {
    compileMarkdown: function() {
      return marked(this.markdown);
    }
  },
  created() {
    const id = this.$route.params.id;
    this.getPost(id);
    this.getTags();
    this.getCategories();
    marked.setOptions({
      langPrefix: "",
      highlight: function(code, lang) {
        return hljs.highlightAuto(code, [lang]).value;
      }
    });
  },
  methods: {
    addTag(newTag) {
      console.log(newTag);
      const tag = {
        name: newTag
      };
      this.tagOptions.push(tag);
      this.tags.push(tag);
    },
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
            this.categoryId = res.data.category.id;
            this.tags = res.data.tags;
            this.title = res.data.title;
            this.markdown = res.data.md_body;
            this.status = res.data.status;

            this.loading = false;
          });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    },
    customLabel(option) {
      return `${option.name}`;
    },
    async getTags() {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/tags?page=1&limit=9999`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.tagOptions = res.data;
            this.loading = false;
          });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    },
    async getCategories() {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/categories?page=1&limit=9999`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("access_token")
            }
          })
          .then(res => {
            this.categories = res.data;
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
        const tagIds = this.tags.map(obj => {
          return {
            id: obj.id
          };
        });
        await apiClient
          .patch(
            `/private/posts/${this.$route.params.id}`,
            {
              category_id: this.categoryId,
              tags: tagIds,
              title: this.title,
              md_body: this.markdown,
              html_body: this.compileMarkdown,
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
});
</script>

<style scoped>
textarea {
  resize: none;
  min-height: 600px;
}
</style>
