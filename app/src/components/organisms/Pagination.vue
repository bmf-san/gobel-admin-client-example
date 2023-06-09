<template>
  <div class="pagination">
    <ul class="pagination-list">
      <li v-for="i of pagecount" :key="i">
        <router-link
          :to="{ name: name, query: { page: i, limit: limit } }"
          :class="{ current: isCurrent(i) }"
          ><span>{{ i }}</span></router-link
        >
      </li>
    </ul>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "Pagination",
  props: {
    name: {
      type: String,
      required: true,
    },
    page: {
      type: Number,
      required: true,
      default: 1,
    },
    limit: {
      type: Number,
      required: true,
      default: 10,
    },
    pagecount: {
      type: Number,
      required: true,
      default: 10,
    },
  },
  setup(props) {
    const isCurrent = (i) => {
      return i == props.page;
    };
    return {
      isCurrent,
    };
  },
});
</script>

<style scocped>
.current {
  background: var(--main-color);
  color: #fff !important;
}
</style>
