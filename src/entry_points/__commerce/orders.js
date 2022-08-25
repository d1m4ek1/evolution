// JS IMPORTS
import { createApp } from "vue";
import "lazysizes";

import { websocket } from "../../assets/javascript/websockets";
import { router } from "../../router/__commerce/orders.router";
import hideElemDir from "../../assets/javascript/hideElem";
import "../../assets/javascript/stickyHeader";

import MODULE_SIGN_OUT from "../../assets/javascript/modules/SignOut.module";
import MODULE_CHECK_AUTHORIZE_USER from "../../assets/javascript/modules/CheckAuthorize.module";

import StickyHeader from "../../assets/javascript/stickyHeader";
StickyHeader();

import UiMessageNotif from "../../assets/UIComponents/Notifications/UiMessageNotif.vue";

const app = createApp({
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
      const headerStickyElement = document.querySelector(".header_sticky");
      const mainElement = document.querySelector(".main");

      headerStickyElement.style.transform = "translateX(-200px)";

      mainElement.classList.remove("main_squeeze_before_add");
      mainElement.classList.add("main_squeeze_before_remove");

      setTimeout(() => {
        mainElement.classList.remove("main_squeeze_before_remove");
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

app.use(router).mount("#wrapper");
