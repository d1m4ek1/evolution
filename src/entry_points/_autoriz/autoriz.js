// JS
import { createApp } from "vue";
// import VueRouter from 'vue-router';
import { router } from "../../router/Authorization/authorization.router";
import "lazysizes";
import "../../assets/javascript/stickyHeader";
import "../../assets/javascript/websockets";
import MODULE_STICKY_HEADER from "../../assets/javascript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "../../assets/javascript/modules/SignOut.module";
import StickyHeader from "../../assets/javascript/stickyHeader";
StickyHeader();

// window.Vue = require('vue');

// Vue.use(VueRouter);

const app = createApp({
  delimiters: ["{%", "%}"],
  props: {
    activetitle: String,
  },
  data: () => ({
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    activeTitle: {
      signup: false,
    },
    activeNumber: 0,
  }),
  methods: {
    onChangenBackground(data) {
      this.activeNumber = data.backgroundNumber;
      this.changeBackground();
    },
    signOut() {
      MODULE_SIGN_OUT();
    },
    hideStickyHeader() {
      MODULE_STICKY_HEADER();
    },
    deletePreloader() {
      setTimeout(() => {
        this.preload = false;
      }, 1000);
    },
    changeBackground() {
      const route = window.location.pathname.split("/")[1];

      const slideFirst = document.querySelector(".slide-first");
      const slideSecond = document.querySelector(".slide-second");

      if (route === "signin" || this.activeNumber === 1) {
        slideFirst.style.left = "0";
        slideSecond.style.left = "100%";
        this.activeTitle.signup = false;
      }
      if (route === "signun" || this.activeNumber === 2) {
        slideFirst.style.left = "-100%";
        slideSecond.style.left = "0";
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
  created() {
    this.deletePreloader();
  },
  mounted() {
    this.changeBackground();
  },
});

app.use(router).mount("#wrapper");

// (() => new Vue({
//   props: ['activetitle'],
//   el: '#wrapper',
//   delimiters: ['{%', '%}'],
//   data: {
//     preload: true,
//     commerce: false,
//     community: false,
//     settings: false,
//     activeTitle: {
//       signup: false,
//     },
//   },
//   methods: {
//     signOut() {
//       MODULE_SIGN_OUT()
//     },
//     hideStickyHeader() {
//       MODULE_STICKY_HEADER()
//     },
//     deletePreloader() {
//       setTimeout(() => {
//         this.preload = false;
//       }, 1000);
//     },
//     changeBackground() {
//       const route = this.$route.matched[0].props.default.activetitle;

//       const slideFirst = document.querySelector('.slide-first')
//       const slideSecond = document.querySelector('.slide-second')

//       if (route === 'signin') {
//         slideFirst.style.left = '0';
//         slideSecond.style.left = '100%';
//         this.activeTitle.signup = false;
//       } else {
//         slideFirst.style.left = '-100%';
//         slideSecond.style.left = '0';
//         this.activeTitle.signup = true;
//       }
//     },
//   },
//   computed: {
//     dropHeader() {
//       return !(document.documentElement.clientWidth <= 960);
//     },
//     showStickyHeader() {
//       return document.documentElement.clientWidth <= 960;
//     },
//     setDeletePreloader() {
//       return {
//         deletePreload: true,
//       };
//     },
//     rotateArrowCommunity() {
//       return {
//         arrow_list_open: this.community,
//       };
//     },
//     rotateArrowCommerce() {
//       return {
//         arrow_list_open: this.commerce,
//       };
//     },
//     rotateArrowSettings() {
//       return {
//         arrow_list_open: this.settings,
//       };
//     },
//   },
//   updated() {
//     this.changeBackground();
//   },
//   created() {
//     this.changeBackground();
//     this.deletePreloader();
//   },
//   router,
// }))();
