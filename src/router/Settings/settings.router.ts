import VueRouter from 'vue-router';

// SETTINGS COMPONENTS
import PageAppearance from '@/pages/Settings/components/PageAppearance.vue';
import Profile from '@/pages/Settings/components/Profile.vue';
import PersonalData from '@/pages/Settings/components/PersonalData.vue';
import Shop from '@/pages/Settings/components/Shop.vue';


export default new VueRouter({
  routes: [
    {
      path: '/customize',
      component: Profile,
    },
    {
      path: '/customize/shop',
      component: Shop,
    },
    {
      path: '/customize/page-appearance',
      component: PageAppearance,
    },
    {
      path: '/customize/personal-data',
      component: PersonalData,
    },
  ],
  mode: 'history',
});
