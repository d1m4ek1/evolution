import VueRouter from 'vue-router';

import AllMessages from '../../../pages/__community/Messages/components/AllMessages.vue';
import FavoritesMessage from '../../../pages/__community/Messages/components/FavoritesMessage.vue';
import MessageExchange from '../../../pages/__community/Messages/components/MessageExchange.vue';

export default new VueRouter({
  routes: [
    {
      path: '/messages',
      component: AllMessages,
    },
    {
      path: '/messages/favorites',
      component: FavoritesMessage,
    },
    {
      path: '/messages/control',
      component: MessageExchange,
    },
  ],
  mode: 'history',
});
