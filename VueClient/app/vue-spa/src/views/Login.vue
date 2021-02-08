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
  .Emesssage{
    color: red;
  }
</style>

<template>
  <div class="home">
    <div class=RegiserContent>
      <button @click="RedirectRegister" id="newRegister">新規登録</button>
    </div>
    <div class="loginContents">
      <h1>ログイン</h1>
      <div class="Emesssage">{{ErrorMessage}}</div>
      <p class="loginContentLabel">メールアドレス：</p>
      <input v-model="address" @keydown.enter="Send" class="loginContent">
        <p class="loginContentLabel">パスワード：</p>
      <input v-model="pass" @keydown.enter="Send" type="password" class="loginContent">
      <button @click="Send" id="loginButton" class="loginContent">ログイン</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'Login',
  data() {
    return {
      address: '',
      pass: '',
      ErrorMessage: '',
    };
  },
  created() {

  },
  methods: {
    async Send() {
      this.check().then((res) => {
        if (res) console.log(1);
        else { this.ErrorMessage = 'メールアドレスアドレスまたはパスワードが違います。'; this.pass = ''; }
      });
      // if (this.check())console.log('ok!'); else console.log('no!');
    },
    async check() {
      let ok = true;
      if (!this.address) {
        ok = false;
      }
      if (!this.pass) {
        ok = false;
      }
      if (!ok) return ok;

      try {
        const url = 'http://localhost:8000/login';
        const data = {
          mail: this.address,
          pass: this.pass,
        };
        const returnData = await window.fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        }).then((res) => res.json());
        if (returnData.reslut === 'false') return false;
        this.$store.state.JWTtoken = returnData.JWT;
        this.$router.push({ name: 'Home' });
        return true;
      } catch (e) {
        console.log(e);
        return e;
      }
    },
    RedirectRegister() {
      this.$router.push({ name: 'register' });
    },
  },
};
</script>
