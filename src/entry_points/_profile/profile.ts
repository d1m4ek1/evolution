// JS
import Vue from 'vue';
import VueRouter from 'vue-router';
import router from '../../router/Profile/profile.router';
import hideElemDir from '../../assets/typescript/hideElem';
import 'lazysizes';
import '../../assets/typescript/stickyHeader';
import MODULE_CHECK_AUTHORIZE_USER from '../../assets/typescript/modules/CheckAuthorize.module';
import MODULE_STICKY_HEADER from "../../assets/typescript/modules/StickyHeader.module";
import MODULE_SIGN_OUT from "@/assets/typescript/modules/SignOut.module";
import StickyHeader from "@/assets/typescript/stickyHeader";
StickyHeader()

window.Vue = require('vue');

Vue.directive('hide-elem', hideElemDir);
Vue.use(VueRouter);

(() => new Vue({
  el: '#wrapper',
  delimiters: ['{%', '%}'],
  data: {
    preload: true,
    commerce: false,
    community: false,
    settings: false,
    aboutme: {
      title: '',
      content: '',
    },
    isSubscriber: false,
    isCountSubscribers: 0
  },
  methods: {
    signOut() {
      MODULE_SIGN_OUT()
    },
    hideStickyHeader() {
      MODULE_STICKY_HEADER()
    },
    deletePreloader() {
      setTimeout(() => {
        this.preload = false;
      }, 1000);
    },
    getDataProfile() {
      const userId = window.location.pathname.split('/');
      fetch(`/api/get_data_profile?get_data=all&user_id=${userId[1]}`).then(
        (response) => {
          response.json().then((data) => {
            this.aboutme.title = data.aboutmeTitle;
            this.aboutme.content = data.aboutmeContent;
          });
        },
      );
    },
    appendSubscriber(id) {
      fetch(`/api/append_subscriber?append_id=${id}`, {
        method: "POST"
      }).then((response: Response) => {
        if (response.ok) {
          this.isSubscriber = true
          this.getCountSubscriber();
        }
      }).catch((error) => console.error(error))
    },
    checkSubcriber() {
      let userId: string = window.location.pathname.split('/')[1]
      fetch(`/api/check_subscriber?check_id=${userId}`, {
        method: "GET"
      }).then((response: Response) => {
        if (!response.ok) {
          console.error(response.statusText)
          return
        }

        response.json().then((data) => {
          this.isSubscriber = data.isSubscriber
          this.getCountSubscriber();
        })
      }).catch(error => console.error(error))
    },
    deleteSubscriber(id) {
      fetch(`/api/delete_subscriber?delete_id=${id}`, {
        method: "DELETE"
      }).then((response: Response) => {
        if (!response.ok) {
          this.isSubscriber = true
          console.error(response.statusText)
        } else {
          this.isSubscriber = false
          this.getCountSubscriber();
        }
      }).catch(error => console.error(error))
    },
    getCountSubscriber() {
      let userId: string = window.location.pathname.split('/')[1]
      fetch(`/api/count_subscriber?check_id=${userId}`,
          {
            method: "GET"
          }).then((response: Response) => {
            if (!response.ok) {
              console.error(response.statusText)
              return
            }

            response.json().then(data => {
              this.isCountSubscribers = data.isCountSubscriber
            })
      }).catch(error => console.error(error))
    }
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
    this.getDataProfile();
    this.checkSubcriber();
    this.getCountSubscriber();
    MODULE_CHECK_AUTHORIZE_USER();
  },
  router,
}))();
