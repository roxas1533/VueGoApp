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
    #logoutButton{
        position: absolute;
        bottom: 0em;
        left: 0em;
    }
</style>
<template>
    <div class="Talk">
            <textarea class="TalkComponent" v-model="content" rows="8" cols="30" placeholder="Talk!"></textarea>
            <button class="TalkComponent" id="talkButton" @click="Talk">Talk</button>
            <button class="TalkComponent" id="logoutButton" @click="logout">ログアウト</button>
    </div>
</template>
<script>
export default {
  data() {
    return {
      content: '',
    };
  },
  methods: {
    logout() {
      this.$store.state.JWTtoken = '';
      this.$store.state.loadID = 0;
      this.$router.push({ name: 'Login' });
    },
    async Talk() {
      if (this.content !== '') {
        const url = 'http://localhost:8000/talk';
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
