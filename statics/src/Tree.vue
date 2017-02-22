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
import ArrowIcon from './ArrowIcon.vue'

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
      const splitpaths = path.split('/')
      const paths = []
      for (let i = 0; i < splitpaths.length; i++) {
        if (splitpaths[i]) {
          paths.push('children', splitpaths[i])
        }
      }
      return get(this.$store.state.tree, paths)
    },
    toggle() {
      const node = this.node
      node.open = !node.open
      if (!node.loaded) {
        node.loaded = true
        this.$store.dispatch('getNode', node.path)
      }
      this.$store.commit('setActive', node.path)
      this.$store.commit('set', node)
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
