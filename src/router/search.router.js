import SeacrhPage from '../pages/Search/SearchPage.vue';
import MyAlbums from '../pages/Search/MyAlbums.vue';
import Favorites from '../pages/Search/Favorites.vue';

export const search = [
  {
    path: '/beats',
    component: SeacrhPage,
  },
  {
    path: '/beats/my-albums',
    component: MyAlbums,
  },
  {
    path: '/beats/favorites',
    component: Favorites,
  },
];
