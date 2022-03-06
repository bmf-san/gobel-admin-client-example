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
      <tr v-for="post of posts" :key="post">
        <td>{{ post.id }}</td>
        <td>
          {{ post.title }}
        </td>
        <td>
          {{ post.category.name }}
        </td>
        <td>
          {{ post.status }}
        </td>
        <td>{{ post.created_at }}</td>
        <td>{{ post.updated_at }}</td>
        <td>
          <router-link :to="{ name: 'EditPost', params: { id: post.id } }"
            >Edit</router-link
          >
        </td>
        <td>
          <button-delete
            :text="buttonText"
            @delete-items="$emit('deleteItems', post.id)"
          />
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import { defineComponent } from "vue";
import ButtonDelete from "@/components/atoms/ButtonDelete.vue";

export default defineComponent({
  name: "ListPosts",
  components: {
    ButtonDelete
  },
  props: {
    buttonText: {
      type: String,
      required: true,
      default: ""
    },
    tableHeaders: {
      type: Object,
      required: true,
      default: () => {}
    },
    posts: {
      type: Object,
      required: true,
      default: () => {}
    }
  },
  emits: ["deleteItems"]
});
</script>

<style scoped>
table {
  margin: 0 auto;
}
</style>
