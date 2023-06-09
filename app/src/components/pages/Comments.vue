<template>
  <Main>
    <div class="container">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Comments</h1>
          <ListComments :table-headers="tableHeaders" :comments="comments" />
        </div>
      </div>
      <div class="row">
        <div class="col">
          <Pagination
            name="Comments"
            :page="page"
            :limit="limit"
            :pagecount="pageCount"
            @click="getComments(page, limit)"
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
import ListComments from "@/components/organisms/ListComments";
import Pagination from "@/components/organisms/Pagination";
import apiComments from "@/modules/api/comments";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

export default defineComponent({
  name: "Comments",
  components: {
    Main,
    Loader,
    ListComments,
    Pagination,
  },
  setup() {
    const route = useRoute();
    const tableHeaders = [
      "ID",
      "Body",
      "Status",
      "Created at",
      "Updated at",
      "Edit",
    ];
    const loading = ref("");
    const comments = ref({});
    const page = ref(defaultPage);
    const limit = ref(defaultLimit);
    const pageCount = ref(defaultPageCount);

    const getComments = async (paramPage, paramLimit) => {
      loading.value = true;
      apiComments
        .getCommentsByQueryParams(paramPage, paramLimit)
        .then((res) => {
          comments.value = res.data;
          page.value = Number(res.headers["pagination-page"]);
          limit.value = Number(res.headers["pagination-limit"]);
          pageCount.value = Number(res.headers["pagination-pagecount"]);
          loading.value = false;
        })
        .catch((e) => {
          console.log(e);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    onBeforeRouteUpdate(async (to, _, next) => {
      getComments(to.query.page, to.query.limit);
      next();
    });
    onMounted(() => {
      if (route.query.page == null || route.query.limit == null) {
        getComments(defaultPage, defaultLimit);
      } else {
        getComments(route.query.page, route.query.limit);
      }
    });
    return {
      tableHeaders,
      loading,
      comments,
      page,
      limit,
      pageCount,
    };
  },
});
</script>
