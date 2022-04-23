import VueRouter from "vue-router";

// SETTINGS COMPONENTS
import PageAppearance from "/src/pages/Settings/components/PageAppearance.vue"
import Profile from "/src/pages/Settings/components/Profile.vue"
import PersonalData from "/src/pages/Settings/components/PersonalData.vue"
import Shop from "/src/pages/Settings/components/Shop.vue"

export default new VueRouter({
    routes: [
        {
            path: "/customize",
            component: Profile,
        },
        {
            path: "/customize/shop",
            component: Shop,
        },
        {
            path: "/customize/page-appearance",
            component: PageAppearance,
        },
        {
            path: "/customize/personal-data",
            component: PersonalData,
        },
    ],
    mode: "history",
});