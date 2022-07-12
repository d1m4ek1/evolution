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
        <p v-if="messageItems.error">{{ messageItems.error }}</p>
        <template v-else>
          <template v-for="item in messageItems">
            <div :class="defineMessage(item.sender_id, 'main')" :id="item.message_id">
              <div :class="defineMessage(item.sender_id, 'content')">
                <p>{{ item.message }}</p>
                <p class="message_date">Date: {{ item.date }}</p>
              </div>
            </div>
          </template>
        </template>
        <template v-if="!newMessageItems.error">
          <template v-for="item in newMessageItems">
            <div :class="defineMessage(item.sender_id, 'main')" :id="item.message_id">
              <div class="indicator_new_view_message" :class="defineMessage(item.sender_id, 'content')">
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
import GetCookie from '@/assets/typescript/getCookie';
import { EventMessageSend, websocket } from '@/assets/typescript/websockets';

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
        name: "",
        logo: "",
        netStatus: "",
      },
      myId: Number,
      friendId: '',
      messageItems: [],
      newMessageItems: [],
      messageListener: {
        message: String
      }
    };
  },
  methods: {
    scrollBottom() {
      const messageView = document.getElementById("message_view");
      messageView.scrollTop = messageView.scrollHeight;
    },
    defineMessage(senderId, block) {
      if (senderId === this.myId) {
        if (block === "main") return "message_view_me"
        if (block === "content") return "message_view_me__content"
      } else {
        if (block === "main") return "message_view_friend"
        if (block === "content") return "message_view_friend__content"
      }
    },
    getDataChat() {
      let chatId = window.location.pathname.split('/')[2].replace("chat_", "")

      if (chatId !== "" && chatId !== undefined && chatId !== null) {
        fetch(`/api/check_chat?chat_id=${chatId}`, {
          method: "GET"
        }).then(response => {
          response.json().then(data => {
            let userDataOne = JSON.parse(data.userDataOne)
            let userDataTwo = JSON.parse(data.userDataTwo)


            if (userDataOne.userId !== this.myId) {
              Object.keys(this.preData).forEach(key => {
                this.preData[key] = userDataOne[key]
              })
              this.friendId = userDataOne.userId
            }
            if (userDataTwo.userId !== this.myId) {
              Object.keys(this.preData).forEach(key => {
                this.preData[key] = userDataTwo[key]
              })
              this.friendId = userDataTwo.userId
            }

            if (data.messages !== null) {
              this.messageItems = JSON.parse(`[${data.messages}]`)
            } else {
              this.messageItems = {error: "Вы никому не писали"}
            }
            if (data.newMessages !== null) {
              this.newMessageItems = JSON.parse(`[${data.newMessages}]`)
            } else {
              this.newMessageItems = {error: "Новых сообщений нет"}
            }
          })
        })
      }
    },
    setLogo() {
      if (this.preData.logo === 'not_logo.png') {
        return '/user_images/profile/logo/notLogo/not_logo.png'
      }
      return `/user_images/profile/logo/saved/${this.preData.logo}`
    },
    getDate() {
      let date = new Date()
      let day = date.getDay() < 10 ? `0${date.getDay()}` : date.getDay()
      let month = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : date.getMonth() + 1
      return `${day}.${month}.${date.getFullYear()}`
    },
    sendMessage() {
      if(this.messageListener.message !== "" && this.messageListener.message !== undefined) {
        if (!Array.isArray(this.newMessageItems)) {
          this.newMessageItems = []
        }
        this.newMessageItems.push({
          sender_id: this.myId,
          message: this.messageListener.message,
          date: this.getDate()
        })
        EventMessageSend(JSON.stringify({
          chatId: window.location.pathname.split('/')[2].replace("chat_", ""),
          sender_id: this.myId,
          recipient_id: this.friendId,
          message: this.messageListener.message,
          date: this.getDate()
        }))

        this.messageListener.message = ""
        document.querySelector(".write_message").innerHTML = ""
      }
    },
    onListenMessage(event) {
      this.messageListener.message = event.target.innerText
      if (event.target.innerText === "\n") {
        document.querySelector(".write_message").innerHTML = ""
      }
    },
    setNewMessages() {
      setTimeout(()=> {
        if (Array.isArray(this.newMessageItems) && this.newMessageItems.length !== 0 || this.newMessageItems.error) {
          if (!Array.isArray(this.messageItems)) {
            this.messageItems = []
          }
          let counter = 0
          for (let i = 0; i < this.newMessageItems.length; i++) {
            if (this.myId !== this.newMessageItems[i].sender_id) {
              this.messageItems.push(this.newMessageItems[i])
              counter++
            }
          }
          if (counter !== 0) {
            let str = []
            for (let i = 0; i < this.newMessageItems.length; i++) {
              str.push(`${JSON.stringify(this.newMessageItems[i])}`)
            }
            EventMessageSend(JSON.stringify({
              isMessageCheck: true,
              message: JSON.stringify(str).replace(/^.|.$/g,""),
              chatId: window.location.pathname.split('/')[2].replace("chat_", ""),
              recipient_id: this.friendId
            }))
            this.newMessageItems = []
          }
        }
      }, 500)
    }
  },
  created() {
    this.myId = Number(GetCookie("userId"))

    if (this.chatData !== undefined) {
      if (this.chatData.messages !== null) {
        this.messageItems = JSON.parse(`[${this.chatData.messages}]`)
      } else {
        this.messageItems = {error: "Вы никому не писали"}
      }
      if (this.chatData.newMessages !== null) {
        this.newMessageItems = JSON.parse(`[${this.chatData.newMessages}]`)
      } else {
        this.newMessageItems = {error: "Новых сообщений нет"}
      }

      Object.keys(this.preData).forEach(key => {
        this.preData[key] = this[key]
      })

      if (this.myId !== this.chatData.userIDOne) {
        this.friendId = this.chatData.userIDOne
      }
      if (this.myId !== this.chatData.userIDTwo) {
        this.friendId = this.chatData.userIDTwo
      }
    } else {
      this.getDataChat()
    }

    websocket.addEventListener("message", (event) => {
      if (event.data === "checked") {
        if (!Array.isArray(this.messageItems)) {
          this.messageItems = []
        }
        this.messageItems.push(...this.newMessageItems)
        this.newMessageItems = []
        return
      }
      if (this.newMessageItems.error) {
        this.newMessageItems = []
      }
      this.newMessageItems.push(JSON.parse(event.data))
    })

    this.setNewMessages()
  },
  mounted() {
    this.scrollBottom();
  },
  updated() {
    this.setNewMessages()
  }
};
</script>
