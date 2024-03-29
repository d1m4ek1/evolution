<template>
  <div class="global_player container" :class="showHidePlayer">
    <div v-if="$store.getters.getAudioSettings.loadedAudio" class="global_player__block">
      <div class="header_player">
        <div class="timer">
          <progress-bar-time
            :duration="$store.state.audioSettings.loadedAudio.duration"
            :currentTime="$store.state.audioSettings.loadedAudio.currentTime"
          ></progress-bar-time>
          <div class="timeUpdate">
            <p class="now_time">{{ time.nowTime }}</p>
            <p class="amount_time">{{ time.totalTime }}</p>
          </div>
        </div>
        <div class="audio_item">
          <button class="btn btn_open" @click="showPlayer()">{{ showHideText }}</button>
          <div class="treck_info">
            <div class="treck_cover">
              <img :src="'/user_files/beats/cover_album/' + $store.getters.getAudioSettings.album.covers" />
            </div>
            <div class="treck_names">
              <p>{{ $store.getters.getAudioSettings.name }}</p>
              <p>{{ $store.getters.getAudioSettings.album.creator }}</p>
            </div>
          </div>
          <div class="treck_price_buy">{{ $store.getters.getAudioSettings.price }}</div>
          <div class="treck_control">
            <button class="btn" v-if="!$store.getters.getAudioSettings.played" @click="playTreck()">Play</button>
            <button class="btn" v-else @click="pauseTreck()">Pause</button>
          </div>
        </div>
      </div>
      <div class="content_player">
        <div class="sidebar_left">
          <div class="album_item" v-if="$store.getters.getAudioSettings.album.id !== undefined">
            <div class="album_item__cover">
              <img :src="'/user_files/beats/cover_album/' + $store.getters.getAudioSettings.album.covers" />
              <button v-if="!$store.getters.getAudioSettings.album.played" @click="playTreck()">Play</button>
              <button v-else @click="pauseTreck()">Pause</button>
            </div>
            <h1>{{ $store.getters.getAudioSettings.album.name }}</h1>
            <h3>Date: {{ $store.getters.getAudioSettings.album.dateOfRelease }}</h3>
          </div>
        </div>
        <div class="content_player__main">
          <template v-if="$store.getters.getAudioSettings.album.trecks.length !== 0">
            <button class="btn" @click="showHideTrecks('.audio_items')">{{ clipPathTrecksBlockText }} trecks</button>
            <div class="audio_items" :class="clipPathTrecksBlock">
              <div
                v-for="(treckItem, idx) in $store.getters.getAudioSettings.album.trecks"
                class="audio_item"
                :key="'player_treck_item_' + idx"
              >
                <div class="treck_info">
                  <div class="treck_cover">
                    <img :src="'/user_files/beats/cover_album/' + $store.getters.getAudioSettings.album.covers" />
                  </div>
                  <div class="treck_names">
                    <p>{{ treckItem.name }}</p>
                    <p>{{ $store.getters.getAudioSettings.album.creator }}</p>
                  </div>
                </div>
                <div class="treck_price_buy">{{ treckItem.price }}</div>
                <div class="treck_control">
                  <button class="btn" v-if="!treckItem.played" @click="playTreck(idx)">Play</button>
                  <button class="btn" v-else @click="pauseTreck(idx)">Pause</button>
                </div>
              </div>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ProgressBarTime from './ProgressBarTime.vue';
