// Profile routings
import Profile from '../pages/Profile/Profile.vue';

export const profile = [
  {
    path: '/:username',
    component: Profile,
    props: {
      loadComponent: 'news',
    },
  },
  {
    path: '/:username/aboutMe',
    component: Profile,
    props: {
      loadComponent: 'aboutme',
    },
  },
];
