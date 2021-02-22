import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    JWTtoken: '',
    loadID: 0,
    // APIserver: 'http://roxas-71a9bf3a.localhost.run',
    APIserver: 'http://localhost:8000',
    userName: '',
    userId: 0,
    // websocketserver: 'roxas-71a9bf3a.localhost.run',
    timelineKind: ['Global', 'Home'],
    websocketserver: 'localhost:8000',
    websocketUpdate: new Date().getTime(),
    favlist: { favid: 0, state: false },
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  },
  getters: {
    getwebsocketUpdate: (state) => state.websocketUpdate,
    getFavlist: (state) => state.favlist,
  },
});
