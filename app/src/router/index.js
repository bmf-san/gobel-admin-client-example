import { createRouter, createWebHistory } from "vue-router";
import Signin from "../pages/Signin.vue";
import Home from "../pages/Home.vue";
import Posts from "../pages/Posts.vue";
import EditPost from "../pages/EditPost.vue";
import CreatePost from "../pages/CreatePost.vue";
import Categories from "../pages/Categories.vue";
import CreateCategory from "../pages/CreateCategory.vue";
import EditCategory from "../pages/EditCategory.vue";
import Tags from "../pages/Tags.vue";
import CreateTag from "../pages/CreateTag.vue";
import EditTag from "../pages/EditTag.vue";
import Comments from "../pages/Comments.vue";
import EditComment from "../pages/EditComment.vue";
import Error from "../pages/Error.vue";
import apiClient from "../modules/apiClient";
import storage from "../storage";

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
    path: "/posts/create",
    name: "CreatePost",
    component: CreatePost,
    meta: { requiresAuth: true }
  },
  {
    path: "/posts/:id",
    name: "EditPost",
    component: EditPost,
    meta: { requiresAuth: true }
  },
  {
    path: "/categories",
    name: "Categories",
    component: Categories,
    meta: { requiresAuth: true }
  },
  {
    path: "/categories/create",
    name: "CreateCategory",
    component: CreateCategory,
    meta: { requiresAuth: true }
  },
  {
    path: "/categories/:id",
    name: "EditCategory",
    component: EditCategory,
    meta: { requiresAuth: true }
  },
  {
    path: "/tags",
    name: "Tags",
    component: Tags,
    meta: { requiresAuth: true }
  },
  {
    path: "/tags/create",
    name: "CreateTag",
    component: CreateTag,
    meta: { requiresAuth: true }
  },
  {
    path: "/tags/:id",
    name: "EditTag",
    component: EditTag,
    meta: { requiresAuth: true }
  },
  {
    path: "/comments",
    name: "Comments",
    component: Comments,
    meta: { requiresAuth: true }
  },
  {
    path: "/comments/:id",
    name: "EditComment",
    component: EditComment,
    meta: { requiresAuth: true }
  },
  {
    path: "/:catchAll(.*)",
    name: "Error",
    component: Error
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // if not signed in
    if (storage.getIsSignin() == false) {
      router.push({ name: "Signin" });
    }
    // private pages
    apiClient
      .get("/private/me", {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken()
        }
      })
      .then(function() {
        next();
      })
      .catch(function(error) {
        console.log(error);
        router.push({ name: "Signin" });
      });
  } else {
    // public pages
    next();
  }
});

export default router;
