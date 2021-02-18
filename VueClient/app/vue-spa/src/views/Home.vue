<style scoped>

  .Home{
    --display:none;
    display: flex;
    width: 100%;
    position: relative;
  }
  .HomeContents{
    height: 95vh;
    margin-right: 1em;
    margin-left: 1em;
    margin-top: 1em;
    background-color: #0f0a3b;
  }
  .ProfileEdit{
    position: absolute;
    width: 58%;
    max-width: 550px;
    top: 50%;
    left: 50%;
    transform: translateY(-50%) translateX(-50%);
    margin: auto;
    display: var(--display);
  }
  .overlay{
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.7);
    display: var(--display);
  }
</style>
<template>
  <div class="Home" id="Home">
    <TalkArea @open="open" class="HomeContents"></TalkArea>
    <TimeLine class="HomeContents"></TimeLine>
    <div class=overlay></div>
    <ProfileEdit @close="closeEdit" class="ProfileEdit" id="ProfileEdit"></ProfileEdit>
  </div>
</template>

<script>
import TalkArea from '../components/TalkArea.vue';
import TimeLine from '../components/TimeLineArea.vue';
import ProfileEdit from '../components/ProfileEdit.vue';

export default {
  created() {
    if (this.$store.state.JWTtoken === '') {
      this.$router.push({ name: 'Login' });
    }
  },
  components: {
    TalkArea,
    TimeLine,
    ProfileEdit,
  },
  methods: {
    open() {
      document.getElementById('Home').style.setProperty('--display', 'block');
    },
    closeEdit() {
      document.getElementById('Home').style.setProperty('--display', 'none');
    },
  },
};
</script>
