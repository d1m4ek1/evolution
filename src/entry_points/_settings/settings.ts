// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Settings/settings.router';
import hideElemDir from '../../assets/typescript/hideElem';
import 'lazysizes';
import '../../assets/typescript/stickyHeader';
import MODULE_CHECK_AUTHORIZE_USER from '../../assets/typescript/modules/CheckAuthorize.module';
import MODULE_STICKY_HEADER from "../../assets/typescript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "@/assets/typescript/modules/SignOut.module";
import StickyHeader from "@/assets/typescript/stickyHeader";
StickyHeader()

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
    setTitle(str: string) {
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
    MODULE_CHECK_AUTHORIZE_USER();
  },
  router,
}))();
