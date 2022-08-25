import { createWebHistory, createRouter } from "vue-router";

import InSocial from "../../pages/__projects/inSocial/components/Route_components/inSocial.vue";
import FavoritesMessage from "../../pages/__projects/inSocial/components/Route_components/FavoritesMessage.vue";
import MessageExchange from "../../pages/__projects/inSocial/components/Message_block/MessageExchange.vue";

const routes = [
  {
    path: "/inSocial",
    component: InSocial,
  },
  {
    path: "/inSocial/favorites",
    component: FavoritesMessage,
  },
  {
    path: "/inSocial/chat_:id",
    name: "chat",
    component: MessageExchange,
    props: true,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
