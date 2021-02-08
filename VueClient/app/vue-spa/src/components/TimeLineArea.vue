<style scoped>
    .TimeLineContents{
        width: 20em;
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
</style>

<template>
    <div class="TimeLineContents" id="TimeLineContents">
    </div>
</template>

<script>
import Vue from 'vue';
import talkContent from './ContentArea.vue';

let socket;
let timeLineC;
export default {
  destoryed() {
    socket.close();
  },
  components: {
  },
  methods: {
    async ScrollE() {
      if (this.$store.state.loadID !== 1) {
        if (Math.round((timeLineC.scrollTop / (timeLineC.scrollHeight - timeLineC.clientHeight)) * 100) > 80) {
          const url = `http://localhost:8000/get/20/${this.$store.state.loadID}`;
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
            this.$store.state.loadID = returnData.result.pop().ID;
          }
        }
      }
    },
    AddContentEnd(data) {
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
        instance.$mount();
        // c.appendch;
        // c.insertBefore(instance.$el, c.firstChild);
        timeLineC.appendChild(instance.$el);
      }
    },
    AddContentTop(data) {
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
        instance.$mount();
        timeLineC.insertBefore(instance.$el, timeLineC.firstChild);
      }
    },
  },
  async mounted() {
    timeLineC = document.getElementById('TimeLineContents');

    timeLineC.addEventListener('scroll', this.ScrollE);

    if (this.$store.state.JWTtoken !== '') {
      const url = 'http://localhost:8000/get/20/0';
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
        this.$store.state.loadID = returnData.result.pop().ID;
      }
      socket = new WebSocket('ws://localhost:8000/home/getTimeLine', [this.$store.state.JWTtoken]);
      const that = this;
      socket.onmessage = function (evt) {
        const data = JSON.parse(evt.data);
        that.AddContentTop(data);
      };
    }
  },
};
</script>
