// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Directory/directory.router';
import hideElemDir from '../../assets/typescript/hideElem';
import 'lazysizes';
import '../../assets/typescript/stickyHeader';
import MODULE_CHECK_AUTHORIZE_USER from '../../assets/typescript/modules/CheckAuthorize.module';
import MODULE_STICKY_HEADER from "../../assets/typescript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "@/assets/typescript/modules/SignOut.module";
import StickyHeader from "@/assets/typescript/stickyHeader";
StickyHeader()

window.Vue = require('vue');

Vue.use(VueRouter);
Vue.directive('hide-elem', hideElemDir);

function directoryFixed() {
  if (document.querySelector('.directory_fixed')) {
    const directoryFixedBlock = document.querySelector('.directory_fixed');
    let top = Number();

    window.addEventListener('scroll', () => {
      if (window.scrollY <= 30) {
        top = directoryFixedBlock.getBoundingClientRect().top;
      } else {
        top = directoryFixedBlock.getBoundingClientRect().top - 50;
      }
      setTimeout(() => {
        const directoryFixedBlock: HTMLElement = document.querySelector('.directory_fixed');
        directoryFixedBlock.style.transform = `translateY(${top + window.scrollY}px)`;
      }, 300);
    });
  }
}
directoryFixed();

(() => new Vue({
  el: '#wrapper',
  delimiters: ['{%', '%}'],
  data: {
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    showHideDirectory: true,
    showHideComponent: true,
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
    showComponent() {
      if (document.documentElement.clientWidth <= 600) {
        this.showHideDirectory = false;
        this.showHideComponent = true;
      } else {
        this.showHideDirectory = true;
        this.showHideComponent = true;
      }
    },
    hideComponent(data) {
      this.showHideDirectory = data.hide;
      this.showHideComponent = false;
    },
    preShowComponent() {
      if (document.documentElement.clientWidth <= 600) {
        this.showHideComponent = false;
      } else if (
        document.documentElement.clientWidth > 600
      ) {
        this.showHideComponent = true;
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
  created() {
    this.deletePreloader();
    this.preShowComponent();
    MODULE_CHECK_AUTHORIZE_USER();
  },
  router,
}))();
