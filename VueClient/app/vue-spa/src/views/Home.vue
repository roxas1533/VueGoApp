<style scoped>

  .col{
    --Edisplay:none;
    --Pdisplay:none;
    display: flex;
    width: 100%;
    position: relative;

  }
  .Home{
    display:block;
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
  }
#model-container{
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0px;
    left: 0px;
    pointer-events: none;
}
  .ProfileEdit{
    display: var(--Edisplay);
  }
</style>
<template>
  <div class="Home" id="Home">
    <div class="col">
      <TalkArea @open="open" class="HomeContents"></TalkArea>
      <TimeLine @showProfile="showProfile" class="HomeContents"></TimeLine>
      <div class=overlay @click="closeEdit"></div>
      <ProfileEdit @close="closeEdit" class="ProfileEdit" id="ProfileEdit"></ProfileEdit>
    </div>
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
      instance.$mount();
      // c.appendch;
      // c.insertBefore(instance.$el, c.firstChild);
      // mc.style.height = '100%';
      mc.appendChild(instance.$el);
      // document.getElementById('Home').style.setProperty('--Pdisplay', 'block');
    },
  },
};
</script>
