<template>
  <div class="block_notif" v-if="showNotif">
    <div class="block_notif_header_message">
      <h4>Новое сообщение от {{ name }}</h4>
    </div>
    <div class="block_notif_content_message">
      <div class="img">
        <img src="" alt="Логотип отпровителя" />
      </div>
      <a :href="`/inSocial/chat_${chatData.chatId}`">Читать...</a>
      <a :href="`/${chatData.sender_id}`">Профиль отправителя</a>
    </div>
  </div>
</template>

<style scoped>
.block_notif {
  position: fixed;
  left: 10px;
  bottom: 10px;
  min-width: 100px;
  max-width: 400px;
  height: auto;
  border-radius: 10px;
  border: 1px solid;
  background-color: white;
  padding: 10px;
}
.anim {
  animation-name: slickNotif;
  animation-duration: 1s;
}
.block_notif_content_message {
  display: flex;
  flex-direction: column;
}
.block_notif_content_message a {
  width: max-content;
}
</style>

<script>
export default {
  props: {
    chatData: Object,
    showNotif: false,
  },
  watch: {
    chatData: {
      deep: true,
      handler(value) {
        if (value !== undefined) {
          fetch(`/api/get_user_data_default?sender_id=${value.sender_id}`, {
            method: "GET",
          }).then((response) => {
            response.json().then((data) => {
              this.logo = data.logo;
              this.name = data.name;
            });
          });
        }
      },
    },
  },
  data: () => ({
    logo: undefined,
    name: undefined,
  }),
};
</script>