import { createWebHistory, createRouter } from "vue-router";

import Orders from "../../pages/__commerce/Orders/components/Orders.vue";

const routes = [
  {
    path: "/orders",
    component: Orders,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
