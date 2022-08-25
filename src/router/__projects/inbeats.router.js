import { createWebHistory, createRouter } from "vue-router";

import InBeats from "../../pages/__projects/inBeats/components/inBeats.vue";
import UserPage from "../../pages/__projects/inBeats/components/UserPage.vue";

const routes = [
  {
    path: "/inBeats",
    component: InBeats,
  },
  {
    path: "/inBeats/user_:id",
    component: UserPage,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
