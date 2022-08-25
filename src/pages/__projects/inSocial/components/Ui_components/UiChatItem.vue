<template>
  <div
    class="message_item"
    :class="{ new_message_animation: chatDataItem.newMessages.length !== 0 }"
    :id="`message_item_${chatDataItem.chatId}`"
  >
    <div class="message_item__info">
      <a @click="openChat()"></a>
      <div class="message_item__logo" :class="dataFriend.netStatus">
        <img class="lazyload" :data-src="setLogo" />
      </div>
      <div class="message_item__name">
        <h2>Чат с {{ dataFriend.name }}</h2>
      </div>
      <div class="message_item__control">
        <button class="btn">В избранное</button>
        <button class="btn">Удалить чат</button>
      </div>
    </div>
    <div class="message_item__banner">
      <img class="lazyload" :data-src="setBanner" />
    </div>
  </div>
</template>

<script>
import GetCookie from "../../../../../assets/javascript/getCookie.js";

export default {
  props: {
    chatDataItem: Object,
  },
  data() {
    return {
      myId: Number,
      dataFriend: {
        id: "",
        name: "",
        logo: "",
        banner: "",
        netStatus: "",
        chatData: Object,
      },
      friendId: 0,
    };
  },
  methods: {
    initDataFriend() {
      if (this.chatDataItem.userIDOne !== this.myId) {
        let jsonUser = this.chatDataItem.userDataOne;
        this.dataFriend = {
          id: this.chatDataItem.userIDOne,
          name: jsonUser.name,
          logo: jsonUser.logo,
          banner: jsonUser.banner,
        };
        this.friendId = this.chatDataItem.userIDOne;
      }
      if (this.chatDataItem.userIDTwo !== this.myId) {
        let jsonUser = this.chatDataItem.userDataTwo;
        this.dataFriend = {
          id: this.chatDataItem.userIDTwo,
          name: jsonUser.name,
          logo: jsonUser.logo,
          banner: jsonUser.banner,
        };
        this.friendId = this.chatDataItem.userIDTwo;
      }
    },
    openChat() {
      let l = this.chatDataItem;
      this.$router.push({
        name: "chat",
        params: {
          id: this.chatDataItem.chatId,
          chatDataId: this.chatDataItem.chatId,
          name: this.dataFriend.name,
          logo: this.dataFriend.logo,
        },
      });
    },
  },
  computed: {
    setLogo() {
      if (this.dataFriend.logo !== "") {
        if (this.dataFriend.logo === "not_logo.png") {
          return "/user_files/profile/logo/notLogo/not_logo.png";
        }
        return `/user_files/profile/logo/saved/${this.dataFriend.logo}`;
      }
    },
    setBanner() {
      if (this.dataFriend.banner !== "") {
        if (this.dataFriend.banner === "not_banner.png") {
          return "/user_files/profile/banner/notBanner/not_banner.png";
        }
        return `/user_files/profile/banner/saved/${this.dataFriend.banner}`;
      }
    },
  },
  created() {
    this.myId = Number(GetCookie("userId"));
    this.initDataFriend();
  },
};
</script>
