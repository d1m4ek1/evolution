import Directory from '../pages/Directory/Directory.vue';

export const directory = [
  {
    path: '/directory',
    component: Directory,
    props: { paragraph: 'greetings', typeContent: 'hello' },
  },
  {
    path: '/directory/basic-documentation',
    component: Directory,
    props: { paragraph: 1, typeContent: 'basic' },
  },
  {
    path: '/directory/service-information',
    component: Directory,
    props: { paragraph: 1, typeContent: 'serviceInformation' },
  },
  {
    path: '/directory/service-founders',
    component: Directory,
    props: { paragraph: 1, typeContent: 'serviceFounders' },
  },
  {
    path: '/directory/privacy-policy',
    component: Directory,
    props: { paragraph: 2, typeContent: 'privacyPolicy' },
  },
  {
    path: '/directory/terms-of-use',
    component: Directory,
    props: { paragraph: 2, typeContent: 'termsOfUs' },
  },
  {
    path: '/directory/profile',
    component: Directory,
    props: { paragraph: 9, typeContent: 'profile' },
  },
  {
    path: '/directory/commerce',
    component: Directory,
    props: { paragraph: 9, typeContent: 'commerce' },
  },
  {
    path: '/directory/community',
    component: Directory,
    props: { paragraph: 9, typeContent: 'community' },
  },
  {
    path: '/directory/settings',
    component: Directory,
    props: { paragraph: 9, typeContent: 'settings' },
  },
];
