<template>
  <div class="editpost">
    <h1>EditPost</h1>
    <Loader v-show="loading" />
    <div>
      <form>
        <textarea v-model="markdown"></textarea>
      </form>
      <div class="preview" v-html="compileMarkdown"></div>
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";
import apiClient from "../modules/apiClient";
import marked from 'marked';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';

export default {
  name: "EditPost",
  components: {
    Loader
  },
  data() {
    return {
      loading: true,
      post: "",
      markdown: ""
    };
  },
  created() {
    const id = this.$route.params.id;
    this.getPost(id);
    marked.setOptions({
      langPrefix: '',
      highlight: function (code, lang) {
        return hljs.highlightAuto(code, [lang]).value
      }
    });
  },
  computed: {
    compileMarkdown: function () {
      return marked(this.markdown)
    }
  },
  methods: {
    async getPost(id) {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/posts/${id}`, {
            headers: {
              'Authorization': "Bearer " + localStorage.getItem("access_token"),
            }
          })
          .then(res => {
            this.posts = res.data;
            this.loading = false;
        });
      } catch (e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    },
    // TODO: add a method for edit.
  }
};
</script>

<style scoped>
</style>
