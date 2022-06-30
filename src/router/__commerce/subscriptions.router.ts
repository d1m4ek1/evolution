import VueRouter from 'vue-router';

import Subscriptions from '@/pages/__commerce/Subscriptions/components/Subscriptions.vue';

export default new VueRouter({
  routes: [
    {
      path: '/subscriptions',
      component: Subscriptions,
    },
  ],
  mode: 'history',
});
