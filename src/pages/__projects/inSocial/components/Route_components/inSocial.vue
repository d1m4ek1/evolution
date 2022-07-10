<template>
  <div class="messages show_content">
    <section class="title">
      <h1>Все сообщения</h1>
    </section>
    <section class="meesages_content__folders">
      <div class="messages_folders">
        <div class="messages_folder">
          <router-link tag="button" :to="`/inSocial`" class="btn"
            >Все сообщения</router-link
          >
        </div>
        <div class="messages_folder">
          <router-link tag="button" :to="`/inSocial/favorites`" class="btn"
            >Избранные сообщения</router-link
          >
        </div>
      </div>
    </section>
    <section>
      <div class="mesages_content">
        <h2>Чаты</h2>
        <p v-if="chats === undefined">{{ notifications.nullChats }}</p>
      </div>
      <div class="mesages_content">
        <h2>Подписки</h2>
        <template v-if="checkOnNullArray(subscriptions)">
          <div class="message_items" v-for="item in subscriptions">
            <ui-message-item :name="item.name"
                             :net-status="item.netStatus"
                             :banner="item.banner"
                             :logo="item.logo"
                             :userid="item.userId"
                             :is-subscriptions="true"
                             @remove-subs="onRemoveSubs"></ui-message-item>
          </div>
        </template>
        <p v-else>{{ notifications.nullSubscriptions }}</p>
      </div>
    </section>
    <section>
      <div class="mesages_content">
        <h2>Подписчики</h2>
        <template v-if="checkOnNullArray(subscribers)">
          <div class="message_items" v-for="item in subscribers">
            <ui-message-item :name="item.name"
                             :net-status="item.netStatus"
                             :banner="item.banner"
                             :logo="item.logo"
                             :userid="item.userId"
                             :is-subscriptions="false"
                             @remove-subs="onRemoveSubs"></ui-message-item>
          </div>
        </template>
        <p v-else>{{ notifications.nullSubscribers }}</p>
      </div>
    </section>
  </div>
</template>


<script>
import UIMessageItem from '@/pages/__projects/inSocial/components/Ui_components/UiMessageItem';

export default {
  props: {
    subscribers: Array | null,
    subscriptions: Array | null,
    chats: Array | null,
  },
  data() {
    return {
      notifications: {
        nullSubscribers: "На вас пока никто не подписан",
        nullSubscriptions: "Вы не подписаны ни на одного пользователя",
        nullChats: "Вы еще не переписывались"
      }
    }
  },
  methods: {
    onRemoveSubs(data) {
      if (data !== undefined) {
        for (let i = 0; i < this[data.arrayName].length; i++) {
          const subs = this[data.arrayName][i]
          if (subs.userId === data.ident) {
            this[data.arrayName].splice(i, 1)
          }
        }
      }
    },
    checkOnNullArray(array) {
      if (array === null) {
        return false
      }
      if (array.length === 0) {
        return false
      }

      return true
    }
  },
  components: {
    "ui-message-item": UIMessageItem }
}
</script>
