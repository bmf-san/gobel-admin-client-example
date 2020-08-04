import Vue from "vue";
import VueRouter from "vue-router";
import Signin from "../pages/Signin.vue";
import Home from "../pages/Home.vue";
import Posts from "../pages/Posts.vue";
import EditPost from "../pages/EditPost.vue";
import apiClient from "../modules/apiClient";

Vue.use(VueRouter);

const routes = [
  {
    path: "/signin",
    name: "Signin",
    component: Signin
  },
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: { requiresAuth: true }
  },
  {
    path: "/posts",
    name: "Posts",
    component: Posts,
    meta: { requiresAuth: true }
  },
  {
    path: "/posts/:id",
    name: "EditPost",
    component: EditPost,
    meta: { requiresAuth: true }
  }
  // {
  //   path: "/posts/create",
  //   name: "NewPost",
  //   component: NewPost,
  // meta: { requiresAuth: true }
  // },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    apiClient
      .get("/private/me", {
        headers: {
          'Authorization': "Bearer " + localStorage.getItem("access_token"),
        }
      })
      .then(function() {
        // TODO: Get admin data from response and set it to state or something for output admin data to views. Maybe this case needs store pattern.
        next();
      })
      .catch(function(error) {
        console.log(error);
        next({
          path: "/signin",
        });
      });
  } else {
    next()
  }
});

export default router;