export default {
  components: { ProgressBarTime },
  data() {
    return {
      hidenPlayer: true,
      time: {
        nowTime: '0:00',
        totalTime: '',
        nowTimeNumber: 0,
      },
      clipPath: {
        trecksBlock: false,
      },
    };
  },
  watch: {
    '$store.getters.getAudioSettings.treckId': {
      handler(value) {
        const audioSettings = this.$store.state.audioSettings;

        this.time.nowTime = '0:00';
        this.time.nowTimeNumber = 0;

        const setTotalTime = setInterval(() => {
          if (audioSettings.loadedAudio.duration) {
            let totalMinute = Math.floor(audioSettings.loadedAudio.duration / 60);
            let totalSecond = Math.floor(audioSettings.loadedAudio.duration - totalMinute * 60);

            this.time.totalTime = `${totalMinute < 10 ? '0' + totalMinute : totalMinute} : ${
              totalSecond < 10 ? '0' + totalSecond : totalSecond
            }`;
          } else {
            clearInterval(setTotalTime);
          }
        }, 100);

        if (this.time.nowTimeNumber !== 0) {
          audioSettings.loadedAudio.currentTime = this.time.nowTimeNumber;
        }

        const timer = setInterval(() => {
          let minute = Math.floor(audioSettings.loadedAudio.currentTime / 60);
          let second = Math.floor(audioSettings.loadedAudio.currentTime - minute * 60);
          this.time.nowTime = `${minute < 10 ? '0' + minute : minute} : ${second < 10 ? '0' + second : second}`;

          this.time.nowTimeNumber = audioSettings.loadedAudio.currentTime;
        }, 100);

        if (!value || this.time.nowTimeNumber === audioSettings.loadedAudio.duration) {
          clearInterval(timer);
        }
      },
    },
  },
  methods: {
    showPlayer() {
      this.hidenPlayer = !this.hidenPlayer;
    },
    showHideTrecks() {
      this.clipPath.trecksBlock = !this.clipPath.trecksBlock;
    },
    playTreck(treckIdInArray) {
      const audioSettings = this.$store.state.audioSettings;

      if (treckIdInArray === undefined) {
        for (let i = 0; i < audioSettings.album.trecks.length; i++) {
          const treckItem = audioSettings.album.trecks[i];

          if (treckItem.id === audioSettings.treckId) {
            treckIdInArray = i;
            break;
          }
        }
      }

      this.pauseTreck(treckIdInArray);

      if (audioSettings.treckId !== audioSettings.album.trecks[treckIdInArray].id) {
        audioSettings.treckId = audioSettings.album.trecks[treckIdInArray].id;
        audioSettings.loadedAudio = new Audio('/user_files/beats/trecks/' + audioSettings.album.trecks[treckIdInArray].path);
        audioSettings.loadedAudio.volume = audioSettings.volume;
        audioSettings.name = audioSettings.album.trecks[treckIdInArray].name;
        audioSettings.price = audioSettings.album.trecks[treckIdInArray].price;
      }

      audioSettings.loadedAudio.play();
      audioSettings.played = true;
      audioSettings.album.played = true;

      if (treckIdInArray) {
        audioSettings.album.trecks[treckIdInArray].played = true;
      } else {
        for (let i = 0; i < audioSettings.album.trecks.length; i++) {
          const treckItem = audioSettings.album.trecks[i];

          if (treckItem.id === audioSettings.treckId) {
            treckItem.played = true;
            return;
          }
        }
      }
    },
    pauseTreck(treckIdInArray) {
      const audioSettings = this.$store.state.audioSettings;

      audioSettings.loadedAudio.pause();
      audioSettings.played = false;
      audioSettings.album.played = false;

      for (let i = 0; i < audioSettings.album.trecks.length; i++) {
        const treckItem = audioSettings.album.trecks[i];

        if (treckItem.id === audioSettings.treckId) {
          treckItem.played = false;
          return;
        }
      }

      if (treckIdInArray !== undefined) {
        audioSettings.album.trecks[treckIdInArray].played = false;
      }
    },
  },
  computed: {
    showHidePlayer() {
      document.body.style.overflow = this.hidenPlayer ? 'unset' : 'hidden';
      return {
        hidden: this.hidenPlayer,
        showed: !this.hidenPlayer,
        'hide-full': !this.$store.getters.getAudioSettings.loadedAudio,
        'show-full': this.$store.getters.getAudioSettings.loadedAudio,
      };
    },
    showHideText() {
      return this.hidenPlayer ? 'Show' : 'Hidde';
    },
    clipPathTrecksBlock() {
      return {
        cliped: this.clipPath.trecksBlock,
      };
    },
    clipPathTrecksBlockText() {
      return this.clipPath.trecksBlock ? 'Show' : 'Hidden';
    },
  },
};
</script>
