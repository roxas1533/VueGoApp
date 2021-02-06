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
  .Error{
      color: red;
  }
</style>

<template>
  <div class="home">
        <div class=RegiserContent>
            <button @click="RedirectRegister" id="newRegister">ログイン</button>
        </div>
        <div class="loginContents">
            <h1>新規登録</h1>
            <p class="loginContentLabel">ユーザ名：<span class="Error">{{userError}}</span></p>
            <input v-model="username" class="loginContent">
            <p class="loginContentLabel">メールアドレス：<span class="Error">{{mailError}}</span></p>
            <input v-model="address" type="email" class="loginContent">
            <p class="loginContentLabel">パスワード：<span class="Error">{{passError}}</span></p>
            <input v-model="password" type="password" class="loginContent">
            <p class="loginContentLabel">パスワード(確認用)：<span class="Error">{{pass2Error}}</span></p>
            <input v-model="password2" type="password" class="loginContent">
            <button @click="Send" id="loginButton" class="loginContent">登録</button>
         </div>
  </div>
</template>

<script>

export default {
  name: 'Register',
  data() {
    return {
      username: '',
      address: '',
      password: '',
      password2: '',
      userError: '',
      mailError: '',
      passError: '',
      pass2Error: '',
    };
  },
  methods: {
    Send() {
      if (this.Check() === true) console.log('ok!'); else console.log('だめ');
    },
    Check() {
      let ok = true;
      if (!this.username) {
        this.userError = '入力してください';
        ok = false;
      } else this.userError = '';
      if (!this.address) {
        this.mailError = '入力してください';
        ok = false;
      } else this.mailError = '';
      if (!this.password) {
        this.passError = '入力してください';
        ok = false;
      } else this.passError = '';
      if (this.password2 !== this.password) {
        this.pass2Error = '一致しません';
        ok = false;
      } else this.pass2Error = '';
      return ok;
    },
    RedirectRegister() {
      this.$router.push({ name: 'Login' });
    },
  },
};
</script>
