<template>
  <div class="messages show_content">
    <section class="message_header">
      <router-link tag="button" class="btn" to="/messages"
        >Назад</router-link
      >
      <h1>Name</h1>
      <div class="message_block_img online">
        <a href="/1"></a>
        <img src="/profile/logo/notLogo/not_logo.png" />
      </div>
    </section>
    <section class="message__main_content">
      <div id="message_view" class="message_view">
        <div class="message_view_friend">
          <div class="message_view_friend__content">
            <p>
              Lorem, ipsum dolor sit amet consectetur adipisicing elit. Non,
              veniam?
            </p>
            <p class="message_date">Date: 10.12.2016</p>
          </div>
        </div>
        <div class="message_view_me">
          <div class="message_view_me__content">
            <p>Lorem, ipsum dolor sit amet consectetur?</p>
            <p class="message_date">Date: 10.12.2016</p>
          </div>
        </div>
        <div class="message_view_friend">
          <div class="message_view_friend__content">
            <p>
              Lorem ipsum dolor sit amet consectetur adipisicing elit.
              Reiciendis, magni dignissimos? Id delectus ab labore tenetur
              placeat, beatae cupiditate sit quaerat iste harum cum nesciunt
              quos magnam laudantium quas perspiciatis eligendi numquam
              provident? Ipsam, vero.
            </p>
            <p class="message_date">Date: 10.12.2016</p>
          </div>
        </div>
        <div class="message_view_me">
          <div class="message_view_me__content">
            <p>
              Lorem ipsum dolor sit, amet consectetur adipisicing elit. Odio
              vitae qui fugit debitis, repellat sit placeat?
            </p>
            <p class="message_date">Date: 10.12.2016</p>
          </div>
        </div>
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
export default {
  data() {
    return {
      test: 'test'
    };
  },
  methods: {
    scrollBottom: function() {
      let messageView = document.getElementById("message_view");
      messageView.scrollTop = messageView.scrollHeight;
    }
  },
  mounted: function() {
    this.scrollBottom()
  }
};
</script>
