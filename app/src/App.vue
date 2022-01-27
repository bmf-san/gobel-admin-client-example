<template>
  <div>
    <Navigation />
    <router-view />
    <Footer />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Navigation from "./components/Navigation";
import Footer from "./components/Footer";

export default defineComponent({
  name: "App",
  components: {
    Navigation,
    Footer
  }
});
</script>
