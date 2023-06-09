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
      <tr v-for="category of categories" :key="category">
        <td>{{ category.id }}</td>
        <td>
          {{ category.name }}
        </td>
        <td>{{ category.created_at }}</td>
        <td>{{ category.updated_at }}</td>
        <td>
          <router-link
            :to="{ name: 'EditCategory', params: { id: category.id } }"
            >Edit</router-link
          >
        </td>
        <td>
          <button-delete
            :text="buttonText"
            @delete-items="$emit('deleteItems', category.id)"
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
  name: "ListCategories",
  components: {
    ButtonDelete,
  },
  props: {
    buttonText: {
      type: String,
      required: true,
      default: "",
    },
    tableHeaders: {
      type: Object,
      required: true,
      default: () => {},
    },
    categories: {
      type: Object,
      required: true,
      default: () => {},
    },
  },
  emits: ["deleteItems"],
});
</script>

<style scoped>
table {
  margin: 0 auto;
}
</style>
