<template>
  <Main>
    <div class="container">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Tags</h1>
          <router-link :to="{ name: 'CreateTag' }">
            <button>Create</button>
          </router-link>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <ListTags
            button-text="Delete"
            :table-headers="tableHeaders"
            :tags="tags"
            @delete-items="deleteTag"
          />
        </div>
      </div>
      <div class="row">
        <div class="col">
          <Pagination
            name="Tags"
            :page="page"
            :limit="limit"
            :pagecount="pageCount"
            @click="getTags(page, limit)"
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
import ListTags from "@/components/organisms/ListTags";
import Pagination from "@/components/organisms/Pagination";
import apiTags from "@/modules/api/tags";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

export default defineComponent({
  name: "Tags",
  components: {
    Main,
    Loader,
    ListTags,
    Pagination
  },
  setup() {
    const route = useRoute();
    const tableHeaders = [
      "ID",
      "Name",
      "Created at",
      "Updated at",
      "Edit",
      "Delete"
    ];
    const loading = ref("");
    const tags = ref({});
    const page = ref(defaultPage);
    const limit = ref(defaultLimit);
    const pageCount = ref(defaultPageCount);

    const getTags = async (paramPage, paramLimit) => {
      loading.value = true;
      apiTags
        .getTagsByQueryParams(paramPage, paramLimit)
        .then(res => {
          tags.value = res.data;
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
    const deleteTag = async id => {
      if (!confirm("Is it really okay to delete it?")) {
        return;
      }
      try {
        loading.value = true;
        apiTags.deleteTag(id).then(() => {
          getTags(page.value, limit.value);
          loading.value = false;
        });
      } catch (e) {
        console.log(e);
      } finally {
        loading.value = false;
      }
    };
    onBeforeRouteUpdate(async (to, _, next) => {
      getTags(to.query.page, to.query.limit);
      next();
    });
    onMounted(() => {
      if (route.query.page == null || route.query.limit == null) {
        getTags(defaultPage, defaultLimit);
      } else {
        getTags(route.query.page, route.query.limit);
      }
    });
    return {
      tableHeaders,
      loading,
      tags,
      page,
      limit,
      pageCount,
      deleteTag
    };
  }
});
</script>
