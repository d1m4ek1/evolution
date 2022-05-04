<template>
  <section class="container settings_content">
    <!-- BACKUPS -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Ключи восстановления аккаунта</h2>
        <p class="note">
          Поля ключей восстановления аккаунта нужно заполнить все, либо оставить
          пустыми
        </p>
      </div>
      <div class="change_content">
        <h3>Ключи восстановления #1</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>erwe</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input type="text" maxlength="100" placeholder="Первый ключ..."/>
          </div>
        </div>
        <h3>Ключи восстановления #2</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>erwe</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input type="text" maxlength="100" placeholder="Второй ключ..."/>
          </div>
        </div>
        <h3>Ключи восстановления #3</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>erwe</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input type="text" maxlength="100" placeholder="Третий ключ..."/>
          </div>
        </div>
        <h3>Ключи восстановления #4</h3>
        <div class="inlineblock">
          <div class="inlineblock_left">
            <div class="old_text">
              <h2>erwe</h2>
            </div>
          </div>
          <div class="inlineblock_right">
            <input type="text" maxlength="100" placeholder="Четвертый ключ..."/>
          </div>
        </div>
      </div>
    </div>
    <!-- CHANGE PASSWORD -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Смена пароля</h2>
      </div>
      <div class="change_content">
        <div class="change_content__title">
          <h3>Новый пароль</h3>
          <p class="note">
            Безопаснее всего создать пароль длиною более 16 символов, что и
            делает генератор
          </p>
        </div>
        <div class="inlineblock">
          <input type="password" placeholder="Новый пароль..."/>
        </div>
        <div class="inlineblock">
          <button class="btn">Сгенерировать пароль</button>
        </div>
      </div>
    </div>
    <!-- CHANGE PASSWORD -->
    <div class="change_section">
      <div class="change_section__title">
        <h2>Смена электронной почты</h2>
      </div>
      <div class="change_content">
        <div class="change_content__title">
          <h3>Новая электронная почта</h3>
          <p class="note">
            Безопаснее всего создать пароль длиной более 16 символов, что и
            делает генератор
          </p>
        </div>
        <div class="inlineblock">
          <input type="email" placeholder="example@example.com"/>
        </div>
        <div class="inlineblock">
          <button @click="completePassword()" class="btn">Сегенерировать пароль</button>
        </div>
      </div>
    </div>
    <!-- SAVE BUTTON OR DEFAULT SETTINGS -->
    <div class="change_section btns">
      <div class="change_content">
        <div class="change_content__title">
          <h3>Подтверждение</h3>
          <p class="note">
            Только с подтверждением пароля, можно изменить персональные данные.<br>
            <a href="#">Забыли пароль?</a>
          </p>
        </div>
        <div class="inlineblock">
          <input type="password" />
        </div>
      </div>
      <button @click="completePassword()" class="btn">Подвердить пароль</button>
      <button
        v-if="completePassword.completed"
        @click="saveSettings()"
        class="btn"
      >
        Сохранить
      </button>
      <button @click="resetSettings()" class="btn btn-red">
        Сбросить все изменения
      </button>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      completePassword: {
        password: '',
        completed: false,
      },
    };
  },
};
</script>
