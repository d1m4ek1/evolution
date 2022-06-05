import VueRouter from 'vue-router';

import InBeats from '../../../pages/__projects/inBeats/components/inBeats.vue';

export default new VueRouter({
  routes: [
    {
      path: '/inbeats',
      component: InBeats,
    },
  ],
  mode: 'history',
});
