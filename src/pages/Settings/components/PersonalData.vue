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
          <input
            v-model="userData.password.new"
            type="password"
            placeholder="Новый пароль..."
          />
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
              v-model="userData.email.new"
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
      <template v-if="!completePassword.completed">
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
              v-model="completePassword.password"
              name="now_password"
              type="password"
              placeholder="Текущий пароль..."
            />
          </div>
        </div>
        <button @click="sendOnConfirmPassword()" class="btn">
          Подвердить пароль
        </button>
      </template>
      <template v-else>
        <div class="change_content__title">
          <p>
            Сохраняя данные вы согласны с
            <a href="/directory/terms-of-use"> пользовательским соглашением</a>
            и
            <a href="/directory/privacy-policy"
              >политикой конфиденциальности.</a
            >
            <span class="note"></span>
          </p>
        </div>
        <button @click="saveSettings()" class="btn">Сохранить</button>
      </template>
      <button @click="resetSettings()" class="btn btn-red">
        Сбросить все изменения
      </button>
    </div>
  </section>
</template>

<script>
import MD5 from 'crypto-js/md5';

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
              param: 'backupkey_one',
              error: false,
            },
            {
              new: '',
              placeholder: 'Второй ключ',
              param: 'backupkey_two',
              error: false,
            },
            {
              new: '',
              placeholder: 'Третий ключ',
              param: 'backupkey_three',
              error: false,
            },
            {
              new: '',
              placeholder: 'Четвертый ключ',
              param: 'backupkey_four',
              error: false,
            },
          ],
        },
        password: {
          param: 'password',
          new: '',
        },
        email: {
          old: '',
          param: 'email',
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
      const defUrl = '/api/save_settings/personal_data?';
      const urlParam = [];
      let fullUrl = '';
      fullUrl = defUrl;

      let voidCounter = 0;

      for (let i = 0; i < this.userData.backupKeys.keys.length; i += 1) {
        const temporaryKey = this.userData.backupKeys.keys[i];

        if (temporaryKey.new !== '' && voidCounter !== 4) {
          voidCounter += 1;
        }
      }

      if (voidCounter === 4) {
        voidCounter = 0;
        for (let j = 0; j < this.userData.backupKeys.keys.length; j += 1) {
          const key = this.userData.backupKeys.keys[j];
          urlParam.push(`${key.param}=${encodeURIComponent(MD5(key.new))}`);
        }
      }

      if (this.userData.password.new !== '') {
        urlParam.push(`${this.userData.password.param}=${MD5(this.userData.password.new)}`);
      }
      if (this.userData.email.old !== this.userData.email.new && this.userData.email.new !== '') {
        urlParam.push(`${this.userData.email.param}=${this.userData.email.new}`);
      }

      fullUrl += urlParam.join('&');

      if (fullUrl !== defUrl) {
        fetch(fullUrl, {
          method: 'POST',
        })
          .then((response) => {
            if (response.ok) {
              setTimeout(() => {
                window.location.reload();
              }, 3000);
            }
          })
          .catch((error) => console.error(error));
      }
    },
    sendOnConfirmPassword() {
      if (this.completePassword.password !== '') {
        fetch(`/api/confirm?conf_pass=${MD5(this.completePassword.password)}`, {
          method: 'GET',
        })
          .then((response) => {
            response.json().then((data) => {
              this.completePassword.completed = data.pass;
            });
          })
          .catch((error) => console.log(error));
      }
    },
  },
  created() {
    this.getUserData();
  },
};
</script>
