import { createWebHistory, createRouter } from "vue-router";

import Subscriptions from "../../pages/__commerce/Subscriptions/components/Subscriptions.vue";

const routes = [
  {
    path: "/subscriptions",
    component: Subscriptions,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
