<template>
  <Main>
    <div class="container">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Posts</h1>
          <router-link :to="{ name: 'CreatePost' }">
            <button>Create</button>
          </router-link>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <ListPosts
            button-text="Delete"
            :table-headers="tableHeaders"
            :posts="posts"
            @delete-items="deletePost"
          />
        </div>
      </div>
      <div class="row">
        <div class="col">
          <Pagination
            name="Posts"
            :page="page"
            :limit="limit"
            :pagecount="pageCount"
            @click="getPosts(page, limit)"
          />
        </div>
      </div>
    </div>
  </Main>
</template>

<script>
const defaultPage = 1;
const defaultLimit = 10;
const defaultPageCount = 10;

import { defineComponent, ref, onMounted } from "vue";
import Main from "@/components/templates/Main";
import Loader from "@/components/atoms/Loader";
import ListPosts from "@/components/organisms/ListPosts";
import Pagination from "@/components/organisms/Pagination";
import apiPosts from "@/modules/api/posts";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

export default defineComponent({
  name: "Posts",
  components: {
    Main,
    Loader,
    ListPosts,
    Pagination
  },
  setup() {
    const route = useRoute();
    const tableHeaders = [
      "ID",
      "Title",
      "Category",
      "Status",
      "Created at",
      "Updated at",
      "Edit",
      "Delete"
    ];
    const loading = ref("");
    const posts = ref({});
    const page = ref(defaultPage);
    const limit = ref(defaultLimit);
    const pageCount = ref(defaultPageCount);

    const getPosts = async (paramPage, paramLimit) => {
      loading.value = true;
      apiPosts
        .getPostsByQueryParams(paramPage, paramLimit)
        .then(res => {
          posts.value = res.data;
          page.value = Number(res.headers["pagination-page"]);
          limit.value = Number(res.headers["pagination-limit"]);
          pageCount.value = Number(res.headers["pagination-pagecount"]);
          loading.value = false;
        })
        .catch(e => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    const deletePost = async id => {
      if (!confirm("Is it really okay to delete it?")) {
        return;
      }
      try {
        loading.value = true;
        apiPosts.deletePost(id).then(() => {
          getPosts(page.value, limit.value);
          loading.value = false;
        });
      } catch (e) {
        console.log(e);
      } finally {
        loading.value = false;
      }
    };
    onBeforeRouteUpdate(async (to, _, next) => {
      getPosts(to.query.page, to.query.limit);
      next();
    });
    onMounted(() => {
      if (route.query.page == null || route.query.limit == null) {
        getPosts(defaultPage, defaultLimit);
      } else {
        getPosts(route.query.page, route.query.limit);
      }
    });
    return {
      tableHeaders,
      loading,
      posts,
      page,
      limit,
      pageCount,
      deletePost
    };
  }
});
</script>
