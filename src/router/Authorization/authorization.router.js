import { createWebHistory, createRouter } from "vue-router";

// Authorization routings
import SignIn from "../../pages/Authorization/components/SignIn.vue";
import SignUp from "../../pages/Authorization/components/SignUp.vue";

const routes = [
  {
    path: "/signin",
    component: SignIn,
    props: { activetitle: "signin" },
  },
  {
    path: "/signup",
    component: SignUp,
    props: { activetitle: "signup" },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  linkActiveClass: "active-title",
  routes,
});
