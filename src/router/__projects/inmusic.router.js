import { createWebHistory, createRouter } from "vue-router";

import InMusic from "../../pages/__projects/inMusic/components/inMusic.vue";

const routes = [
  {
    path: "/inMusic",
    component: InMusic,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
