<template lang="html">
  <div class="tree">
    <div
      class="header"
      v-on:click="toggle"
      :style="{ paddingLeft: level * 15 + 8 + 'px' }"
      :class="{ active }"
    >
      <ArrowIcon v-if="node.isDir" :down="!!node.open" />
      <span :class="{ leaf: !node.isDir }">{{ node.name }}</span>
    </div>
    <Tree v-for="n in node.children" v-if="node.open" :path="n.path"></Tree>
  </div>
</template>

<script>
import get from 'lodash/get'
import merge from 'lodash/merge'
import ArrowIcon from './ArrowIcon.vue'
import utils from './utils.js'

export default {
  name: 'Tree',
  components: { ArrowIcon },
  props: {
    path: {
      type: String,
      default: ''
    }
  },
  methods: {
    getNodeByPath(path = '') {
      if (!path || path === '' || path === '.') {
        return this.$store.state.tree
      }
      const paths = utils.getPaths(path)
      return get(this.$store.state.tree, paths)
    },
    toggle() {
      const node = this.node
      if (node.isDir) {
        node.open = !node.open
        if (!node.loaded) {
          node.loaded = true
          this.$store.dispatch('getNode', node.path)
          .then((n) => {
            merge(node, n)
            this.$store.commit('set', node)
            this.$store.commit('setActive', node.path)
          })
        } else {
          this.$store.commit('set', node)
          this.$store.commit('setActive', node.path)
        }
      } else {
        if (!node.loaded) {
          node.loaded = true
          this.$store.dispatch('getNode', node.path)
          .then((n) => {
            merge(node, n)
            this.$store.commit('setTextArea', node.content)
            this.$store.commit('setActive', node.path)
            this.$store.commit('set', node)
          })
        } else {
          this.$store.commit('setActive', node.path)
          this.$store.commit('set', node)
        }
      }
    },
  },
  computed: {
    node() {
      return this.getNodeByPath(this.path)
    },
    level() {
      const path = this.path
      if (!path || path === '' || path === '.') {
        return 0
      }
      return path.split('/').filter(s => s).length
    },
    active() {
      return this.$store.state.active === this.node.path
    }
  },
}
</script>

<style lang="css" scoped>
  .tree {
    box-sizing: border-box;
    width: 100%;
  }
  .header {
    cursor: pointer;
    color: #9DA5B3;
    height: 2em;
    line-height: 2em;
  }
  .header:hover, .active {
    background-color: #2C313A;
  }
  .header .leaf {
    padding-left: 18px;
  }
</style>
