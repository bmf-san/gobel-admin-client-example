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
      <tr v-for="tag of tags" :key="tag">
        <td>{{ tag.id }}</td>
        <td>
          {{ tag.name }}
        </td>
        <td>{{ tag.created_at }}</td>
        <td>{{ tag.updated_at }}</td>
        <td>
          <router-link :to="{ name: 'EditTag', params: { id: tag.id } }"
            >Edit</router-link
          >
        </td>
        <td>
          <button-delete
            :text="buttonText"
            @delete-items="$emit('deleteItems', tag.id)"
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
  name: "ListTags",
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
    tags: {
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
