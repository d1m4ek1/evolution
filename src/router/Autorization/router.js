import VueRouter from 'vue-router';

// Autorization routings
import SignIn from '../../pages/Autorization/components/SignIn.vue';
import SignUp from '../../pages/Autorization/components/SignUp.vue';

export default new VueRouter({
  routes: [
    {
      path: '/signin',
      component: SignIn,
      props: { activetitle: 'signin' },
    },
    {
      path: '/signup',
      component: SignUp,
      props: { activetitle: 'signup' },
    },
  ],
  mode: 'history',
  linkActiveClass: 'active-title',
});
