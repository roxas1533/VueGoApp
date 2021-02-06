<style scoped>
  .RegiserContent{
    text-align: right;
  }
  .loginContents{
    display: inline-block;
    width:300pt;
  }
  #newRegister{
    border-radius: 5px;
    font-size: 1.3em;
  }
  #loginButton{
    border-radius: 5px;
    font-size: 1.5em;
  }
  input{
    font-size: 1.2em;
    height: 1.5em;
  }
  .loginContent{
    width: 100%;
    display: block;
    margin-bottom: 1em;
  }
  .loginContentLabel{
    text-align: left;
  }
</style>

<template>
  <div class="home">
    <div class=RegiserContent>
      <button @click="RedirectRegister" id="newRegister">新規登録</button>
    </div>
    <div class="loginContents">
      <h1>ログイン</h1>
      <p class="loginContentLabel">メールアドレス：</p>
      <input @keydown.enter="Send" class="loginContent">
        <p class="loginContentLabel">パスワード：</p>
      <input @keydown.enter="Send" type="password" class="loginContent">
      <button @click="Send" id="loginButton" class="loginContent">ログイン</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'Login',
  methods: {
    Send() {
    },
    RedirectRegister() {
      this.$router.push({ name: 'register' });
    },
  },
};
</script>
