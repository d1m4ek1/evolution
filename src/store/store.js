import { createStore } from "vuex";

import GetCookie from "../assets/javascript/getCookie.js";

export const store = createStore({
  state() {
    return {
      chatData: undefined,
      userId: GetCookie("userId"),
      audioSettings: {
        treckId: undefined,
        loadedAudio: null,
        name: "",
        producer: "",
        feat: "",
        tags: [],
        typeTreck: "",
        volume: 0.05,
        played: false,
        repeat: false,
        price: undefined,
        album: {
          name: "",
          id: undefined,
          cover: "",
          typeTrecks: "",
          date: "",
          played: false,
          trecks: [],
        },
      },
    };
  },
  getters: {
    getChatData: (state) => state.chatData,
    getUserId: (state) => state.userId,
    getAudioSettings: (state) => state.audioSettings,
  },
});
