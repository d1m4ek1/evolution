import VueRouter from 'vue-router';

import Music from '../../../pages/__community/Music/components/Music.vue';

export default new VueRouter({
  routes: [
    {
      path: '/music',
      component: Music,
    },
  ],
  mode: 'history',
});
