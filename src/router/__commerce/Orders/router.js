import VueRouter from 'vue-router';

import Orders from '../../../pages/__commerce/Orders/components/Orders.vue';

export default new VueRouter({
  routes: [
    {
      path: '/orders',
      component: Orders,
    },
  ],
  mode: 'history',
});
