// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Settings/router';
import hideElemDir from '../../assets/js/hideElem';
import '../../assets/js/delayedLoading';
import '../../assets/js/stickyHeader';

window.Vue = require('vue');

Vue.directive('hide-elem', hideElemDir);
Vue.use(VueRouter);

(() => new Vue({
  el: '#wrapper',
  delimiters: ['{%', '%}'],
  data: {
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    title: {
      show: false,
      value: '',
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
      document.querySelector(
        '.header_sticky',
      ).style.transform = 'translateX(-200px)';
      document
        .querySelector('.main')
        .classList.remove('main_squeeze_before_add');
      document
        .querySelector('.main')
        .classList.add('main_squeeze_before_remove');
      document
        .querySelector('.main__body_content')
        .classList.remove('main_squeeze');
      document
        .querySelector('.main__body_content')
        .classList.add('main_unclench');
      document.body.style.overflow = 'unset';
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
    setTitle(str) {
      this.title.show = true;
      this.title.value = str;
      document.title = str;
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
  created() {
    this.deletePreloader();
  },
  router,
}))();
