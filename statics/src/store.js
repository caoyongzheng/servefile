import Vuex from 'vuex'
import Vue from 'vue'
import merge from 'lodash/merge'
import set from 'lodash/set'
import get from 'lodash/get'
import utils from './utils.js'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    tree: {},
    active: ''
  },
  mutations: {
    set(state, node) {
      const path = node.path
      if (path === '.' || path === '') {
        state.tree = merge({},state.tree, node)
      } else {
        const paths = utils.getPaths(path)
        const tree = {}
        set(tree, paths, node)
        state.tree = merge({},state.tree, tree)
      }
    },
    setActive(state, active) {
      state.active = active
    },
  },
  actions: {
    getNode(context, p = '') {
      fetch(`/treenode?path=${p}`)
      .then(r => r.json())
      .then(treenode => {
        context.commit('set', treenode)
      })
    },
  },
})

export default store
