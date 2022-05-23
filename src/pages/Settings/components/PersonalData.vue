<template>
  <section class="container settings_content">
    <!-- BACKUPS -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Ключи восстановления аккаунта</h2>
        <p>
          Поля ключей восстановления аккаунта нужно заполнить все, либо оставить
          пустыми <span class="note"></span>
        </p>
        <p v-if="userData.backupKeys.available">
          У вас есть ключ восстановления, но вы может изменить его в любое время
        </p>
        <p v-else>
          Ключа восстановления нет, создайте ключи, затем запишите их
          <span class="note"></span>
        </p>
      </div>
      <div class="change_content">
        <template v-for="(item, idx) in userData.backupKeys.keys">
          <h3 :key="idx">Ключи восстановления #{{ idx + 1 }}</h3>
          <div class="inlineblock">
            <input
              v-model="item.new"
              @input="fieldsNotEmpty(idx)"
              type="text"
              maxlength="100"
              :class="{ error: item.error }"
              :placeholder="item.placeholder"
            />
          </div>
        </template>
      </div>
    </div>
    <!-- CHANGE PASSWORD -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Смена пароля</h2>
      </div>
      <div class="change_content">
        <div class="change_content__title">
          <h3>Новый пароль</h3>
          <p>
            Безопаснее всего создать пароль длиною более 16 символов, что и
            делает генератор <span class="note"></span>
          </p>
        </div>
        <div class="inlineblock">
          <input type="password" placeholder="Новый пароль..." />
        </div>
        <div class="inlineblock">
          <button class="btn">Сгенерировать пароль</button>
        </div>
      </div>
    </div>
    <!-- CHANGE EMAIL -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Смена электронной почты</h2>
      </div>
      <div class="change_content">
        <div class="change_content__title">
          <h3>Новая электронная почта</h3>
          <p>
            Электронная почта нужна для получения рекламных новостей (не
            обязательно), аутентификации, а также для подтверждения изменения
            каких-либо персональных данных <span class="note"></span>
          </p>
        </div>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>{{ userData.email.old }}</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              type="email"
              placeholder="example@example.com"
              maxlength="80"
            />
          </div>
        </div>
      </div>
    </div>
    <!-- SAVE BUTTON OR DEFAULT SETTINGS -->
    <div class="change_section btns">
      <div class="change_content">
        <div class="change_content__title">
          <h3>Подтверждение</h3>
          <p>
            Только с подтверждением пароля, можно изменить персональные
            данные.<span class="note"></span><br />
            <a href="#">Забыли пароль?</a>
          </p>
        </div>
        <div class="inlineblock">
          <input
            name="now_password"
            type="password"
            placeholder="Текущий пароль..."
          />
        </div>
      </div>
      <div class="change_content__title">
        <p>
          Сохраняя данные вы согласны с
          <a href="/directory/terms-of-use"> пользовательским соглашением</a> и
          <a href="/directory/privacy-policy">политикой конфиденциальности.</a>
          <span class="note"></span>
        </p>
      </div>
      <button @click="completePassword()" class="btn">Подвердить пароль</button>
      <button
        v-if="completePassword.completed"
        @click="saveSettings()"
        class="btn"
      >
        Сохранить
      </button>
      <button @click="resetSettings()" class="btn btn-red">
        Сбросить все изменения
      </button>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      completePassword: {
        password: '',
        completed: false,
      },
      userData: {
        backupKeys: {
          available: false,
          keys: [
            {
              new: '',
              placeholder: 'Первый ключ',
              error: false,
            },
            {
              new: '',
              placeholder: 'Второй ключ',
              error: false,
            },
            {
              new: '',
              placeholder: 'Третий ключ',
              error: false,
            },
            {
              new: '',
              placeholder: 'Четвертый ключ',
              error: false,
            },
          ],
        },
        email: {
          old: '',
          new: '',
        },
      },
    };
  },
  methods: {
    resetSettings() {
      return window.location.reload();
    },
    getUserData() {
      fetch('/api/get_settings/personal_data').then((response) => {
        response.json().then((data) => {
          this.userData.backupKeys.available = data.bcpk;
          this.userData.email.old = data.eml;
        });
      });
    },
    fieldsNotEmpty(id) {
      this.userData.backupKeys.keys.forEach((item, index) => {
        if (index !== id && item.new === '') {
          this.userData.backupKeys.keys[index].error = true;
        } else {
          this.userData.backupKeys.keys[index].error = false;
        }
      });
      let countError = 0;
      if (this.userData.backupKeys.keys[id].new === '') {
        this.userData.backupKeys.keys.forEach((item) => {
          if (item.new === '' && countError !== 4) countError += 1;
        });
        if (countError === 4) {
          countError = 0;
          this.userData.backupKeys.keys.forEach((item, index) => {
            this.userData.backupKeys.keys[index].error = false;
          });
        }
      }
    },
    saveSettings() {
      const defUrl = '/api/save_settings/profile?';
      const urlParam = [];
      const fullUrl = '';
    },
  },
  created() {
    this.getUserData();
  },
};
</script>
