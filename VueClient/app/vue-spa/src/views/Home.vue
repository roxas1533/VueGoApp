<style scoped>

  .col{
    --Edisplay:none;
    --Pdisplay:none;
    display: flex;
    width: 100%;
    position: relative;
    height: 98vh;
  }
  .Home{
    display:block;
    position: relative;
  }
  .HomeContents{
    background-color: #0f0a3b;
  }
#model-container{
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0px;
    left: 0px;
    pointer-events: none;
}
</style>
<template>
  <div class="Home" id="Home">
    <div class="col">
      <TalkArea @open="open" class="HomeContents"></TalkArea>
      <TimeLine @showProfile="showProfile" class="HomeContents" type="Global"></TimeLine>
      <TimeLine @showProfile="showProfile" class="HomeContents" type="Home"></TimeLine>
      <div class=overlay @click="closeEdit"></div>
      <!-- <ProfileEdit @close="closeEdit" class="ProfileEdit" id="ProfileEdit"></ProfileEdit> -->
    </div>
    <div id="model-container"></div>
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
  },
  methods: {
    open() {
      const mc = document.getElementById('model-container');
      const store = this.$store;
      const ComponentClass = Vue.extend(ProfileEdit);
      const instance = new ComponentClass({
        store,
      });
      this.$store.state.websocketUpdate = new Date().getTime();

      instance.$mount();
      // c.appendch;
      // c.insertBefore(instance.$el, c.firstChild);
      // mc.style.height = '100%';
      mc.appendChild(instance.$el);
    },

    showProfile(sc, id) {
      const mc = document.getElementById('model-container');
      const store = this.$store;
      const ComponentClass = Vue.extend(PersonalProfile);
      const instance = new ComponentClass({
        store,
        propsData: {
          screenname: sc,
          userID: id,
        },
      });
      instance.$on('showProfile', this.showProfile);
      instance.$on('showEditProfile', this.open);
      instance.$mount();
      // c.appendch;
      // c.insertBefore(instance.$el, c.firstChild);
      // mc.style.height = '100%';
      mc.appendChild(instance.$el);
    },
  },
};
</script>
