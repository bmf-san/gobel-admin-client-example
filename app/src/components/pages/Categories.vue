<template>
  <Main>
    <div class="container">
      <Loader v-show="loading" />
      <div class="row">
        <div class="col">
          <h1>Categories</h1>
          <router-link :to="{ name: 'CreateCategory' }">
            <button>Create</button>
          </router-link>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <ListCategories
            button-text="Delete"
            :table-headers="tableHeaders"
            :categories="categories"
            @delete-items="deleteCategory"
          />
        </div>
      </div>
      <div class="row">
        <div class="col">
          <Pagination
            name="Categories"
            :page="page"
            :limit="limit"
            :pagecount="pageCount"
            @click="getCategories(page, limit)"
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
import ListCategories from "@/components/organisms/ListCategories";
import Pagination from "@/components/organisms/Pagination";
import apiCategories from "@/modules/api/categories";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

export default defineComponent({
  name: "Categories",
  components: {
    Main,
    Loader,
    ListCategories,
    Pagination,
  },
  setup() {
    const route = useRoute();
    const tableHeaders = [
      "ID",
      "Name",
      "Created at",
      "Updated at",
      "Edit",
      "Delete",
    ];
    const loading = ref("");
    const categories = ref({});
    const page = ref(defaultPage);
    const limit = ref(defaultLimit);
    const pageCount = ref(defaultPageCount);

    const getCategories = async (paramPage, paramLimit) => {
      loading.value = true;
      apiCategories
        .getCategoriesByQueryParams(paramPage, paramLimit)
        .then((res) => {
          categories.value = res.data;
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
    const deleteCategory = async (id) => {
      if (!confirm("Is it really okay to delete it?")) {
        return;
      }
      try {
        loading.value = true;
        apiCategories.deleteCategory(id).then(() => {
          getCategories(page.value, limit.value);
          loading.value = false;
        });
      } catch (e) {
        console.log(e);
      } finally {
        loading.value = false;
      }
    };
    onBeforeRouteUpdate(async (to, _, next) => {
      getCategories(to.query.page, to.query.limit);
      next();
    });
    onMounted(() => {
      if (route.query.page == null || route.query.limit == null) {
        getCategories(defaultPage, defaultLimit);
      } else {
        getCategories(route.query.page, route.query.limit);
      }
    });
    return {
      tableHeaders,
      loading,
      categories,
      page,
      limit,
      pageCount,
      deleteCategory,
    };
  },
});
</script>
