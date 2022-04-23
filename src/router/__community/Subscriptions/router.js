import VueRouter from 'vue-router';

import Subscriptions from '../../../pages/__community/Subscriptions/components/Subscriptions.vue';

export default new VueRouter({
  routes: [
    {
      path: '/subscriptions',
      component: Subscriptions,
    },
  ],
  mode: 'history',
});
