import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../pages/Home.vue";
import Post from "../pages/Post.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/posts",
    name: "Post",
    component: Post,
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
