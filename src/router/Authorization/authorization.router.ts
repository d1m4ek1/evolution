import VueRouter from 'vue-router';

// Authorization routings
import SignIn from '@/pages/Authorization/components/SignIn.vue';
import SignUp from '@/pages/Authorization/components/SignUp.vue';

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
