<style scoped>
  textarea{
      resize: none;
  }
  .TalkComponent{
      display: block;
      margin-left: 0px;
  }
  #logoutButton{
     border-radius: 10px;
     font-size: 1.6em;
  }
  #talkButton{
    color: white;
    height: 30px;
    border: none;
    border-radius: 100vh;
    outline: none;
    font-size: 1.2em;
    width: 80px;
    background-color: rgb(0, 183, 255);
    display: inline-block;
    cursor: pointer;
  }
  #talkButton:disabled{
    color: rgba(0, 183, 255,0.5);
    background-color: rgba(0, 183, 255,0.2);
    cursor:default ;
  }
  #talkButton:hover:active{
    background-color: rgb(0, 89, 255);
    outline: none;
  }
 .Talk{
     display: inline-block;
     text-align: right;
     position: relative;
     padding-left: 10px;
     padding-right: 10px;
     padding-top: 10px;
 }
  ::-webkit-scrollbar{
    width: 10px;
  }
  ::-webkit-scrollbar-track{
    background: #4444;
    border-left: solid 1px #ececec;
  }
  ::-webkit-scrollbar-thumb{
    background: #888;
    border-radius: 10px;
    box-shadow: inset 0 0 0 2px #888;
  }
  .setting{
      position: absolute;
      display: flex;
      bottom: 0em;
      left: 0px;
      width: 100%;
  }
  #logoutButton{
    margin-top: auto;
    border-radius: 0px;
    background: none;
    border: none;
  }
  #logoutButton:hover{
      color: white;
      background-color: #221788;

  }
  #logoutButton:focus{
    outline: none;
  }
  #logoutButton:focus ~ #settingmenu{
    /* position: absolute; */
    display: inline;
  }
  #logoutButton:hover ~ .settingString{
      display: inline;
      color: white;
  }
  .settingString{
    display: none;
    background-color: #221788;
    padding-left: 10px;
    padding-right: 10px;
  }
  :root{
    --width:38px;
  }
  #settingmenu:hover{
    display: inline;
  }
  #settingmenu{
     --width:0px;
      position: absolute;
      /* display: inline-block; */
      left: var(--width);
      display: none;
      text-align: center;
      bottom: 0px;
      background-color: #221788;
      pointer-events: auto;
      border-radius: 5px;
  }
  hr{
    margin: 0;
  }
  .menuText{
    display: block;
    color: white;
    margin-top: 10px;
    margin-bottom: 10px;
    padding-left: 10px;
    padding-right: 10px;
  }
  .menuText:hover{
    background-color: #20dcfd;
    cursor: pointer;
  }

</style>
<template>
    <div class="Talk">
      <textarea class="TalkComponent" @keyup="check" v-model="talkContent" rows="8" cols="30" placeholder="Talk!"></textarea>
      <button disabled class="TalkComponent" id="talkButton" @click="Talk" >Talk</button>
      <div class="setting">
        <button class="TalkComponent" id="logoutButton" @click="set">
          <i class="fas fa-cog"></i>
        </button>
        <div class="settingString">設定</div>
        <div id="settingmenu">
          <a class="menuText" @click="open">プロフィール</a>
          <hr>
          <a class="menuText" @click="logout">ログアウト</a>
        </div>
      </div>

    </div>
</template>
<script>
export default {
  mounted() {
    const w = document.getElementById('logoutButton').clientWidth;
    document.documentElement.style.setProperty('--width', `${w}px`);
    document.getElementById('settingmenu').style.left = `${w}px`;
  },

  data() {
    return {
      talkContent: '',
    };
  },
  methods: {
    check() {
      if (this.talkContent !== '') {
        document.getElementById('talkButton').removeAttribute('disabled');
      } else {
        document.getElementById('talkButton').setAttribute('disabled', true);
      }
    },
    open() {
      this.$emit('open');
    },
    set() {
      const w = document.getElementById('logoutButton').clientWidth;
      document.documentElement.style.setProperty('--width', w);
    },
    logout() {
      this.$store.state.JWTtoken = '';
      this.$store.state.loadID = 0;
      this.$router.push({ name: 'Login' });
    },
    async Talk() {
      if (this.talkContent !== '') {
        const url = `${this.$store.state.APIserver}/talk`;
        const data = {
          Content: this.talkContent,
        };
        const returnData = await window.fetch(url, {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${this.$store.state.JWTtoken}`,
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        }).then((res) => res.json());
        if (returnData.result !== 'ok') {
          this.$store.state.JWTtoken = '';
          this.$router.push({ name: 'Login' });
        }
        this.talkContent = '';
        document.getElementById('talkButton').setAttribute('disabled', true);
      }
    },
  },
};
</script>
