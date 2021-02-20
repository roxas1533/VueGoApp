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
      background: rgb(20, 32, 43);
      border-left: solid thin #ececec;
      border-radius: 0 0 10px 0px;

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

    .TitleBar,.WrapEdit,.ffconent{
        display: flex;
        background-color: rgb(20, 32, 43);
        position: relative;
        align-items:center;
    }

    .WrapEdit{
        border-top: solid thin white;
        border-bottom: solid thin black;
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
        border: solid black thin;
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
        height: 500px;
        border-radius: 0px 0px 15px 15px;
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
    .ffconent{
      border-bottom: solid black thin;
      height: 70px;
      align-items: start;
    }
    .ff{
      border-left: thin black solid;
    }
    #ffButon{
      margin-left: auto;
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
            <div class="ffconent">
                <FFcontent title="TALKS" :num=0></FFcontent>
                <FFcontent title="FOLLOWING" class="ff" :num=1></FFcontent>
                <FFcontent title="FOLLOWERS" class="ff" :num=0></FFcontent>
                <div id="ffButon">
                </div>
            </div>
            <div class="userContent" id="userContent">
            </div>
        </div>
    </div>
</template>

<script>
import Vue from 'vue';
import talkContent from './ContentArea.vue';
import overray from './overray.vue';
import FFcontent from './FFContent.vue';
import FollowButton from './FollowButton.vue';
import unFollowButton from './UnFollowButton.vue';

export default {
  components: {
    overray,
    FFcontent,
  },
  props: {
    screenname: String,
    userID: Number,
  },
  async mounted() {
    let returnData = await window.fetch(`${this.$store.state.APIserver}/isFollow/${this.userID}`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${this.$store.state.JWTtoken}`,
        'Content-Type': 'text/plain',
      },
    }).then((res) => res.json());
    let ComponentClass;
    if (returnData.result === true) {
      ComponentClass = Vue.extend(unFollowButton);
    } else {
      ComponentClass = Vue.extend(FollowButton);
    }
    const instance = new ComponentClass();
    instance.$on('follow', this.follow);
    instance.$on('unfollow', this.unfollow);
    instance.$mount();
    document.getElementById('ffButon').appendChild(instance.$el);
    //----------------------------------------------------------------------------------------------------------
    const url = `${this.$store.state.APIserver}/getusers/${this.userID}/20/0`;
    returnData = await window.fetch(url, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${this.$store.state.JWTtoken}`,
        'Content-Type': 'application/json',
      },
    }).then((res) => res.json());
    if (returnData.result !== null) {
      returnData.result.forEach((element) => {
        this.AddContentEnd(element);
      });
      this.$store.state.loadID = returnData.result.pop().ID;
    }
  },

  data() {
    return {
      profile: `${this.$store.state.APIserver}/profile/${this.userID}.png?${(new Date()).getMinutes()}`,
      innerusername: this.$store.state.userName,
      Follow: '',
    };
  },
  methods: {
    async follow() {
      const url = `${this.$store.state.APIserver}/follow/${this.userID}`;
      const returnData = await window.fetch(url, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${this.$store.state.JWTtoken}`,
          'Content-Type': 'application/json',
        },
      }).then((res) => res.json());
      if (returnData.result) {
        const ComponentClass = Vue.extend(unFollowButton);
        const instance = new ComponentClass();
        instance.$on('unfollow', this.unfollow);
        instance.$mount();
        document.getElementById('ffButon').appendChild(instance.$el);
      }
    },
    async unfollow() {
      const url = `${this.$store.state.APIserver}/unfollow/${this.userID}`;
      const returnData = await window.fetch(url, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${this.$store.state.JWTtoken}`,
          'Content-Type': 'application/json',
        },
      }).then((res) => res.json());
      if (returnData.result) {
        const ComponentClass = Vue.extend(FollowButton);
        const instance = new ComponentClass();
        instance.$on('follow', this.follow);
        instance.$mount();
        document.getElementById('ffButon').appendChild(instance.$el);
      }
    },
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
    showProfile(sc, id) {
      this.$emit('showProfile', sc, id);
      this.$destroy();
      if (this.$el.parentNode) this.$el.parentNode.removeChild(this.$el);
    },
    close() {
      this.$destroy();
      if (this.$el.parentNode) this.$el.parentNode.removeChild(this.$el);
    },
  },
};
</script>
