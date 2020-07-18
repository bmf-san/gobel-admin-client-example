import Vue from "vue";
import VueRouter from "vue-router";
import Signin from "../pages/Signin.vue";
import Home from "../pages/Home.vue";
import Post from "../pages/Post.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/signin",
    name: "Signin",
    component: Signin,
  },
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/posts",
    name: "Post",
    component: Post,
  },
  // {
  //   path: "/posts/create",
  //   name: "NewPost",
  //   component: NewPost,
  // },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
