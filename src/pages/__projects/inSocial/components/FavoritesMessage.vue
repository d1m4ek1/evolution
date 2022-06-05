<template>
  <div class="messages show_content">
    <section class="title">
      <h1>Избранные сообщения</h1>
    </section>
    <section class="meesages_content__folders">
      <div class="messages_folders">
        <div class="messages_folder">
          <router-link tag="button" to="/insocial" class="btn"
            >Все сообщения</router-link
          >
        </div>
        <div class="messages_folder">
          <router-link tag="button" to="/insocial/favorites" class="btn"
            >Избранные сообщения</router-link
          >
        </div>
      </div>
    </section>
    <section class="mesages_content">
      <div class="message_items">
        <div class="message_item">
          <router-link to="/insocial/control"></router-link>
          <div class="message_item__info">
            <div class="message_item__logo online">
              <img class="lazyload" data-src="/profile/logo/notLogo/not_logo.png" />
            </div>
            <div class="message_item__name">
              <h2>Name</h2>
            </div>
            <div class="message_item__control">
              <button class="btn">управлять</button>
            </div>
          </div>
          <div class="message_item__banner">
            <img class="lazyload" data-src="/profile/banner/notBanner/not_banner.png" />
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
