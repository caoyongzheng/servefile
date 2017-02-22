import Vue from 'vue'
import Main from './Main.vue'
import store from './store.js'

new Vue({
  el: '#main',
  store,
  render: h => h(Main)
})
