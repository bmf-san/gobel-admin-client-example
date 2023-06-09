<template>
  <table>
    <thead>
      <tr>
        <th v-for="val of tableHeaders" :key="val">
          {{ val }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="comment of comments" :key="comment">
        <td>{{ comment.id }}</td>
        <td>
          {{ comment.body }}
        </td>
        <td>{{ comment.status }}</td>
        <td>{{ comment.created_at }}</td>
        <td>{{ comment.updated_at }}</td>
        <td>
          <router-link :to="{ name: 'EditComment', params: { id: comment.id } }"
            >Edit</router-link
          >
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "ListComments",
  props: {
    tableHeaders: {
      type: Object,
      required: true,
      default: () => {},
    },
    comments: {
      type: Object,
      required: true,
      default: () => {},
    },
  },
});
</script>

<style scoped>
table {
  margin: 0 auto;
}
</style>
