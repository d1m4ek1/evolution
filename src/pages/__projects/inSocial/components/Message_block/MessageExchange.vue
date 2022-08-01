<template>
  <div class="messages show_content">
    <section class="message_header">
      <router-link tag="button" class="btn" to="/inSocial">Назад</router-link>
      <h1>{{ preData.name }}</h1>
      <div class="message_block_img" :class="preData.netStatus">
        <a :href="`/${friendId}`"></a>
        <img :src="setLogo()" />
      </div>
    </section>
    <section class="message__main_content">
      <div id="message_view" class="message_view">
        <template
          v-if="$store.state.chatData !== undefined && arrayIdChat !== null"
        >
          <template
            v-for="(item, idx) in $store.state.chatData[arrayIdChat].messages"
            :key="idx + 'def_message'"
          >
            <div
              :class="defineMessage(item.sender_id, 'main')"
              :id="item.message_id"
            >
              <div :class="defineMessage(item.sender_id, 'content')">
                <p>{{ item.message }}</p>
                <p class="message_date">Date: {{ item.date }}</p>
              </div>
            </div>
          </template>
          <template
            v-for="(item, idx) in $store.state.chatData[arrayIdChat]
              .newMessages"
            :key="idx + 'new_message'"
          >
            <div
              :class="defineMessage(item.sender_id, 'main')"
              :id="item.message_id"
            >
              <div
                class="indicator_new_view_message"
                :class="defineMessage(item.sender_id, 'content')"
              >
                <p>{{ item.message }}</p>
                <p class="message_date">Date: {{ item.date }}</p>
              </div>
            </div>
          </template>
        </template>
      </div>
      <div class="message_send">
        <div
          @input="onListenMessage"
          class="write_message placeholder_block"
          contenteditable
        ></div>
        <div class="control_message">
          <button @click="sendMessage()" class="btn send">Отправить</button>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.message_block_img img {
  height: 100%;
}
.placeholder_block:empty:before {
  content: "Ваше сообщение...";
}
</style>

<script>
import GetCookie from "../../../../../assets/javascript/getCookie.js";
import {
  EventMessageSend,
  websocket,
} from "../../../../../assets/javascript/websockets.js";

export default {
  props: {
    chatDataId: String,
    name: String,
    logo: String,
    netStatus: String,
  },
  data() {
    return {
      preData: {
        name: "",
        logo: "",
        netStatus: "",
      },
      myId: Number,
      friendId: "",
      chatIdParsed: 0,
      arrayIdChat: null,
      messageListener: {
        message: String,
      },
      dataSetted: false,
    };
  },
  watch: {
    "$store.getters.getChatData": {
      deep: true,
      handler() {
        if (this.$store.getters.getChatData !== undefined) {
          if (!this.dataSetted) {
            this.setStartMessage();
          }
          this.setNewMessages();
        }
      },
    },
  },
  methods: {
    scrollBottom() {
      const messageView = document.getElementById("message_view");
      messageView.scrollTop = messageView.scrollHeight;
    },
    defineMessage(senderId, block) {
      if (senderId === this.myId) {
        if (block === "main") return "message_view_me";
        if (block === "content") return "message_view_me__content";
      } else {
        if (block === "main") return "message_view_friend";
        if (block === "content") return "message_view_friend__content";
      }
    },
    setLogo() {
      if (this.preData.logo === "not_logo.png") {
        return "/user_images/profile/logo/notLogo/not_logo.png";
      }
      return `/user_images/profile/logo/saved/${this.preData.logo}`;
    },
    getDate() {
      const date = new Date();
      const day = date.getDay() < 10 ? `0${date.getDay()}` : date.getDay();
      const month =
        date.getMonth() + 1 < 10
          ? `0${date.getMonth() + 1}`
          : date.getMonth() + 1;
      return `${day}.${month}.${date.getFullYear()}`;
    },
    sendMessage() {
      if (
        this.messageListener.message !== "" &&
        this.messageListener.message !== undefined
      ) {
        this.$store.state.chatData[this.arrayIdChat].newMessages.push({
          sender_id: this.myId,
          message: this.messageListener.message,
          date: this.getDate(),
        });

        EventMessageSend(
          JSON.stringify({
            chatId: this.$attrs.id,
            sender_id: this.myId,
            recipient_id: this.friendId,
            message: this.messageListener.message,
            date: this.getDate(),
          })
        );
        this.messageListener.message = "";
        document.querySelector(".write_message").innerHTML = "";
      }
    },
    onListenMessage(event) {
      this.messageListener.message = event.target.innerText;
      if (event.target.innerText === "\n") {
        document.querySelector(".write_message").innerHTML = "";
      }
    },
    setNewMessages() {
      const chatItem = this.$store.state.chatData[this.arrayIdChat];
      if (chatItem.newMessages.length !== 0) {
        let counter = 0;
        for (let i = 0; i < chatItem.newMessages.length; i++) {
          if (this.myId !== chatItem.newMessages[i].sender_id) {
            counter++;
          }
        }

        if (counter !== 0) {
          const str = [];
          for (let i = 0; i < chatItem.newMessages.length; i++) {
            str.push(`${JSON.stringify(chatItem.newMessages[i])}`);
          }
          EventMessageSend(
            JSON.stringify({
              isMessageCheck: true,
              message: JSON.stringify(str).replace(/^.|.$/g, ""),
              chatId: this.$attrs.id,
              recipient_id: this.friendId,
            })
          );

          chatItem.messages.push(...chatItem.newMessages);
          chatItem.newMessages = [];

          this.dataSetted = true;
        }
      }
    },
    setStartMessage() {
      for (let i = 0; i < this.$store.state.chatData.length; i++) {
        const element = this.$store.state.chatData[i];

        if (element.chatId === this.chatIdParsed) {
          this.arrayIdChat = i;
          break;
        }
      }

      Object.keys(this.preData).forEach((key) => {
        this.preData[key] = this[key];
      });

      const chatItem = this.$store.state.chatData[this.arrayIdChat];

      if (this.myId !== chatItem.userIDOne) {
        this.friendId = chatItem.userIDOne;

        this.preData.logo = chatItem.userDataOne.logo;
        this.preData.name = chatItem.userDataOne.name;
      }
      if (this.myId !== chatItem.userIDTwo) {
        this.friendId = chatItem.userIDTwo;

        this.preData.logo = chatItem.userDataTwo.logo;
        this.preData.name = chatItem.userDataTwo.name;
      }
    },
  },
  created() {
    this.myId = Number(GetCookie("userId"));
    this.chatIdParsed = Number(this.$attrs.id);

    if (this.name !== undefined && this.logo !== undefined) {
      this.setStartMessage();
      this.dataSetted = true;
    }
    if (this.$store.getters.chatData !== undefined) {
      this.setNewMessages();
    }
  },
  mounted() {
    this.scrollBottom();
  },
};
</script>
