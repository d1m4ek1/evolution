<template>
  <section class="container settings_content">
    <your-albums :albums-data="albums"></your-albums>
    <create-album></create-album>
  </section>
</template>

<script>
import YourAlbums from "./inBeats/YourAlbums.vue";
import CreateAlbum from "./inBeats/CreateAlbum.vue";

export default {
  data() {
    return {
      albums: null,
    };
  },
  methods: {
    getAlbums() {
      const getAlbumsPromise = new Promise((resolve, reject) => {
        setTimeout(() => {
          resolve({
            names: ["1", "2", "3"],
            covers: ["/scrshot.png", "/scrshot.png", "/scrshot.png"],
            dateOfRelease: ["2022-08-12", "2022-08-12", "2022-08-12"],
            typeTrecks: ["beats", "sounds", "beats"],
            isPrivate: [false, true, false],
            albumId: ["fdg234fgvcb", "fddfg4445vcb", "32452hgf45"],
            trecks: [
              [
                {
                  name: "adele",
                  path: "/adele-hello.mp3",
                  id: 1,
                  price: 233,
                  played: false,
                },
                {
                  name: "adele",
                  path: "/adele-hello.mp3",
                  id: 2,
                  price: 233,
                  played: false,
                },
                {
                  name: "adele",
                  path: "/adele-hello.mp3",
                  id: 3,
                  price: 233,
                  played: false,
                },
              ],
              [
                {
                  name: "be-mine",
                  path: "/ofenbach-be-mine.mp3",
                  id: 4,
                  price: 233,
                  played: false,
                },
              ],
              [
                {
                  name: "adele",
                  path: "/adele-hello.mp3",
                  id: 5,
                  price: 233,
                  played: false,
                },
                {
                  name: "be-mine",
                  path: "/ofenbach-be-mine.mp3",
                  id: 6,
                  price: 233,
                  played: false,
                },
              ],
            ],
          });
        }, 500);
      });

      getAlbumsPromise.then((response) => {
        this.albums = response;
      });
    },
  },
  created() {
    this.getAlbums();
  },
  components: {
    CreateAlbum,
    YourAlbums,
  },
};
</script>
