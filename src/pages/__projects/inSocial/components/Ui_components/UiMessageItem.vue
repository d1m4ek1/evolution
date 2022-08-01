<template>
  <div class="message_item" :id="`message_item_${userid}`">
    <div class="message_item__info">
      <div class="message_item__logo" :class="netStatus">
        <img class="lazyload" :data-src="setLogo" />
      </div>
      <div class="message_item__name">
        <h2>{{ name }}</h2>
      </div>
      <div class="message_item__control">
        <button @click="openChat(userid)" class="btn">Сообщение</button>

        <button @click="openProfile(userid)" class="btn">Профиль</button>

        <button
          v-if="isSubscriptions"
          @click="deleteSubscriber(userid)"
          class="btn"
        >
          Отписаться
        </button>
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
    userid: Number,
    name: String,
    logo: String,
    banner: String,
    netStatus: String,
    isSubscriptions: Boolean,
  },
  methods: {
    deleteSubscriber(id) {
      fetch(`/api/delete_subscriber?delete_id=${id}`, {
        method: "DELETE",
      })
        .then((response) => {
          if (!response.ok) {
            console.error(response.statusText);
          } else {
            if (this.isSubscriptions) {
              this.$emit("remove-subs", {
                ident: id,
                arrayName: "subscriptions",
              });
            } else {
              this.$emit("remove-subs", {
                ident: id,
                arrayName: "subscribers",
              });
            }
          }
        })
        .catch((error) => console.error(error));
    },
    openProfile(id) {
      window.location.pathname = id;
    },
    openChat(id) {
      let userId = GetCookie("userId");
      if (userId !== undefined) {
        fetch(`/api/check_chat?user_id_two=${id}`, {
          method: "GET",
        }).then((response) => {
          response.json().then((data) => {
            this.$router.push({
              name: "chat",
              params: {
                id: data.chatId,
                chatData: data,
                name: this.name,
                logo: this.logo,
                netStatus: this.netStatus,
              },
            });
          });
        });
      }
    },
  },
  computed: {
    setLogo() {
      if (this.logo === "not_logo.png") {
        return "/user_images/profile/logo/notLogo/not_logo.png";
      }
      return `/user_images/profile/logo/saved/${this.logo}`;
    },
    setBanner() {
      if (this.banner === "not_banner.png") {
        return "/user_images/profile/banner/notBanner/not_banner.png";
      }
      return `/user_images/profile/banner/saved/${this.banner}`;
    },
  },
};
</script>
