// JS IMPORTS
import { createApp } from "vue";
// import VueRouter from 'vue-router';
import { router } from "../../router/__projects/insocial.router";
import hideElemDir from "../../assets/javascript/hideElem";
import "lazysizes";
import "../../assets/javascript/stickyHeader";
import MODULE_STICKY_HEADER from "../../assets/javascript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "../../assets/javascript/modules/SignOut.module";
import MODULE_CHECK_AUTHORIZE_USER from "../../assets/javascript/modules/CheckAuthorize.module";
import StickyHeader from "../../assets/javascript/stickyHeader";
import { websocket } from "../../assets/javascript/websockets";
import { store } from "../../store/store";

StickyHeader();

const app = createApp({
  delimiters: ["{%", "%}"],
  data: () => ({
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    subscribers: [],
    subscriptions: [],
  }),
  methods: {
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
    getDataSubscriptions() {
      fetch("/api/get_user_card_messages").then((response) => {
        if (!response.ok) {
          console.error(response.statusText);
          return;
        }
        response.json().then((data) => {
          this.subscribers = data.isCardSubscribers;
          this.subscriptions = data.isCardSubscriptions;
        });
      });
    },
    getAllChats() {
      fetch("/api/get_all_chats", {
        method: "GET",
      }).then((response) => {
        response.json().then((data) => {
          const chatData = data;
          for (let i = 0; i < chatData.length; i++) {
            if (!Array.isArray(chatData[i].newMessages)) {
              chatData[i].newMessages = [];
            } else {
              for (let j = 0; j < chatData[i].newMessages.length; j++) {
                const item = chatData[i].newMessages[j];

                chatData[i].newMessages[j] = JSON.parse(item);
              }
            }

            if (!Array.isArray(chatData[i].messages)) {
              chatData[i].messages = [];
            } else {
              for (let j = 0; j < chatData[i].messages.length; j++) {
                const item = chatData[i].messages[j];

                chatData[i].messages[j] = JSON.parse(item);
              }
            }
          }

          this.$store.state.chatData = data;
          for (let i = 0; i < this.$store.state.chatData.length; i++) {
            const element = this.$store.state.chatData[i];

            element.userDataOne = JSON.parse(element.userDataOne);
            element.userDataTwo = JSON.parse(element.userDataTwo);
          }
        });
      });
    },
    clearNewMessages(iterator, moveItems) {
      if (moveItems) {
        for (
          let j = 0;
          j < this.$store.state.chatData[iterator].newMessages.length;
          j++
        ) {
          const item = this.$store.state.chatData[iterator].newMessages[j];
          this.$store.state.chatData[iterator].messages.push(item);
        }
      }
      this.$store.state.chatData[iterator].newMessages = [];
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
    this.getDataSubscriptions();
    MODULE_CHECK_AUTHORIZE_USER();
    this.getAllChats();

    const connectionToWebsocket = setInterval(() => {
      if (websocket) {
        websocket.addEventListener("message", (event) => {
          const json = JSON.parse(event.data);

          if (json.checked) {
            for (let i = 0; i < this.$store.state.chatData.length; i++) {
              const chatItem = this.$store.state.chatData[i];
              if (Number(chatItem.chatId) === Number(json.chatId)) {
                this.clearNewMessages(i, true);
                return;
              }
            }
            return;
          }

          for (let i = 0; i < this.$store.state.chatData.length; i++) {
            const chatItem = this.$store.state.chatData[i];

            if (Number(chatItem.chatId) === Number(json.chatId)) {
              delete json.chatId;
              this.$store.state.chatData[i].newMessages.push(json);
              return;
            }
          }
        });

        clearInterval(connectionToWebsocket);
      }
    }, 500);
  },
  directives: {
    "hide-elem": hideElemDir,
  },
});

app.use(store).use(router).mount("#wrapper");
