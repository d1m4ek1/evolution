// JS IMPORTS
import { createApp, ref } from 'vue';
import { router } from '../router/index.router.js';
import { store } from '../store/store.js';
import hideElemDir from '../assets/javascript/hideElem.js';
import 'lazysizes';
import '../assets/javascript/websockets.js';
import StickyHeader from '../assets/javascript/stickyHeader.js';

import MODULE_CHECK_AUTHORIZE_USER from '../assets/javascript/modules/CheckAuthorize.module.js';
import MODULE_STICKY_HEADER from '../assets/javascript/modules/StickyHeader.module.js';
import MODULE_SIGN_OUT from '../assets/javascript/modules/SignOut.module.js';
import { websocket } from '../assets/javascript/websockets.js';

import UiMessageNotif from '../assets/UIComponents/Notifications/UiMessageNotif.vue';
import GlobalPlayer from '../assets/UIComponents/Players/GlobalPlayer.vue';

StickyHeader();

const app = createApp({
  delimiters: ['{%', '%}'],
  data: () => ({
    preload: true,
    beats: false,
    settings: false,
    chatData: undefined,
    showNotifMessage: false,
    userAuthorized: undefined,
    showMobileHeader: false,
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
    clientWidth() {
      this.showMobileHeader = document.documentElement.clientWidth <= 960;

      window.addEventListener('resize', () => {
        this.showMobileHeader = document.documentElement.clientWidth <= 960;
      });
    },
    getDataSubscriptions() {
      fetch('/api/get_user_card_messages').then((response) => {
        if (!response.ok) {
          console.error(response.statusText);
          return;
        }
        response.json().then((data) => {
          if (!data.isAuthorized) {
            return;
          }
          this.subscribers = data.isCardSubscribers;
          this.subscriptions = data.isCardSubscriptions;
        });
      });
    },
    async getAllChats() {
      const response = await fetch('/api/get_all_chats', {
        method: 'GET',
      });

      const jsonResponse = await response.json();

      if (jsonResponse) {
        for (let i = 0; i < jsonResponse.length; i++) {
          if (!Array.isArray(jsonResponse[i].newMessages)) {
            jsonResponse[i].newMessages = [];
          } else {
            for (let j = 0; j < jsonResponse[i].newMessages.length; j++) {
              const item = jsonResponse[i].newMessages[j];

              jsonResponse[i].newMessages[j] = JSON.parse(item);
            }
          }

          if (!Array.isArray(jsonResponse[i].messages)) {
            jsonResponse[i].messages = [];
          } else {
            for (let j = 0; j < jsonResponse[i].messages.length; j++) {
              const item = jsonResponse[i].messages[j];

              jsonResponse[i].messages[j] = JSON.parse(item);
            }
          }
        }

        this.$store.state.messengerData.chats = jsonResponse;
        for (let i = 0; i < this.$store.state.messengerData.chats.length; i++) {
          const element = this.$store.state.messengerData.chats[i];

          element.userDataOne = JSON.parse(element.userDataOne);
          element.userDataTwo = JSON.parse(element.userDataTwo);
        }

        return true;
      }
    },
    clearNewMessages(iterator, moveItems) {
      if (moveItems) {
        for (let j = 0; j < this.$store.state.messengerData.chats[iterator].newMessages.length; j++) {
          const item = this.$store.state.messengerData.chats[iterator].newMessages[j];
          this.$store.state.messengerData.chats[iterator].messages.push(item);
        }
      }
      this.$store.state.messengerData.chats[iterator].newMessages = [];
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
    rotateArrowBeats() {
      return {
        arrow_list_open: this.beats,
      };
    },
    rotateArrowSettings() {
      return {
        arrow_list_open: this.settings,
      };
    },
  },
  async created() {
    this.deletePreloader();
    this.clientWidth();

    this.getDataSubscriptions();

    await MODULE_CHECK_AUTHORIZE_USER().then(async (response) => {
      this.userAuthorized = response;

      if (response) {
        const responseChats = await this.getAllChats();

        if (responseChats) {
          const setConnectToWebsocket = setInterval(() => {
            if (websocket) {
              websocket.addEventListener('message', (event) => {
                const json = JSON.parse(event.data);

                if (json.checked) {
                  for (let i = 0; i < this.$store.state.messengerData.chats.length; i++) {
                    const chatItem = this.$store.state.messengerData.chats[i];
                    if (Number(chatItem.chatId) === Number(json.chatId)) {
                      this.clearNewMessages(i, true);
                      return;
                    }
                  }
                  return;
                }

                for (let i = 0; i < this.$store.state.messengerData.chats.length; i++) {
                  const chatItem = this.$store.state.messengerData.chats[i];

                  if (Number(chatItem.chatId) === Number(json.chatId)) {
                    delete json.chatId;
                    this.$store.state.messengerData.chats[i].newMessages.push(json);
                    return;
                  }
                }
              });

              clearInterval(setConnectToWebsocket);
            }
          }, 100);
        }
      }
    });
  },
  directives: {
    'hide-elem': hideElemDir,
  },
  components: {
    'ui-message-notificate': UiMessageNotif,
    GlobalPlayer,
  },
});
app.use(store).use(router).mount('#wrapper');
