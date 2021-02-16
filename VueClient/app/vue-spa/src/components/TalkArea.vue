<style scoped>
  textarea{
      resize: none;
  }
  .TalkComponent{
      display: block;
      margin-left: 0px;
  }
 #talkButton,#logoutButton{
     border-radius: 10px;
     font-size: 1.6em;
     display: inline-block;
 }
 .Talk{
     display: inline-block;
     text-align: right;
     position: relative;
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
  }

</style>
<template>
    <div class="Talk">
      <textarea class="TalkComponent" v-model="content" rows="8" cols="30" placeholder="Talk!"></textarea>
      <button class="TalkComponent" id="talkButton" @click="Talk">Talk</button>
      <div class="setting">
        <button class="TalkComponent" id="logoutButton" @click="set">
          <i class="fas fa-cog"></i>
        </button>
        <div class="settingString">設定</div>
        <div id="settingmenu">
          <a class="menuText">プロフィール</a>
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
      content: '',
    };
  },
  methods: {
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
      console.log(getComputedStyle(document.getElementById('settingmenu')).left);

      if (this.content !== '') {
        const url = `${this.$store.state.APIserver}/talk`;
        const data = {
          Content: this.content,
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
        this.content = '';
      }
    },
  },
};
</script>
