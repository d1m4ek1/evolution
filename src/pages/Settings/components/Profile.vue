<template>
  <section class="container settings_content">
    <div class="change_section">
      <h2>Изменить мой псевдоним</h2>
      <div class="change_content">
        <h3>Псевдоним сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>{{ userData.name.old }}</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input type="text">
          </div>
        </div>
      </div>
    </div>
    <div class="change_section">
      <h2>Изменить язык</h2>
      <div class="change_content">
        <h3>Язык сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>Русский</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <select>
              <option value="ru">Русский</option>
              <option value="en">English</option>
            </select>
          </div>
        </div>
      </div>
    </div>
    <div class="change_section">
      <h2>Изменить аватарку</h2>
      <div class="change_content">
        <h3>Аватарка сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left img">
            <img :src="userData.logo.old" alt="Нет аватарки" />
          </div>
          <div class="upload_content">
            <input id="upload_logo_file" type="file" name="upload_file" />
            <label for="upload_logo_file" id="label_input_drop">
              <h2>Загрузить аватарку</h2>
              <h3>Расширения .jpg, .png, .jpeg, .gif</h3>
            </label>
          </div>
        </div>
      </div>
    </div>
    <div class="change_section">
      <h2>Изменить баннер</h2>
      <div class="change_content">
        <h3>Баннер сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left banner">
            <img
              :src="userData.banner.old"
              alt="Нет баннера"
            />
          </div>
          <div class="upload_content">
            <input id="upload_banner_file" type="file" name="upload_file" />
            <label for="upload_banner_file" id="label_input_drop">
              <h2>Загрузить баннер</h2>
              <h3>Расширения .jpg, .png, .jpeg, .gif</h3>
            </label>
          </div>
        </div>
      </div>
    </div>
    <div class="change_section">
      <h2>Изменить информацию "Обо мне"</h2>
      <div class="change_content">
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.aboutMe.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <div class="write_text notve_text" contenteditable=""></div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      userData: {
        logo: {
          old: '',
          new: '',
        },
        banner: {
          old: '',
          new: '',
        },
        name: {
          old: '',
          new: '',
        },
        aboutMe: {
          old: '',
          new: '',
        },
      },
    };
  },
  methods: {
    validateLogoBanner(logo = String, banner = String) {
      if (logo === 'not_logo.png') {
        this.userData.logo.old = '/profile/logo/notLogo/not_logo.png';
      } else {
        this.userData.logo.old = `/profile/logo/saved/${logo}`;
      }
      if (banner === 'not_banner.png') {
        this.userData.banner.old = '/profile/banner/notBanner/not_banner.png';
      } else {
        this.userData.banner.old = `/profile/banner/saved/${banner}`;
      }
    },
    getUserData() {
      fetch('/api/get_settings/profile').then((response = Object) => {
        response.json().then((data = {
          logo: String,
          banner: String,
          name: String,
          aboutMe: String,
        }) => {
          this.validateLogoBanner(data.logo, data.banner);
          this.userData.name.old = data.name;
          this.userData.aboutMe.old = data.aboutMe;
        });
      });
    },
  },
  created() {
    this.getUserData();
  },
};
</script>
