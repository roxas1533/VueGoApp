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
    userId: '1',
    // websocketserver: 'roxas-71a9bf3a.localhost.run',
    websocketserver: 'localhost:8000',
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  },
});
