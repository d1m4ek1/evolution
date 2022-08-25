import { createStore } from "vuex";

export const store = createStore({
  state() {
    return {
      chatData: undefined,
    };
  },
  getters: {
    getChatData: (state) => state.chatData,
  },
});
