// SETTINGS COMPONENTS
import Settings from '../pages/Settings/Settings.vue';

export const settings = [
  {
    path: '/customize',
    component: Settings,
    props: { activeCustomize: 'profile', title: 'Настройки профиля' },
  },
  {
    path: '/customize/page-appearance',
    component: Settings,
    props: { activeCustomize: 'page-appearance', title: 'Настройки вида приложения' },
  },
  {
    path: '/customize/personal-data',
    component: Settings,
    props: { activeCustomize: 'personal-data', title: 'Настройки личных данных' },
  },
];
