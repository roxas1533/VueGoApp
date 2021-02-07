<style scoped>
  .Home{
    display: flex;
  }
  .HomeContents{
    height: 95vh;
    margin-right: 1em;
    margin-left: 1em;
    margin-top: 1em;
    background-color: #0a0014;
  }
</style>
<template>
  <div class="Home">
    <TalkArea class="HomeContents"></TalkArea>
    <TimeLine class="HomeContents"></TimeLine>
  </div>
</template>

<script>
import TalkArea from '../components/TalkArea.vue';
import TimeLine from '../components/TimeLineArea.vue';

export default {
  components: {
    TalkArea,
    TimeLine,
  },
};
</script>
