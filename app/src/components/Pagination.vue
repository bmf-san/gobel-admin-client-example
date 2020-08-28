<template>
  <ul>
    <li v-for="i of pagecount" :key="i">
      <router-link
        :to="{ name: name, query: { page: i, limit: limit } }"
        :class="{ current: isCurrent(i) }"
        >{{ i }}</router-link
      >
    </li>
  </ul>
</template>

<script>
export default {
  name: "Pagination",
  props: {
    name: String,
    page: Number,
    limit: Number,
    pagecount: Number,
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
