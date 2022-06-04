<template>
  <section class="container settings_content">
    <h1 class="connect_bd">Связано с Базой данных</h1>
    <!-- NICKNAME -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить мой псевдоним</h2>
      </div>
      <div class="change_content">
        <h3>Псевдоним сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>{{ userData.name.old }}</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input v-model="userData.name.new" type="text" maxlength="100" />
          </div>
        </div>
      </div>
    </div>
    <!-- LANGUAGE -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить язык сайта</h2>
      </div>
      <div class="change_content">
        <h3>Язык сайта сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>{{ userData.language.old }}</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <select v-model="userData.language.new">
              <option value="ru">Русский</option>
              <option value="en">English</option>
            </select>
          </div>
        </div>
      </div>
    </div>
    <!-- THEME PAGE -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить тему сайта</h2>
      </div>
      <div class="change_content">
        <h3>Тема сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>{{ userData.themePage.old }}</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <select v-model="userData.themePage.new">
              <option value="white">Светлая тема</option>
              <option value="dark">Темная тема</option>
            </select>
          </div>
        </div>
      </div>
    </div>
    <!-- AVA -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить аватарку</h2>
      </div>
      <div class="change_content">
        <h3>Аватарка сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left img">
            <img class="lazyload" :data-src="userData.logo.old" />
          </div>
          <div class="upload_content">
            <input
              @change="traceFile($event, 'logo')"
              id="upload_logo_file"
              type="file"
              name="upload_file"
            />
            <label for="upload_logo_file" id="label_input_drop">
              <template v-if="userData.logo.new === ''">
                <h2>Загрузить аватарку</h2>
                <h3>Расширения .jpg, .png, .jpeg, .gif</h3>
              </template>
              <template v-else>
                <div class="img">
                  <img class="lazyload" :data-src="userData.logo.preview" />
                </div>
              </template>
            </label>
          </div>
        </div>
        <button
          v-if="userData.logo.new !== ''"
          @click="(userData.logo.preview = ''), (userData.logo.new = '')"
          class="btn btn-red"
        >
          Сбросить аватарку
        </button>
        <p class="error" v-if="userData.logo.error !== ''">{{ userData.logo.error }}</p>
      </div>
    </div>
    <!-- BANNER -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить баннер</h2>
      </div>
      <div class="change_content">
        <h3>Баннер сейчас</h3>
        <div class="inlineblock">
          <div class="inlineblock_left banner">
            <img class="lazyload" :data-src="userData.banner.old" alt="Нет баннера" />
          </div>
          <div class="upload_content">
            <input
              @change="traceFile($event, 'banner')"
              id="upload_banner_file"
              type="file"
              name="upload_file"
            />
            <label for="upload_banner_file" id="label_input_drop">
              <template v-if="userData.banner.new === ''">
                <h2>Загрузить баннер</h2>
                <h3>Расширения .jpg, .png, .jpeg, .gif</h3>
              </template>
              <template v-else>
                <div class="img">
                  <img class="lazyload" :data-src="userData.banner.preview" alt="Нет баннера" />
                </div>
              </template>
            </label>
          </div>
        </div>
        <button
          v-if="userData.banner.new !== ''"
          @click="(userData.banner.new = ''), (userData.banner.preview = '')"
          class="btn btn-red"
        >
          Сбросить баннер
        </button>
        <p class="error" v-if="userData.banner.error !== ''">{{ userData.banner.error }}</p>
      </div>
    </div>
    <!-- ABOUT ME -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить информацию "Обо мне"</h2>
      </div>
      <div class="change_content">
        <h3>Название</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.aboutMe.title.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              v-model="userData.aboutMe.title.new"
              type="text"
              maxlength="300"
            />
          </div>
        </div>
        <h3>Содержимое</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.aboutMe.content.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <div class="write_text aboutme_content" contenteditable=""></div>
          </div>
        </div>
      </div>
    </div>
    <!-- CONNECTION -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Изменить контактную информацию</h2>
      </div>
      <div class="change_content">
        <h3>Telegram</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.connection.telegram.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              v-model="userData.connection.telegram.new"
              type="text"
              maxlength="250"
            />
          </div>
        </div>
        <h3>Instagram</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.connection.instagram.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              v-model="userData.connection.instagram.new"
              type="text"
              maxlength="250"
            />
          </div>
        </div>
        <h3>Facebook</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.connection.facebook.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              v-model="userData.connection.facebook.new"
              type="text"
              maxlength="250"
            />
          </div>
        </div>
        <h3>VK</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.connection.vk.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              v-model="userData.connection.vk.new"
              type="text"
              maxlength="250"
            />
          </div>
        </div>
        <h3>TikTok</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              {{ userData.connection.tiktok.old }}
            </div>
          </div>
          <div class="inlineblock_right">
            <input
              v-model="userData.connection.tiktok.new"
              type="text"
              maxlength="250"
            />
          </div>
        </div>
      </div>
    </div>
    <!-- SAVE BUTTON OR DEFAULT SETTINGS -->
    <div class="change_section btns">
      <button @click="saveSettings()" class="btn">Сохранить</button>
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
      userData: {
        logo: {
          old: '',
          preview: '',
          file: null,
          error: '',
          new: '',
        },
        banner: {
          old: '',
          preview: '',
          file: null,
          error: '',
          new: '',
        },
        name: {
          old: '',
          new: '',
        },
        language: {
          old: '',
          new: '',
        },
        aboutMe: {
          title: {
            old: '',
            new: '',
          },
          content: {
            old: '',
            new: '',
          },
        },
        themePage: {
          old: '',
          new: '',
        },
        connection: {
          telegram: {
            old: '',
            new: '',
          },
          instagram: {
            old: '',
            new: '',
          },
          facebook: {
            old: '',
            new: '',
          },
          vk: {
            old: '',
            new: '',
          },
          tiktok: {
            old: '',
            new: '',
          },
        },
      },
    };
  },
  methods: {
    traceFile(e = Object, key = String) {
      const extens = ['.png', '.jpg', '.jpeg', '.gif'];
      let extenErrorCount = 0;

      extens.forEach((exten = String) => {
        if (e.target.files[0].name.indexOf(exten) !== -1) {
          const fileName = e.target.files[0].name.split('.');

          if (
            fileName[fileName.length - 1] === 'png'
              || fileName[fileName.length - 1] === 'jpg'
              || fileName[fileName.length - 1] === 'jpeg'
              || fileName[fileName.length - 1] === 'gif'
          ) {
            this.userData[key].new = e.target.files[0].name;
            this.userData[key].error = '';
            this.userData[key].file = e.target.files;

            const reader = new FileReader();
            this.addListeners(reader, key);
            reader.readAsDataURL(e.target.files[0]);
          } else {
            this.userData[key].error = 'Файл не является изображением';
          }
        } else {
          extenErrorCount += 1;
        }
      });

      if (extenErrorCount === 4) {
        this.userData[key].error = 'Файл не является изображением';
      }
    },
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
    validateLanguage(lang = String) {
      switch (lang) {
        case 'ru':
          this.userData.language.old = 'Русский';
          break;
        case 'en':
          this.userData.language.old = 'English';
          break;

        default:
          this.userData.language.old = 'Не задано';
          break;
      }
    },
    getUserData() {
      fetch('/api/get_settings/profile').then((response = Object) => {
        response.json().then((data = Object) => {
          this.validateLogoBanner(data.logo, data.banner);
          this.validateLanguage(data.language);

          this.userData.name.old = data.name;
          this.userData.themePage.old = data.themePage;
          this.userData.aboutMe.title.old = data.aboutMeTitle;
          this.userData.aboutMe.content.old = data.aboutMeContent;

          this.userData.connection.telegram.old = data.telegram;
          this.userData.connection.instagram.old = data.instagram;
          this.userData.connection.facebook.old = data.facebook;
          this.userData.connection.vk.old = data.vk;
          this.userData.connection.tiktok.old = data.tiktok;
        });
      });
    },
    saveSettings() {
      const defUrl = '/api/save_settings/profile?';
      const urlParam = [];
      const formData = new FormData();
      let fullUrl = '';

      this.userData.aboutMe.title.new = document.querySelector('.aboutme_content').innerHTML;

      Object.keys(this.userData).forEach((keyMain = String) => {
        if (keyMain === 'aboutMe') {
          Object.keys(this.userData.aboutMe).forEach((key = String) => {
            if (this.userData.aboutMe[key].new !== '') {
              if (
                this.userData.aboutMe[key].old
                !== this.userData.aboutMe[key].new && this.userData.aboutMe[key].new !== ''
              ) {
                urlParam.push(
                  `aboutMe_${key}=${encodeURIComponent(
                    this.userData.aboutMe[key].new,
                  )}`,
                );
              }
            }
          });
        } else if (keyMain === 'connection') {
          Object.keys(this.userData.connection).forEach((key = String) => {
            if (this.userData.connection[key].new !== '') {
              if (
                this.userData.connection[key].old
                !== this.userData.connection[key].new && this.userData.connection[key].new !== ''
              ) {
                urlParam.push(
                  `${key}=${encodeURIComponent(
                    this.userData.connection[key].new,
                  )}`,
                );
              }
            }
          });
        } else if (keyMain === 'logo' || keyMain === 'banner') {
          if (
            this.userData[keyMain].old
                !== this.userData[keyMain].new && this.userData[keyMain].new !== ''
          ) {
            if (this.userData[keyMain].file !== null) {
              formData.append(keyMain, this.userData[keyMain].file[0]);
              urlParam.push(
                `${keyMain}=${encodeURIComponent(
                  this.userData[keyMain].new,
                )}`,
              );
            }
          }
        } else if (this.userData[keyMain].new !== '') {
          if (this.userData[keyMain].old !== this.userData[keyMain].new) {
            urlParam.push(
              `${keyMain}=${encodeURIComponent(this.userData[keyMain].new)}`,
            );
          }
        }
      });

      fullUrl += defUrl + urlParam.join('&');

      if (fullUrl !== defUrl) {
        fetch(fullUrl, {
          method: 'POST',
          body: formData,
        }).catch((error) => console.error(error));
      }
    },
    resetSettings() {
      Object.keys(this.userData).forEach((keyMain = String) => {
        if (keyMain === 'aboutMe') {
          Object.keys(this.userData.aboutMe).forEach((key = String) => {
            this.userData.aboutMe[key].new = '';
          });
        } else if (keyMain === 'connection') {
          Object.keys(this.userData.connection).forEach((key = String) => {
            this.userData.connection[key].new = '';
          });
        } else {
          this.userData[keyMain].new = '';
        }
      });
    },
    addListeners(reader = Object, key = String) {
      reader.addEventListener('loadstart', () => {
        this.userData[key].preview = reader.result;
      });
      reader.addEventListener('load', () => {
        this.userData[key].preview = reader.result;
      });
      reader.addEventListener('loadend', () => {
        this.userData[key].preview = reader.result;
      });
      reader.addEventListener('progress', () => {
        this.userData[key].preview = reader.result;
      });
      reader.addEventListener('error', () => {
        this.userData[key].preview = reader.result;
      });
      reader.addEventListener('abort', () => {
        this.userData[key].preview = reader.result;
      });
    },
  },
  created() {
    this.getUserData();
  },
};
</script>
