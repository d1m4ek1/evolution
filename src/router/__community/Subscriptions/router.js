import VueRouter from "vue-router";

import Subscriptions from "/src/pages/__community/Subscriptions/components/Subscriptions.vue"

export default new VueRouter({
    routes: [
        {
            path: "/subscriptions",
            component: Subscriptions,
        },
    ],
    mode: "history",
});
