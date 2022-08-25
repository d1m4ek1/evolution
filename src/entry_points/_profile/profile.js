// JS
import { createApp } from "vue";
import { router } from "../../router/Profile/profile.router";
import hideElemDir from "../../assets/javascript/hideElem";
import "lazysizes";
import "../../assets/javascript/stickyHeader";
import MODULE_CHECK_AUTHORIZE_USER from "../../assets/javascript/modules/CheckAuthorize.module";
import MODULE_STICKY_HEADER from "../../assets/javascript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "../../assets/javascript/modules/SignOut.module";
import StickyHeader from "../../assets/javascript/stickyHeader";
import GetCookie from "../../assets/javascript/getCookie";
import { websocket } from "../../assets/javascript/websockets";
StickyHeader();

import UiMessageNotif from "../../assets/UIComponents/Notifications/UiMessageNotif.vue";

const app = createApp({
  delimiters: ["{%", "%}"],
  data: () => ({
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    aboutme: {
      title: "",
      content: "",
    },
    isSubscriber: false,
    isCountSubscribers: 0,
    tester: String,
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
    getDataProfile() {
      const userId = window.location.pathname.split("/");
      fetch(`/api/get_data_profile?get_data=all&user_id=${userId[1]}`).then(
        (response) => {
          response.json().then((data) => {
            this.aboutme.title = data.aboutmeTitle;
            this.aboutme.content = data.aboutmeContent;
          });
        }
      );
    },
    appendSubscriber(id) {
      fetch(`/api/append_subscriber?append_id=${id}`, {
        method: "POST",
      })
        .then((response) => {
          if (response.ok) {
            this.isSubscriber = true;
            this.getCountSubscriber();
          }
        })
        .catch((error) => console.error(error));
    },
    checkSubcriber() {
      let userId = window.location.pathname.split("/")[1];
      fetch(`/api/check_subscriber?check_id=${userId}`, {
        method: "GET",
      })
        .then((response) => {
          if (!response.ok) {
            console.error(response.statusText);
            return;
          }

          response.json().then((data) => {
            this.isSubscriber = data.isSubscriber;
            this.getCountSubscriber();
          });
        })
        .catch((error) => console.error(error));
    },
    deleteSubscriber(id) {
      fetch(`/api/delete_subscriber?delete_id=${id}`, {
        method: "DELETE",
      })
        .then((response) => {
          if (!response.ok) {
            this.isSubscriber = true;
            console.error(response.statusText);
          } else {
            this.isSubscriber = false;
            this.getCountSubscriber();
          }
        })
        .catch((error) => console.error(error));
    },
    getCountSubscriber() {
      let userId = window.location.pathname.split("/")[1];
      fetch(`/api/count_subscriber?check_id=${userId}`, {
        method: "GET",
      })
        .then((response) => {
          if (!response.ok) {
            console.error(response.statusText);
            return;
          }

          response.json().then((data) => {
            this.isCountSubscribers = data.isCountSubscriber;
          });
        })
        .catch((error) => console.error(error));
    },
    openChat() {
      let userId = GetCookie("userId");
      let userIdTwo = window.location.pathname.split("/")[1];
      if (userId !== undefined) {
        fetch(`/api/check_chat?user_id_two=${userIdTwo}`, {
          method: "GET",
        }).then((response) => {
          response.json().then((data) => {
            window.location.pathname = `/inSocial/chat_${data.chatId}`;
          });
        });
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
    this.getDataProfile();
    this.checkSubcriber();
    this.getCountSubscriber();

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
