<style scoped>
    .Timeline{
      width: 20em;
      margin-left: 5px;
    }
    .TimeLineContents{
       height: calc(100% - 50px);
        overflow-y:auto;
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

    .header{
      height: 50px;
      border-bottom: thin solid black;
      box-shadow: 0px 2px 10px black;
      text-align: left;
      /* justify-content: center; */
      align-items:center;
      display: flex;
    }
    .wrapFas{
      width: 24px;
      margin-left: 10px;
    }
    .homeString{
      color:white;
      font-size: 1.2em;
      font-weight: bold;

    }
</style>

<template>
  <div class="Timeline">
      <div class="header">
        <div class="wrapFas">
          <i class="fas fa-home"></i>
        </div>
        <span class="homeString">
          {{type}}
        </span>
      </div>
      <div class="TimeLineContents" ref="TimeLineContents"></div>
  </div>
</template>

<script>
import Vue from 'vue';
import talkContent from './ContentArea.vue';

let socket;
let snc = false;
export default {
  destoryed() {
    socket.close();
  },
  props: {
    type: String,
  },
  components: {
  },
  data() {
    const loadID = 0;
    return { loadID };
  },
  methods: {
    async ScrollE() {
      if (snc === false) {
        if (this.$loadID !== 1) {
          if (Math.round((this.$refs.TimeLineContents.scrollTop / (this.$refs.TimeLineContents.scrollHeight
          - this.$refs.TimeLineContents.clientHeight)) * 100) > 80) {
            let url;
            switch (this.type) {
              case 'Global':
                url = `${this.$store.state.APIserver}/get/20/${this.$loadID}`;
                break;
              case 'Home':
                url = `${this.$store.state.APIserver}/getUsersTimeLine/20/${this.$loadID}`;
                break;
              default:
                break;
            }
            snc = true;
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
              });
              this.$loadID = returnData.result.pop().ID;
              snc = false;
            }
          }
        }
      }
    },
    makeContent(data) {
      if (data.Type === 'push') {
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
        return instance;
      }
      return null;
    },
    AddContentEnd(data) {
      this.$refs.TimeLineContents.appendChild(this.makeContent(data).$el);
    },
    showProfile(username, uid) {
      this.$emit('showProfile', username, uid);
    },
    AddContentTop(data) {
      this.$refs.TimeLineContents.insertBefore(this.makeContent(data).$el, this.$refs.TimeLineContents.firstChild);
    },
  },
  async mounted() {
    this.$refs.TimeLineContents.addEventListener('scroll', this.ScrollE);
    if (this.$store.state.JWTtoken !== '') {
      let url;
      let socketUrl;
      switch (this.type) {
        case 'Global':
          url = `${this.$store.state.APIserver}/get/20/0`;
          socketUrl = `ws://${this.$store.state.websocketserver}/home/streamGlobalTimeLine`;
          break;
        case 'Home':
          url = `${this.$store.state.APIserver}/getUsersTimeLine/20/0`;
          socketUrl = `ws://${this.$store.state.websocketserver}/home/streamHomeTimeLine`;
          break;
        default:
          break;
      }
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
        });
        this.$loadID = returnData.result.pop().ID;
      }
      socket = new WebSocket(socketUrl, [this.$store.state.JWTtoken]);
      socket.onmessage = (evt) => {
        const data = JSON.parse(evt.data);
        this.AddContentTop(data);
      };
    }
  },
};
</script>
