import VueRouter from 'vue-router';

// Authorization routings
import News from '@/pages/Profile/components/UserNews.vue';
import AboutMe from '@/pages/Profile/components/UserAboutMe.vue';

export default new VueRouter({
  routes: [
    {
      path: '/:username',
      component: News,
    },
    {
      path: '/:username/aboutMe',
      component: AboutMe,
    },
  ],
  mode: 'history',
});
