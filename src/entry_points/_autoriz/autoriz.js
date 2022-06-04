// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Autorization/router';
// import hideElemDir from '../../assets/js/hideElem';
import 'lazysizes';
import '../../assets/js/stickyHeader';
import '../../assets/js/checkStatus';

window.Vue = require('vue');

Vue.use(VueRouter);

(() => new Vue({
  props: ['activetitle'],
  el: '#wrapper',
  delimiters: ['{%', '%}'],
  data: {
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    activeTitle: {
      signup: false,
    },
  },
  methods: {
    signOut() {
      fetch('/api/signout_account')
        .then((response) => {
          if (response.ok) {
            document.cookie = 'token=; path=/; max-age=-1;';
            document.cookie = 'userId=; path=/; max-age=-1;';
            window.location.href = '/';
          }
        })
        .catch((err) => console.error(err));
    },
    hideStickyHeader() {
      document.querySelector('.header_sticky').style.transform = 'translateX(-200px)';
      document
        .querySelector('.main')
        .classList.remove('main_squeeze_before_add');
      document
        .querySelector('.main')
        .classList.add('main_squeeze_before_remove');
      setTimeout(() => {
        document
          .querySelector('.main')
          .classList.remove('main_squeeze_before_remove');
      }, 490);
    },
    deletePreloader() {
      setTimeout(() => {
        this.preload = false;
      }, 1000);
    },
    changeBackground() {
      const route = this.$route.matched[0].props.default.activetitle;
      if (route === 'signin') {
        document.querySelector('.slide-first').style.left = '0';
        document.querySelector('.slide-second').style.left = '100%';
        this.activeTitle.signup = false;
      } else {
        document.querySelector('.slide-first').style.left = '-100%';
        document.querySelector('.slide-second').style.left = '0';
        this.activeTitle.signup = true;
      }
    },
  },
  computed: {
    dropHeader() {
      return !(document.documentElement.clientWidth <= 960);
    },
    showStickyHeader() {
      return document.documentElement.clientWidth <= 960;
    },
    setDeletePreloader() {
      return {
        deletePreload: true,
      };
    },
    rotateArrowCommunity() {
      return {
        arrow_list_open: this.community,
      };
    },
    rotateArrowCommerce() {
      return {
        arrow_list_open: this.commerce,
      };
    },
    rotateArrowSettings() {
      return {
        arrow_list_open: this.settings,
      };
    },
  },
  updated() {
    this.changeBackground();
  },
  created() {
    this.changeBackground();
    this.deletePreloader();
  },
  router,
}))();
