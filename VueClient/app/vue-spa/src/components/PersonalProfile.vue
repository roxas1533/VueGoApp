<style scoped>
    .fa-times{
        color: rgb(0, 140, 255);
        font-size: 2em;
    }
    i:hover{
        cursor: pointer;
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
    .TitleBar{
        border-radius: 15px 15px 0px 0px;
        padding-left: 15px;
        padding-right: 15px;
    }

    .TitleBar,.WrapEdit{
        display: flex;
        background-color: rgb(20, 32, 43);
        position: relative;
        align-items:center;
    }

    .WrapEdit{
        border-top: solid 2px white;
        border-bottom: solid 1px black;
        display: flex;
    }
    .edit{
        display: block;
        margin-left: 39px;
        margin-right: 39px;
        padding-bottom: 20px;
        width: 100%;
    }
    .close{
        width: 100% ;
        height: 50px ;
        align-items:center;
        justify-content: flex-end;
        display:flex;
        margin-left: auto;
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

    .boxPadding{
        padding: 12px;
    }
    .wrapImage{
        border: solid black 1px;
        display: inline-block;
        position: relative;
        margin-top: 20px;
        margin-bottom: 20px;
    }
    .userImage{
        display: inline-block;
        width: 112px;
        height: 112px;
        clear:all;
    }
    .screenname{
        font-weight: bold;
        color: white;
        font-size: 1.8em;
    }
    .userID{
        color: white;
    }
    .userContent{
        overflow-y:auto;
        background-color:rgb(20, 32, 43);
        height: 550px;
        border-radius: 0px 0px 15px 15px;
    }
</style>
<template>
    <div>
        <div class="TitleBar">
            <div class="close">
                <div @click="close" class="t-middle">
                    <div class="t-middle-middle">
                        <i @click="close" class="tims fas fa-times"></i>
                    </div>
                </div>
            </div>
        </div>
        <div class="WrapEdit">
            <div class=edit>
                <div class="wrapImage">
                    <img :src="profile" class="userImage">
                </div>
                <div class="screenName">
                    <span class="screenname">{{screenname}}</span>
                </div>
                <div class="wrapuserID">
                    <span class="userID">@{{userID}}</span>
                </div>
            </div>
        </div>
        <div class="userContent" id="userContent">
        </div>
    </div>
</template>

<script>
import Vue from 'vue';
import talkContent from './ContentArea.vue';

export default {
  props: {
    screenname: String,
    userID: Number,
  },
  async mounted() {
    const url = `${this.$store.state.APIserver}/get/20/0`;
    const returnData = await window.fetch(url, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${this.$store.state.JWTtoken}`,
        'Content-Type': 'application/json',
      },
    }).then((res) => res.json());
    if (returnData.result !== null) {
      returnData.result.forEach((element) => {
        this.AddContentEnd(element);
        // console.log(element);
      });
      this.$store.state.loadID = returnData.result.pop().ID;
    }
  },

  data() {
    return {
      profile: `${this.$store.state.APIserver}/profile/${this.$store.state.userId}?${(new Date()).getMinutes()}`,
      innerusername: this.$store.state.userName,
    };
  },
  methods: {
    AddContentEnd(data) {
    // if (data.Type === 'push') {
      const ComponentClass = Vue.extend(talkContent);
      const instance = new ComponentClass({
        propsData: {
          username: data.name,
          content: data.Content,
          id: data.UserID,
          time: data.Time,
        },
      });
      instance.$on('showProfile', this.showProfile);
      instance.$mount();
      // c.appendch;
      // c.insertBefore(instance.$el, c.firstChild);
      document.getElementById('userContent').appendChild(instance.$el);
      // }
    },
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
  },
};
</script>
