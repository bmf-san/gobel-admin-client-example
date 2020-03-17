<template>
  <div class="editpost">
    <h1>EditPost</h1>
    <Loader v-show="loading" />
    <div>
      <form>
        <multiselect 
          v-model="tag" 
          :options="tagOptions"
          :multiple="true"
          :custom-label="customLabel"
          :options-limit="10"
        >
        </multiselect>
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
import Multiselect from 'vue-multiselect'
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';

export default {
  name: "EditPost",
  components: {
    Loader,
    Multiselect 
  },
  data() {
    return {
      loading: true,
      post: "",
      tag: "",
      tagOptions: [],
      markdown: ""
    };
  },
  created() {
    const id = this.$route.params.id;
    this.getPost(id);
    this.getTags();
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
    customLabel (option) {
      return `${option.name}`
    },
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
    async getTags() {
      try {
        this.loading = true;
        await apiClient
          .get(`/private/tags?page=1&limit=9999`, {
            headers: {
              'Authorization': "Bearer " + localStorage.getItem("access_token"),
            }
          })
          .then(res => {
            this.tagOptions = res.data;
            this.loading = false;
          })
      } catch(e) {
        console.log(e);
      } finally {
        this.loading = false;
      }
    }
    // TODO: add a method for edit.
    // TODO: objectから任意のkeyのvalueだけのセットを生成 多分あとで使いそうなのでメモ
    // const tagIds = res.data.map((obj) => obj.id);
    // console.log(tagIds);
  }
};
</script>

<style scoped>
</style>
