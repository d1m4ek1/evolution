// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Profile/profile.router';
import hideElemDir from '../../assets/typescript/hideElem';
import 'lazysizes';
import '../../assets/typescript/stickyHeader';
import MODULE_CHECK_AUTHORIZE_USER from '../../assets/typescript/modules/CheckAuthorize.module';
import MODULE_STICKY_HEADER from "../../assets/typescript/modules/StickyHeader.module";
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
    aboutme: {
      title: '',
      content: '',
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
    getDataProfile() {
      const userId = window.location.pathname.split('/');
      fetch(`/api/get_data_profile?get_data=all&user_id=${userId[1]}`).then(
        (response) => {
          response.json().then((data) => {
            this.aboutme.title = data.aboutmeTitle;
            this.aboutme.content = data.aboutmeContent;
          });
        },
      );
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
    this.getDataProfile();
    MODULE_CHECK_AUTHORIZE_USER();
  },
  router,
}))();
