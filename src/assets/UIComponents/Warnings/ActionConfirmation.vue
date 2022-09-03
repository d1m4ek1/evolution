<template>
  <div class="block_warning" v-if="showPopup">
    <div class="card_warning">
      <h2>{{ title }}</h2>
      <p>{{ textContent }}</p>
      <div class="control">
        <button class="btn" @click="returnResult(true)">Подтвердить</button>
        <button class="btn" @click="returnResult(false)">Отменить</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.block_warning {
  position: fixed;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.2);
  display: flex;
  justify-content: center;
  align-items: center;
  top: 0;
}
.card_warning {
  width: 500px;
  height: 300px;
  background-color: white;
  border-radius: 10px;
  border: 1px solid;
}
.card_warning h2 {
  padding-bottom: 10px;
}
.card_warning p {
  padding-bottom: 20px;
}
.control {
  display: flex;
  justify-content: space-between;
}
</style>

<script>
export default {
  props: {
    title: String,
    textContent: String,
    showPopup: Boolean,
  },
  methods: {
    returnResult(response) {
      this.$emit("hide-popup-warning", {
        hide: true,
        response: response,
      });
    },
  },
};
</script>
