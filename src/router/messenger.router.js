import Messenger from '../pages/Messenger/Messenger.vue';
import FavoritesMessage from '../pages/Messenger/FavoritesMessege.vue';
import MessegeExchange from '../pages/Messenger/MessegeExchange.vue';

export const messenger = [
  {
    path: '/messenger',
    component: Messenger,
  },
  {
    path: '/messenger/favorites',
    component: FavoritesMessage,
  },
  {
    path: '/messenger/chat_:id',
    name: 'chat',
    component: MessegeExchange,
    props: true,
  },
];
