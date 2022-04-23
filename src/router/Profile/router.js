import VueRouter from "vue-router";

// Autorization routings
import News from "/src/pages/Profile/components/UserNews.vue"
import AboutMe from "/src/pages/Profile/components/UserAboutMe.vue"

export default new VueRouter({
    routes: [
        {
            path: "/:username",
            component: News,
        },
        {
            path: "/:username/aboutMe",
            component: AboutMe,
        },
    ],
    mode: "history",
});