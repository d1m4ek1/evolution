import VueRouter from 'vue-router';

// Autorization routings
import SignIn from '../../pages/Autorization/components/SignIn.vue';
import SignUp from '../../pages/Autorization/components/SignUp.vue';

export default new VueRouter({
  routes: [
    {
      path: '/signin',
      component: SignIn,
    },
    {
      path: '/signup',
      component: SignUp,
    },
  ],
  mode: 'history',
});
