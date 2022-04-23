// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Directory/router';
import hideElemDir from '../../assets/js/hideElem';
import '../../assets/js/delayedLoading';
import '../../assets/js/stickyHeader';

window.Vue = require('vue');

Vue.use(VueRouter);
Vue.directive('hide-elem', hideElemDir);

function directoryFixed() {
  if (document.querySelector('.directory_fixed')) {
    const directoryFixedBlock = document.querySelector('.directory_fixed');
    let top = Number();

    window.addEventListener('scroll', () => {
      if (window.scrollY <= 30) {
        top = directoryFixedBlock.getBoundingClientRect().top;
      } else {
        top = directoryFixedBlock.getBoundingClientRect().top - 50;
      }
      setTimeout(() => {
        document.querySelector(
          '.directory_fixed',
        ).style.transform = `translateY(${top + window.scrollY}px)`;
      }, 300);
    });
  }
}
directoryFixed();

(() => new Vue({
  el: '#wrapper',
  delimiters: ['{%', '%}'],
  data: {
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    showHideDirectory: true,
    showHideComponent: true,
  },
  methods: {
    signOut() {
      fetch('/api/signout_account')
        .then((response) => {
          if (response.ok) {
            document.cookie = 'token=; path=/; max-age=-1;';
            document.cookie = 'userId=; path=/; max-age=-1;';
            window.location.href = '/';
          }
        })
        .catch((err) => console.error(err));
    },
    hideStickyHeader() {
      document.querySelector(
        '.header_sticky',
      ).style.transform = 'translateX(-200px)';
      document
        .querySelector('.main')
        .classList.remove('main_squeeze_before_add');
      document
        .querySelector('.main')
        .classList.add('main_squeeze_before_remove');
      document
        .querySelector('.main__body_content')
        .classList.remove('main_squeeze');
      document
        .querySelector('.main__body_content')
        .classList.add('main_unclench');
      setTimeout(() => {
        document
          .querySelector('.main')
          .classList.remove('main_squeeze_before_remove');
      }, 490);
    },
    deletePreloader() {
      setTimeout(() => {
        this.preload = false;
      }, 1000);
    },
    showComponent() {
      if (document.documentElement.clientWidth <= 600) {
        this.showHideDirectory = false;
        this.showHideComponent = true;
      } else {
        this.showHideDirectory = true;
        this.showHideComponent = true;
      }
    },
    hideComponent(data = Object) {
      this.showHideDirectory = data.hide;
      this.showHideComponent = false;
    },
    preShowComponent() {
      if (document.documentElement.clientWidth <= 600) {
        this.showHideComponent = false;
      } else if (
        document.documentElement.clientWidth > 600
      ) {
        this.showHideComponent = true;
      }
    },
  },
  computed: {
    dropHeader() {
      return !(document.documentElement.clientWidth <= 960);
    },
    showStickyHeader() {
      return document.documentElement.clientWidth <= 960;
    },
    setDeletePreloader() {
      return {
        deletePreload: true,
      };
    },
    rotateArrowCommunity() {
      return {
        arrow_list_open: this.community,
      };
    },
    rotateArrowCommerce() {
      return {
        arrow_list_open: this.commerce,
      };
    },
    rotateArrowSettings() {
      return {
        arrow_list_open: this.settings,
      };
    },
  },
  created() {
    this.deletePreloader();
    this.preShowComponent();
  },
  router,
}))();
