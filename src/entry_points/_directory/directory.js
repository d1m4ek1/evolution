// JS
import { createApp } from "vue";
import "lazysizes";

import { router } from "../../router/Directory/directory.router";
import { websocket } from "../../assets/javascript/websockets";
import hideElemDir from "../../assets/javascript/hideElem";
import "../../assets/javascript/stickyHeader";

import MODULE_CHECK_AUTHORIZE_USER from "../../assets/javascript/modules/CheckAuthorize.module";
import MODULE_STICKY_HEADER from "../../assets/javascript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "../../assets/javascript/modules/SignOut.module";

import StickyHeader from "../../assets/javascript/stickyHeader";
StickyHeader();

import UiMessageNotif from "../../assets/UIComponents/Notifications/UiMessageNotif.vue";

function directoryFixed() {
  if (document.querySelector(".directory_fixed")) {
    const directoryFixedBlock = document.querySelector(".directory_fixed");
    let top = Number();

    window.addEventListener("scroll", () => {
      if (window.scrollY <= 30) {
        top = directoryFixedBlock.getBoundingClientRect().top;
      } else {
        top = directoryFixedBlock.getBoundingClientRect().top - 50;
      }
      setTimeout(() => {
        const directoryFixedBlock = document.querySelector(".directory_fixed");
        directoryFixedBlock.style.transform = `translateY(${
          top + window.scrollY
        }px)`;
      }, 300);
    });
  }
}
directoryFixed();

const app = createApp({
  delimiters: ["{%", "%}"],
  data: () => ({
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    showHideDirectory: true,
    showHideComponent: true,
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
      } else if (document.documentElement.clientWidth > 600) {
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
