<style scoped>
    .TimeLineContents{
        width: 20em;
        overflow:auto;
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
    <div class="TimeLineContents">
        <talkContent></talkContent>

    </div>
</template>

<script>
import talkContent from './ContentArea.vue';

export default {
  created() {
    const socket = new WebSocket('ws://localhost:8000/getTimeLine');
    socket.onopen = function () {
      console.log('ok');
    };
    socket.onmessage = function (evt) {
      console.log(evt.data);
    };
  },
  components: {
    talkContent,
  },
};
</script>
