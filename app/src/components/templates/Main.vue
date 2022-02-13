<template>
  <Header />
  <slot></slot>
  <Footer />
</template>

<script>
import { defineComponent } from "vue";
import Header from "@/components/organisms/Header";
import Footer from "@/components/organisms/Footer";

export default defineComponent({
  name: "Main",
  components: {
    Header,
    Footer
  }
});
</script>
