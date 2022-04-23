<template>
  <div class="directory show_content">
    <section class="back_page">
      <button @click="hideComponent()" class="btn">Назад</button>
    </section>
    <section class="mini_title">
      <h2>{{ documentation[paragraph][typeContent].title }}</h2>
    </section>
    <section class="content_explanation">
      <template v-if="documentation[paragraph][typeContent].content != []">
        <template
          v-for="(item, idx) in documentation[paragraph][typeContent].content"
        >
          <p v-if="item.tag == 'p'" :key="idx">
            {{ item.value }}
          </p>
          <ul v-if="item.tag == 'ul'" :key="idx">
            <li v-for="(liItem, idxsmall) in item.value" :key="idxsmall">
              {{ liItem }}
            </li>
          </ul>
          <ol v-if="item.tag == 'ol'" :key="idx">
            <li v-for="(liItem, idxsmall) in item.value" :key="idxsmall">
              {{ liItem }}
            </li>
          </ol>
        </template>
      </template>
    </section>
  </div>
</template>

<script>
export default {
  props: ["paragraph", "typeContent"],
  data() {
    return {
      documentation: {
        greetings: {
          hello: {
            title: "Приветствие!",
            content: [
              {
                tag: "p",
                value: "Какое-либо приветствие!",
              },
            ],
          },
        },
        1: {
          basic: {
            title: "Основная документация",
            content: [
              {
                tag: "p",
                value: "Доброго времени суток!",
              },
              {
                tag: "p",
                value:
                  "В справочнике вы сможете ознакомиться с юридическики вопросами",
              },
              {
                tag: "p",
                value:
                  "А также получить вопросы по тех. поддержке и многое другое",
              },
            ],
          },
          serviceInformation: {
            title: "Информация о сервисе",
          },
          serviceFounders: {
            title: "Основатели сервиса"
          }
        },
        2: {
          privacyPolicy: {
            title: "Политика конфиденциальности",
            content: [
              {
                tag: "ul",
                value: [
                  "Продолжение следует",
                  "Продолжение следует",
                  "Продолжение следует",
                ],
              },
            ],
          },
          termsOfUs: {
            title: "Пользовательское соглашение",
            content: [],
          },
        },
        9: {
          profile: {
            title: "Профиль",
            content: [
              {
                tag: "p",
                value: "Описание"
              }
            ]
          },
          commerce: {
            title: "Коммерция",
            content: [
              {
                tag: "p",
                value: "Описание"
              }
            ]
          },
          community: {
            title: "Сообщество",
            content: [
              {
                tag: "p",
                value: "Описание"
              }
            ]
          },
          settings: {
            title: "Настройки",
            content: [
              {
                tag: "p",
                value: "Описание"
              }
            ]
          },
        }
      },
    };
  },
  methods: {
    hideComponent: function () {
      this.$emit("hidecomponent", {
        hide: true,
      });
      window.history.pushState(null, "Directory", "/directory");
    },
    setTitle: function() {
      document.title = this.documentation[this.paragraph][this.typeContent].title
    }
  },
  updated: function() {
    this.setTitle()
  },
  mounted: function() {
    this.setTitle()
  }
};
</script>
