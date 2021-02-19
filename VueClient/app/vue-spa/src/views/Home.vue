<style scoped>

  .Home{
    --Edisplay:none;
    --Pdisplay:none;
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
  .ProfileEdit,#model-container{
    position: absolute;
    width: 58%;
    max-width: 550px;
    top: 50%;
    left: 50%;
    transform: translateY(-50%) translateX(-50%);
    margin: auto;
  }
  .ProfileEdit{
    display: var(--Edisplay);
  }
  .PersonalProfile{
    display: var(--Pdisplay);
  }
  .overlay{
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.7);
    display: var(--Edisplay);
  }
</style>
<template>
  <div class="Home" id="Home">
    <TalkArea @open="open" class="HomeContents"></TalkArea>
    <TimeLine @showProfile="showProfile" class="HomeContents"></TimeLine>
    <div class=overlay @click="closeEdit"></div>
    <ProfileEdit @close="closeEdit" class="ProfileEdit" id="ProfileEdit"></ProfileEdit>
    <div id="model-container"></div>
    <!-- <PersonalProfile @close="closeEdit" class="PersonalProfile" id="PersonalProfile" :screenname="name" :userID="uid"></PersonalProfile> -->
  </div>
</template>

<script>
import Vue from 'vue';
import TalkArea from '../components/TalkArea.vue';
import TimeLine from '../components/TimeLineArea.vue';
import ProfileEdit from '../components/ProfileEdit.vue';
import PersonalProfile from '../components/PersonalProfile.vue';

export default {
  data() {
    return {
      name: '',
      uid: 0,
    };
  },
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
      document.getElementById('Home').style.setProperty('--Edisplay', 'block');
    },
    closeEdit() {
      document.getElementById('Home').style.setProperty('--Edisplay', 'none');
    },
    showProfile(sc, id) {
      const store = this.$store;
      const ComponentClass = Vue.extend(PersonalProfile);
      const instance = new ComponentClass({
        store,
        propsData: {
          screenname: sc,
          userID: id,
        },
      });
      // instance.$on('showProfile', this.showProfile);
      instance.$mount();
      // c.appendch;
      // c.insertBefore(instance.$el, c.firstChild);
      document.getElementById('model-container').appendChild(instance.$el);
      // document.getElementById('Home').style.setProperty('--Pdisplay', 'block');
    },
  },
};
</script>
