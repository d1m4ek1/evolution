import VueRouter from 'vue-router';

import InMusic from '../../../pages/__projects/inMusic/components/inMusic.vue';

export default new VueRouter({
  routes: [
    {
      path: '/inmusic',
      component: InMusic,
    },
  ],
  mode: 'history',
});
