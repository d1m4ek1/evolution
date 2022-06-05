<template>
  <div class="subcontent show_content">
    <section class="container title">
      <h2>Информация обо мне</h2>
    </section>
    <section class="container subprofile_items">
      <div class="subprofile_item">
        <h3>{{ title }}</h3>
        <div class="subprofile_item__content" v-html="content">
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    content: String,
  },
};
</script>
