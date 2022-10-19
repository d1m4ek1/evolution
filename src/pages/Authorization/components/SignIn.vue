<template>
  <div class="card-sign-content">
    <input v-model="dataSignIn.login" type="text" name="login" placeholder="Логин..." />
    <input v-model="dataSignIn.password" type="password" name="password" placeholder="Пароль..." />
    <div class="link-forgot">
      <a href="#">Забыли логин или пароль?</a>
    </div>
    <button @click="sendData()" class="btn">Вход</button>
  </div>
</template>

<script>
import MD5 from 'crypto-js/md5';

export default {
  props: {
    activetitle: String,
  },
  data() {
    return {
      dataSignIn: {
        login: String(),
        password: String(),
      },
      error: String(),
    };
  },
  methods: {
    sendData() {
      fetch(`/api/signin_account?signin=true&login=${this.dataSignIn.login}&password=${MD5(this.dataSignIn.password)}`).then((response) => {
        response.json().then((data) => {
          if (data.error === undefined && data.num === undefined) {
            window.location.href = `/${data.user_id}`;
          } else {
            this.error = data.error;
          }
        });
      });
    },
  },
};
</script>
