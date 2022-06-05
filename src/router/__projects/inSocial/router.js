import VueRouter from 'vue-router';

import InSocial from '../../../pages/__projects/inSocial/components/inSocial.vue';
import FavoritesMessage from '../../../pages/__projects/inSocial/components/FavoritesMessage.vue';
import MessageExchange from '../../../pages/__projects/inSocial/components/MessageExchange.vue';

export default new VueRouter({
  routes: [
    {
      path: '/insocial',
      component: InSocial,
    },
    {
      path: '/insocial/favorites',
      component: FavoritesMessage,
    },
    {
      path: '/insocial/control',
      component: MessageExchange,
    },
  ],
  mode: 'history',
});
