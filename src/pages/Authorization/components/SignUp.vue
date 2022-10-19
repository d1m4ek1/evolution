<template>
  <div class="card-sign-content">
    <div class="description">
      <div class="description-content" v-for="(item, idx) in inputs.logins" :key="idx + '2'">
        <div class="necessarily">
          <input
            v-model="item.value"
            :type="item.type"
            :name="item.id"
            :placeholder="item.placeholder + '...'"
            :maxlength="item.maxlength"
          />
        </div>
      </div>
      <div class="description-content" v-for="(item, idx) in inputs.passwords" :key="idx + '3'">
        <div class="necessarily">
          <input
            v-model="item.value"
            @input="validPassword($event.target.value, item), samePassword($event.target.value, idx)"
            :class="{ 'not-valid': item.valid }"
            :type="item.type"
            :name="item.id"
            :placeholder="item.placeholder + '...'"
          />
        </div>
      </div>
    </div>
    <button @click="createAccount()" class="btn">Зарегистрироваться</button>
    <p v-if="validation.symbols.valid">{{ validation.symbols.value }}</p>
    <p v-if="validation.same.valid">{{ validation.same.value }}</p>
    <p v-if="validation.allInputs.valid">
      {{ validation.allInputs.value }}
    </p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      inputs: {
        logins: [
          {
            id: 'nickname',
            placeholder: 'Псевдоним',
            type: 'text',
            maxlength: 110,
            value: String(),
          },
          {
            id: 'email',
            placeholder: 'Электронная почта',
            type: 'email',
            value: String(),
          },
          {
            id: 'login',
            placeholder: 'Логин',
            type: 'text',
            value: String(),
          },
        ],
        passwords: [
          {
            id: 'password',
            placeholder: 'Пароль',
            type: 'password',
            value: String(),
            valid: false,
            conf: true,
          },
          {
            id: 'password',
            placeholder: 'Подтвердить пароль',
            type: 'password',
            value: String(),
            valid: false,
          },
        ],
      },
      validation: {
        symbols: {
          valid: false,
          value: 'Разрешены только буквы и цифры!',
        },
        same: {
          valid: false,
          value: 'Пароли не совпадают!',
        },
        backupKey: {
          valid: false,
          value: 'Разрешены только буквы и цифры!',
        },
        allInputs: {
          value: 'Заполните все поля!',
          valid: false,
        },
      },
      urlCreateAccount: [],
    };
  },
  methods: {
    validPassword(e, obj = { valid: Boolean() }) {
      const validPassword = obj;

      if (e.match(/([!@#$%^&*(){}[]:;"'<\.>,\?\/\|~`№\?-_=\+])/g)) {
        validPassword.valid = true;
        this.validation.symbols.valid = true;
      } else {
        this.validation.symbols.valid = false;
        validPassword.valid = false;
      }
    },
    samePassword(e, id) {
      const newId = id === 0 ? 1 : 0;

      if (e !== this.inputs.passwords[newId].value) {
        this.inputs.passwords[newId].valid = true;
        this.validation.same.valid = true;
      } else {
        this.inputs.passwords[newId].valid = false;
        this.validation.same.valid = false;
      }
    },
    createUrl() {
      this.urlCreateAccount = [];

      Object.keys(this.inputs).forEach((key) => {
        for (let i = 0; i < this.inputs[key].length; i += 1) {
          const el = this.inputs[key][i];
          if (el.value !== '' && el.valid !== true) {
            switch (key) {
              case 'passwords':
                if (!el.conf) {
                  this.urlCreateAccount.push(`${el.id}=${MD5(el.value)}`);
                }
                break;
              default:
                if (!el.conf) {
                  this.urlCreateAccount.push(`${el.id}=${el.value}`);
                }
                break;
            }
          } else {
            this.urlCreateAccount = [];
            this.validation.allInputs.valid = true;
            return;
          }
        }
      });
    },
    createAccount() {
      this.createAccount();

      if (this.urlCreateAccount.length !== 0) {
        fetch(`/api/create_account?${this.urlCreateAccount.join('&')}`, {
          method: 'POST',
        })
          .then((response) => {
            response.json().then((data) => {
              if (data.aut) {
                fetch(
                  `/api/signin_account?signin=true&login=${this.inputs.logins[2].value}&password=${MD5(this.inputs.passwords[1].value)}`
                ).then((preresponse) => {
                  preresponse.json().then((predata) => {
                    if (predata.user_id !== undefined) {
                      document.cookie = `userId=${predata.user_id}; path=/;`;
                      window.location.href = `/${predata.user_id}`;
                    }
                  });
                });
              }
            });
          })
          .catch((err) => console.error(err));
      }
    },
  },
};
</script>
