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
export default {
  name: "Pagination",
  props: {
    name: String,
    page: {
      default: 1,
      type: Number
    },
    limit: {
      default: 10,
      type: Number
    },
    pagecount: {
      default: 10,
      type: Number
    },
    refresh: Function
  },
  methods: {
    isCurrent(i) {
      return i == this.page;
    }
  }
};
</script>

<style scoped>
.current {
  color: red;
}
</style>
