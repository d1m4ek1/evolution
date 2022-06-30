// JS IMPORTS
import Vue from 'vue';
import VueRouter from 'vue-router';
import 'lazysizes';
import router from '../../router/__commerce/orders.router';
import hideElemDir from '../../assets/typescript/hideElem';
import '../../assets/typescript/stickyHeader';
import MODULE_SIGN_OUT from "@/assets/typescript/modules/SignOut.module";

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
  },
  methods: {
    signOut() {
      MODULE_SIGN_OUT()
    },
    hideStickyHeader() {
      const headerStickyElement: HTMLElement = <HTMLElement>document.querySelector('.header_sticky')
      const mainElement: HTMLElement = <HTMLElement>document.querySelector('.main')

      headerStickyElement.style.transform = 'translateX(-200px)';

      mainElement.classList.remove('main_squeeze_before_add');
      mainElement.classList.add('main_squeeze_before_remove');

      setTimeout(() => {
        mainElement.classList.remove('main_squeeze_before_remove');
      }, 490);
    },
    deletePreloader() {
      setTimeout(() => {
        this.preload = false;
      }, 1000);
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
