<template>
  <div class="card-sign-content">
    <input
      v-model="dataSignIn.login"
      type="text"
      name="login"
      placeholder="Логин..."
    />
    <input
      v-model="dataSignIn.password"
      type="password"
      name="password"
      placeholder="Пароль..."
    />
    <div class="link-forgot">
      <a href="#">Забыли логин или пароль?</a>
    </div>
    <button @click="sendData()" class="btn">Вход</button>
  </div>
</template>

<script>
import MD5 from "crypto-js/md5";

function rndsh(sumString = Number()) {
  const symbolArr =
    "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM";
  let rtsdnr = "";
  for (let i = 0; i < sumString; i += 1) {
    const index = Math.floor(Math.random() * symbolArr.length);
    rtsdnr += symbolArr[index];
  }
  return rtsdnr;
}
export default {
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
      const cookie = `token=${MD5(rndsh(64)) + rndsh(8)}`;
      fetch(
        `/api/signin_account?signin=true&login=${
          this.dataSignIn.login
        }&password=${MD5(this.dataSignIn.password)}&${cookie}`
      ).then((response) => {
        response.json().then((data) => {
          if (data.error === undefined && data.num === undefined) {
            if (data.olt !== "") {
              document.cookie = `token=${data.olt}; path=/;`;
            } else {
              document.cookie = `${cookie}; path=/;`;
            }
            document.cookie = `userId=${data.user_id}; path=/;`;
            window.location.href = `/${data.user_id}`;
          } else {
            this.error = data.error;
          }
        });
      });
    },
  },
  created() {
    this.$emit("changed-background", {
      backgroundNumber: 1,
    });
  },
};
</script>
