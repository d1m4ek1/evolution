// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '@/router/Authorization/authorization.router';
import 'lazysizes';
import '@/assets/typescript/stickyHeader';
import '@/assets/typescript/websockets';
import MODULE_STICKY_HEADER from "@/assets/typescript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "@/assets/typescript/modules/SignOut.module";

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
      MODULE_SIGN_OUT()
    },
    hideStickyHeader() {
      MODULE_STICKY_HEADER()
    },
    deletePreloader() {
      setTimeout(() => {
        this.preload = false;
      }, 1000);
    },
    changeBackground() {
      const route = this.$route.matched[0].props.default.activetitle;

      const slideFirst: HTMLElement = document.querySelector('.slide-first')
      const slideSecond: HTMLElement = document.querySelector('.slide-second')

      if (route === 'signin') {
        slideFirst.style.left = '0';
        slideSecond.style.left = '100%';
        this.activeTitle.signup = false;
      } else {
        slideFirst.style.left = '-100%';
        slideSecond.style.left = '0';
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
