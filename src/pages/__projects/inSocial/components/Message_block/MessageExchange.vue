<template>
  <div class="messages show_content">
    <section class="message_header">
      <router-link tag="button" class="btn" to="/inSocial">Назад</router-link>
      <h1>{{ preData.name }}</h1>
      <div class="message_block_img" :class="preData.netStatus">
        <a href="/1"></a>
        <img class="lazyload" :data-src="setLogo" />
      </div>
    </section>
    <section class="message__main_content">
      <div id="message_view" class="message_view">
        <template v-for="item in messageItems">
          <div :class="defineMessage(item.sender_id, 'main')" :id="item.message_id">
            <div :class="defineMessage(item.sender_id, 'content')">
              <p>{{ item.message }}</p>
              <p class="message_date">Date: {{ item.date }}</p>
            </div>
          </div>
        </template>
        <template v-for="item in newMessageItems">
          <div :class="defineMessage(item.sender_id, 'main')" :id="item.message_id">
            <div class="indicator_new_view_message" :class="defineMessage(item.sender_id, 'content')">
              <p>{{ item.message }}</p>
              <p class="message_date">Date: {{ item.date }}</p>
            </div>
          </div>
        </template>
      </div>
      <div class="message_send">
        <div
          class="write_message"
          placeholder="Ваше сообщение..."
          contenteditable=""
        ></div>
        <div class="control_message">
          <button class="btn send">Отправить</button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import GetCookie from '@/assets/typescript/getCookie';

export default {
  props: {
    chatData: Object,
    name: String,
    logo: String,
    netStatus: String
  },
  data() {
    return {
      preData: {
        name: String,
        logo: String,
        netStatus: String
      },
      defineMyMessage: Number,
      messageItems: Object,
      newMessageItems: Object
    };
  },
  methods: {
    scrollBottom() {
      const messageView = document.getElementById("message_view");
      messageView.scrollTop = messageView.scrollHeight;
    },
    defineMessage(senderId, block) {
      if (senderId === this.defineMyMessage) {
        if (block === "main") return "message_view_me"
        if (block === "content") return "message_view_me__content"
      } else {
        if (block === "main") return "message_view_friend"
        if (block === "content") return "message_view_friend__content"
      }
    }
  },
  computed: {
    setLogo() {
      if (this.preData.logo === 'not_logo.png') {
        return '/user_images/profile/logo/notLogo/not_logo.png'
      }
      return `/user_images/profile/logo/saved/${this.preData.logo}`
    },
  },
  created() {
    if (this.chatData !== undefined) {
      this.defineMyMessage = Number(GetCookie("userId"))

      if (this.chatData.messages.String !== '') {
        this.messageItems = JSON.parse(this.chatData.messages.String)
      }
      if (this.chatData.newMessages.String !== '') {
        this.newMessageItems = JSON.parse(this.chatData.newMessages.String)
      }

      Object.keys(this.preData).forEach(key => {
        this.preData[key] = this[key]
      })
    }
  },
  mounted() {
    this.scrollBottom();
  },
};
</script>
