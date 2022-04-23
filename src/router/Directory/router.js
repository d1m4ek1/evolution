import VueRouter from 'vue-router';

// Directory routings
import Explanation from '../../pages/Directory/components/Explanation.vue';

export default new VueRouter({
  routes: [
    {
      path: '/directory',
      component: Explanation,
      props: { paragraph: 'greetings', typeContent: 'hello' },
    },
    {
      path: '/directory/basic-documentation',
      component: Explanation,
      props: { paragraph: 1, typeContent: 'basic' },
    },
    {
      path: '/directory/service-information',
      component: Explanation,
      props: { paragraph: 1, typeContent: 'serviceInformation' },
    },
    {
      path: '/directory/service-founders',
      component: Explanation,
      props: { paragraph: 1, typeContent: 'serviceFounders' },
    },
    {
      path: '/directory/privacy-policy',
      component: Explanation,
      props: { paragraph: 2, typeContent: 'privacyPolicy' },
    },
    {
      path: '/directory/terms-of-use',
      component: Explanation,
      props: { paragraph: 2, typeContent: 'termsOfUs' },
    },
    {
      path: '/directory/profile',
      component: Explanation,
      props: { paragraph: 9, typeContent: 'profile' },
    },
    {
      path: '/directory/commerce',
      component: Explanation,
      props: { paragraph: 9, typeContent: 'commerce' },
    },
    {
      path: '/directory/community',
      component: Explanation,
      props: { paragraph: 9, typeContent: 'community' },
    },
    {
      path: '/directory/settings',
      component: Explanation,
      props: { paragraph: 9, typeContent: 'settings' },
    },
  ],
  mode: 'history',
});
