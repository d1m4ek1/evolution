import VueRouter from "vue-router";

import AllMessages from "/src/pages/__community/Messages/components/AllMessages.vue"
import FavoritesMessage from "/src/pages/__community/Messages/components/FavoritesMessage.vue"
import MessageExchange from "/src/pages/__community/Messages/components/MessageExchange.vue"

export default new VueRouter({
    routes: [
        {
            path: "/messages",
            component: AllMessages,
        },
        {
            path: "/messages/favorites",
            component: FavoritesMessage,
        },
        {
            path: "/messages/control",
            component: MessageExchange
        }
    ],
    mode: "history",
});
