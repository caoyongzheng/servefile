import Vuex from 'vuex'
import Vue from 'vue'
import merge from 'lodash/merge'
import set from 'lodash/set'
import get from 'lodash/get'
import utils from './utils.js'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    tree: { open: true },
    active: '',
    textarea: '',
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
    setTextArea(state, content) {
      state.textarea = content
    }
  },
  actions: {
    getNode(context, p = '') {
      return fetch(`${process.env.BaseURL}/treenode?path=${p}`)
      .then(r => r.json())
      .then(treenode => {
        context.commit('set', treenode)
        return treenode
      })
    },
    onNodeClick(context, node) {
      if (node.loaded) {
        if (node.isDir) {
          node.open = !node.open
          context.commit('set', node)
          context.commit('setActive', node.path)
        } else {
          context.commit('setActive', node.path)
          context.commit('setTextArea', node.content)
        }
        return
      }
      fetch(`${process.env.BaseURL}/treenode?path=${node.path}`)
      .then(r => r.json())
      .then(n => {
        node = merge(node, n)
        node.loaded = true
        if (node.isDir) {
          node.open = !node.open
          context.commit('set', node)
          context.commit('setActive', node.path)
        } else {
          context.commit('setActive', node.path)
          context.commit('setTextArea', node.content)
        }
      })
    },
  },
  getters: {
    curnode(state) {
      const path = state.active
      if (!path || path === '' || path === '.') {
        return state.tree
      }
      const paths = utils.getPaths(path)
      return get(state.tree, paths)
    }
  },
})

export default store
