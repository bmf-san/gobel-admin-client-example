<template>
  <Main>
    <div class="container-readable">
      <h1>404</h1>
      <p>Page NotFound</p>
    </div>
  </Main>
</template>

<script>
import { defineComponent } from "vue";
import Main from "@/components/templates/Main";
export default defineComponent({
  name: "NotFound",
  components: {
    Main,
  },
});
</script>
