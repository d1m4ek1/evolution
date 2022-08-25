import { createWebHistory, createRouter } from "vue-router";

// Authorization routings
import News from "../../pages/Profile/components/UserNews.vue";
import AboutMe from "../../pages/Profile/components/UserAboutMe.vue";

const routes = [
  {
    path: "/:username",
    component: News,
  },
  {
    path: "/:username/aboutMe",
    component: AboutMe,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
