import VueRouter from "vue-router";

import Shop from "/src/pages/__commerce/Shop/components/Shop.vue"

export default new VueRouter({
    routes: [
        {
            path: "/shop",
            component: Shop,
        },
    ],
    mode: "history",
});
