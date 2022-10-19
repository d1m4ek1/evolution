import { createStore } from 'vuex';

import GetCookie from '../assets/javascript/getCookie.js';

export const store = createStore({
  state() {
    return {
      messengerData: {
        chats: [],
        subscribers: [],
        subscriptions: [],
      },
      userId: GetCookie('userId'),
      audioSettings: {
        treckId: undefined,
        loadedAudio: null,
        name: '',
        producer: '',
        feat: '',
        tags: [],
        volume: 0.2,
        played: false,
        repeat: false,
        price: undefined,
        album: {
          name: '',
          creator: '',
          albumId: undefined,
          covers: '',
          dateOfRelease: '',
          played: false,
          trecks: [],
        },
      },
      userAlbums: {
        albums: [],
        isInit: false,
      },
    };
  },
  getters: {
    getMessengerData: (state) => state.messengerData,
    getUserId: (state) => state.userId,
    getAudioSettings: (state) => state.audioSettings,
    getUserAlbums: (state) => state.userAlbums,
  },
});
