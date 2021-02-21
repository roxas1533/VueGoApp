<style scoped>
    .fa-times{
        color: rgb(0, 140, 255);
        font-size: 2em;
    }
    i:hover{
        cursor: pointer;
    }
    h2,.detail{
        margin: 0px;
        max-width: 100%;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    h2{
        color: white;
    }
    input:focus{
        outline: none;
    }
    .TitleBar{
        border-radius: 15px 15px 0px 0px;
    }

    .TitleBar,.edit{
        display: flex;
        background-color: rgb(20, 32, 43);
        padding-left: 15px;
        padding-right: 15px;
        position: relative;
        width: 100%;
    }
    .close{
        width: 80px ;
        height: 50px ;
        text-align: left;
        align-items:center;
        display:flex;
    }
   .t-middle{
        width: 40px ;
        height: 40px ;
        text-align: left;
        align-items:center;
        justify-content:center;
        display:flex;
    }
    .t-middle:hover{
        background-color: rgba(0, 183, 255,0.1);
        border-radius: 200px;
        cursor: pointer;
    }
    .editprofile{
        align-items:center;
        /* justify-content:center; */
        display:flex;
        width:100% ;
    }
    span{
        color: white;
        height: 30px;
        border: none;
        border-radius: 100vh;
        font-size: 1.2em;
        width: 80px;
        background-color: rgb(0, 183, 255);
    }
    span:hover{
        cursor: pointer;
        background-color: rgb(0, 124, 173);
    }
    button:focus{
        outline: none;
    }
    .save{
        margin-left:0px;
         align-items:center;
        display:flex;
    }
    .edit{
        border-radius: 0px 0px 15px 15px;
        justify-content:center;
        display: flexbox;
        flex-direction: column;
        border-top: solid 2px white;
    }
    .name{
        position: absolute;
        font-size: 1.35em;
        height: 1.35em;
        width: 99%;
        bottom: 0px;
        box-sizing: border-box;
        background: none;
        border: none;
        left: 0;
        right: 0;
        margin: auto;
        pointer-events:auto;
    }
    .boxPadding{
        padding: 12px;
    }
    .nameBox{
        position: relative;
        height: 3.5em;
        width: 100%;
        display: flex;
        border: solid 2px black;
        border-radius: 5px;
        background-color: rgb(18, 18, 18);
    }
    .nameBox:focus-within {
        border:solid rgb(0, 183, 255) 2px;
    }
    .nameBox:focus-within .detail{
        color: rgb(0, 183, 255);
    }
    .wrapImage{
        border: solid black 1px;
        display: inline-block;
        position: relative;
        width: 112px;
        height: 112px;
    }
    .fa-plus-circle{
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translateY(-50%) translateX(-50%);
    }
    .wrapPlus{
        width: 50px;
        height:  50px;
        position: absolute;
        display: inline-block;
        top: 50%;
        left: 50%;
        transform: translateY(-50%) translateX(-50%);
        margin: auto;
    }
    .wrapPlus:hover{
        cursor: pointer;
        background-color: rgba(18, 18, 18, 0.5);
        border-radius: 100px;
    }
    .userImage{
        display: inline-block;
        width: 112px;
        height: 112px;
        clear:all;
    }
    .overray{
        position: absolute;
    }
    .main{
        position: fixed;
        width: 58%;
        max-width: 550px;
        top: 50%;
        left: 50%;
        transform: translateY(-50%) translateX(-50%);
        margin: auto;
        position: absolute;
        pointer-events: auto;
        z-index: 300;
    }
</style>
<template>
    <div>
        <overray class="overray" @close="close"></overray>
        <div class="main">
            <div class="TitleBar">
                <div class="close">
                    <div @click="close" class="t-middle">
                        <div class="t-middle-middle">
                            <i @click="close" class="tims fas fa-times"></i>
                        </div>
                    </div>
                </div>
                <div class="editprofile">
                    <h2>プロフィールを編集</h2>
                </div>
                <div class="save" @click="saveProfile">
                    <span>保存</span>
                </div>
            </div>
            <div class=edit>
                <div class="wrapImage">
                    <img :src="profile" class="userImage">
                    <label for="file">
                        <div class="wrapPlus">
                        <input @change="se" type="file" id="file" style="display:none;" accept=".png">
                        <i class="fas fa-plus-circle"></i>
                        </div>
                    </label>

                </div>
                <div class="boxPadding">
                    <div class="nameBox" v-on:click="focus">
                        <p class="detail">名前</p>
                        <input type="text" class="name" v-model="innerusername">
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import overray from './overray.vue';

export default {
  props: {
  },
  components: {
    overray,
  },
  data() {
    return {
      profile: `${this.$store.state.APIserver}/profile/${this.$store.state.userId}.png?${(new Date()).getMinutes()}`,
      innerusername: this.$store.state.userName,
    };
  },
  methods: {
    se(e) {
      if (e.target.files[0].type === 'image/png') {
        const reader = new FileReader();
        reader.onload = this.onload;
        reader.readAsDataURL(e.target.files[0]);
      }
    },
    onload(ev) {
      this.profile = ev.target.result;
    },
    focus(e) {
      const c = e.target.lastElementChild;
      if (c) c.focus();
    },
    close() {
      this.$destroy();
      if (this.$el.parentNode) this.$el.parentNode.removeChild(this.$el);
    },
    async saveProfile() {
      const data = {
        UserName: this.innerusername,
        ProfileImage: this.profile,
      };
      const returnData = await window.fetch(`${this.$store.state.APIserver}/update`, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${this.$store.state.JWTtoken}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      }).then((res) => res.json());
      if (returnData.result === 'ok') {
        this.close();
      } else {
        console.log('アップデート失敗');
      }
    },
  },
};
</script>
