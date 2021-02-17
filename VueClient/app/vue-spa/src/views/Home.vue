<style scoped>
  .Home{
    display: flex;
    width: 100%;
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

  }
</style>
<template>
  <div class="Home">
    <TalkArea class="HomeContents"></TalkArea>
    <TimeLine class="HomeContents"></TimeLine>
    <ProfileEdit class="ProfileEdit"></ProfileEdit>
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
};
</script>
