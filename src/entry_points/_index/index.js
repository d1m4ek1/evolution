// JS IMPORTS
import { createApp } from "vue";
import hideElemDir from "../../assets/javascript/hideElem";
import "lazysizes";
import "../../assets/javascript/websockets";
import StickyHeader from "../../assets/javascript/stickyHeader";

import MODULE_CHECK_AUTHORIZE_USER from "../../assets/javascript/modules/CheckAuthorize.module";
import MODULE_STICKY_HEADER from "../../assets/javascript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "../../assets/javascript/modules/SignOut.module";
import { websocket } from "../../assets/javascript/websockets";

import UiMessageNotif from "../../assets/UIComponents/Notifications/UiMessageNotif.vue";

StickyHeader();

window.Vue = require("vue");

const app = createApp({
  delimiters: ["{%", "%}"],
  data: () => ({
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    chatData: undefined,
    showNotifMessage: false,
  }),
  methods: {
    onCloseNotif() {
      this.chatData = undefined;
      this.showNotifMessage = false;
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

    const connectionToWebsocket = setInterval(() => {
      if (websocket) {
        websocket.addEventListener("message", (event) => {
          const json = JSON.parse(event.data);

          if (json.checked) {
            return;
          }

          this.chatData = json;
          this.showNotifMessage = true;
        });

        clearInterval(connectionToWebsocket);
      }
    }, 500);
  },
  directives: {
    "hide-elem": hideElemDir,
  },
  components: {
    "ui-message-notificate": UiMessageNotif,
  },
});
app.mount("#wrapper");

// (() => new Vue({
//   el: '#wrapper',
//   delimiters: ['{%', '%}'],
//   data: {
//     preload: true,
//     commerce: false,
//     community: false,
//     settings: false,
//   },
//   methods: {
//     signOut() {
//       MODULE_SIGN_OUT();
//     },
//     hideStickyHeader() {
//       MODULE_STICKY_HEADER();
//     },
//     deletePreloader() {
//       setTimeout(() => {
//         this.preload = false;
//       }, 1000);
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
//   created() {
//     this.deletePreloader();
//     MODULE_CHECK_AUTHORIZE_USER();
//   },
// }))();
